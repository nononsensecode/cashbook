package model

import (
	"database/sql"
	"time"
)

// DB Database object
var DB *sql.DB

// AccountCode represents the account code in the cashbook
type AccountCode struct {
	ID      int64  `json:"Id"`
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

// CreateTables creates article table
func CreateTables() error {
	statement, err := DB.Prepare(`CREATE TABLE IF NOT EXISTS 
		cash_register(id INTEGER PRIMARY KEY AUTOINCREMENT,
			type TEXT CHECK( type IN ('receipt', 'payment') ),
			ob REAL, date TEXT, cr_no TEXT, acc_code TEXT,
			acc_type TEXT CHECK( acc_type IN ('cash', 'operative', 'non operative', 'treasury' ),
			amount REAL) ) )`)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec()
	if err != nil {
		return err
	}

	statement, err = DB.Prepare(`CREATE TABLE IF NOT EXISTS 
		 acc_code(id INTEGER PRIMARY KEY AUTOINCREMENT, 
			acc_code TEXT, desc TEXT)`)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec()
	if err != nil {
		return err
	}

	return nil
}
