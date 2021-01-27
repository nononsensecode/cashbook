package model

import (
	"bufio"
	"cashbook/cashbookerror"
	"database/sql"
	"log"
	"os"
	"strings"
	"time"
)

// DB Database object
var DB *sql.DB

// AccountCode represents the account code in the cashbook
type AccountCode struct {
	ID      int64  `json:"id"`
	AccCode string `json:"accountCode"`
	Desc    string `json:"description"`
}

//CashRegister represents cash register
type CashRegister struct {
	ID      int64     `json:"id"`
	Type    string    `json:"type"`
	OB      float64   `json:"openingBalance"`
	Date    time.Time `json:"date"`
	CRNo    string    `json:"cashRegisterNumber"`
	AccCode string    `json:"accountCode"`
	AccType string    `json:"accountType"`
	Amount  float64   `json:"amount"`
}

const (
	// Receipt indicates account belongs to the receipt side
	Receipt = "receipt"
	// Payment indicates account belongs to the payment side
	Payment = "payment"
)

const (
	// Cash indicates amount exists as cash
	Cash = "cash"
	// Operative indicates amount deposited in operative account
	Operative = "operative"
	// NonOperative indicates amount deposited in the non operative account
	NonOperative = "non operative"
	// Treasury indicates amount deposited in treasury
	Treasury = "treasury"
)

func readAccountCodes() ([]AccountCode, error) {
	currDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	accFile, err := os.Open(currDir + "/model/acc_codes.txt")
	if err != nil {
		return nil, err
	}
	defer accFile.Close()

	scanner := bufio.NewScanner(accFile)
	var accountCodes []AccountCode
	for scanner.Scan() {
		codeAsArray := strings.Split(scanner.Text(), ",")
		accountCodes = append(accountCodes, AccountCode{AccCode: codeAsArray[0], Desc: codeAsArray[1]})
	}

	return accountCodes, nil
}

// InitializeAccountCodes fills acc_code table with data
func InitializeAccountCodes() error {
	accountCodes, err := readAccountCodes()
	if err != nil {
		return err
	}

	transaction, err := DB.Begin()
	if err != nil {
		return err
	}
	defer transaction.Commit()

	insertString := "INSERT INTO acc_code (acc_code, desc) VALUES "
	vals := []interface{}{}

	for _, accountCode := range accountCodes {
		insertString += "(?, ?),"
		vals = append(vals, accountCode.AccCode, accountCode.Desc)
	}

	insertString = insertString[0 : len(insertString)-1]

	statement, err := DB.Prepare(insertString)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(vals...)
	if err != nil {
		return err
	}

	return nil
}

// CreateTables creates article table
func CreateTables() error {
	statement, err := DB.Prepare(`CREATE TABLE IF NOT EXISTS 
		cash_register(id INTEGER PRIMARY KEY AUTOINCREMENT,
			type TEXT CHECK(type IN ('receipt', 'payment')),
			ob REAL, date TEXT, cr_no TEXT, acc_code TEXT,
			acc_type TEXT CHECK(acc_type IN ('cash', 'operative', 'non operative', 'treasury')),
			amount REAL)`)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec()
	if err != nil {
		return err
	}
	log.Println("cash_register table created successfully")

	statement, err = DB.Prepare(`CREATE TABLE IF NOT EXISTS 
		 acc_code(id INTEGER PRIMARY KEY AUTOINCREMENT, 
			acc_code TEXT, desc TEXT)`)
	if err != nil {
		return err
	}

	_, err = statement.Exec()
	if err != nil {
		return err
	}
	log.Println("acc_code table created successfully")

	return nil
}

// AddAccountCode adds an account code row in the table
func AddAccountCode(accountCode AccountCode) (AccountCode, error) {
	transaction, err := DB.Begin()
	if err != nil {
		return AccountCode{}, err
	}
	defer transaction.Rollback()

	statement, err := DB.Prepare("INSERT INTO acc_code (code, desc) VALUES (?, ?)")
	if err != nil {
		return AccountCode{}, err
	}
	defer statement.Close()

	result, err := statement.Exec(accountCode.AccCode, accountCode.Desc)
	if err != nil {
		return AccountCode{}, err
	}

	err = transaction.Commit()
	if err != nil {
		return AccountCode{}, err
	}

	accountCode.ID, err = result.LastInsertId()
	if err != nil {
		return AccountCode{}, err
	}

	return accountCode, nil
}

// UpdateAccountCode updates an account code
func UpdateAccountCode(accountCode AccountCode) (AccountCode, error) {
	transaction, err := DB.Begin()
	if err != nil {
		return AccountCode{}, err
	}
	defer transaction.Rollback()

	statement, err := DB.Prepare("UPDATE acc_code SET code = ?, desc = ? WHERE id = ?")
	if err != nil {
		return AccountCode{}, err
	}
	defer statement.Close()

	result, err := statement.Exec(accountCode.AccCode, accountCode.Desc, accountCode.ID)
	if err != nil {
		return AccountCode{}, err
	}

	err = transaction.Commit()
	if err != nil {
		return AccountCode{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return AccountCode{}, err
	}

	if rowsAffected > 0 {
		return accountCode, nil
	} else {
		return AccountCode{}, cashbookerror.NewAccountCodeError(500, "No rows affected")
	}
}

// FindAllAccountCodes returns all account codes in the table
func FindAllAccountCodes() ([]AccountCode, error) {
	var accountCodes []AccountCode

	rows, err := DB.Query("SELECT * FROM acc_code")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var accountCode AccountCode

		err = rows.Scan(&accountCode.ID, &accountCode.AccCode, &accountCode.Desc)
		if err != nil {
			return nil, err
		}

		accountCodes = append(accountCodes, accountCode)
	}

	return accountCodes, nil
}

// FindAccountCodeByCode find the account code by the `code` supplied
func FindAccountCodeByCode(code string) (AccountCode, error) {
	var accountCode AccountCode

	err := DB.QueryRow("SELECT * FROM acc_code WHERE code = ?", code).
		Scan(&accountCode.ID, &accountCode.AccCode, &accountCode.Desc)

	switch {
	case err == sql.ErrNoRows:
		return AccountCode{}, cashbookerror.AccountCodeDoesNotExistError(code)
	case err != nil:
		return AccountCode{}, cashbookerror.UnknownAccountCodeError(err)
	default:
		return accountCode, nil
	}
}
