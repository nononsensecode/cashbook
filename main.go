package main

import (
	"cashbook/cashbookerror"
	"cashbook/model"
	"cashbook/utils"
	"database/sql"
	"log"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"

	_ "github.com/mattn/go-sqlite3"
)

var initDb string

func startup() string {
	return initDb
}

func init() {
	var err error
	model.DB, err = sql.Open("sqlite3", "./cashbook.db")
	if err != nil {
		log.Println(err)
		initDb = "Database cannot be created"
	}

	err = model.CreateTables()
	if err != nil {
		log.Println(err)
		initDb = "Tables cannot be created"
	}

	initDb = "Database and tables created successfully"
}

func findAllAccountCodes() string {
	accountCodes, err := model.FindAllAccountCodes()
	if err != nil {
		log.Println(err)
		return utils.RespondWithError(500, "Unknown error")
	}
	
	return utils.RespondWithJSON(accountCodes)
}

func findAccountCodeByCode(code string) string {
	accountCode, err := model.FindAccountCodeByCode(code)
	if err != nil {
		accountCodeError, ok := err.(*cashbookerror.AccountCodeError)
		if ok {
			log.Println(err)
			switch{
			case accountCodeError.Code == 404:
				return utils.RespondWithError(accountCodeError.Code, accountCodeError.Err.Error())
			default:
				return utils.RespondWithError(500, "Unknown error")
			}
		}
	}

	return utils.RespondWithJSON(accountCode)
}

func createAccountCode(accountCode model.AccountCode) string {
	createdAccCode, err := model.AddAccountCode(accountCode)
	if err != nil {
		log.Println(err)
		return utils.RespondWithError(500, "Account could not be created due to unknown reasons")
	}

	return utils.RespondWithJSON(createdAccCode)
}

func updateAccountCode(accountCode model.AccountCode) string {
	updatedAccCode, err := model.UpdateAccountCode(accountCode)
	if err != nil {
		log.Println(err)
		return utils.RespondWithError(500, "Account code could not be updated")
	}

	return utils.RespondWithJSON(updatedAccCode)
}

func main() {

	js := mewn.String("./frontend/dist/my-app/main.js")
	css := mewn.String("./frontend/dist/my-app/styles.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "CashBook",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(startup)
	app.Bind(findAllAccountCodes)
	app.Bind(findAccountCodeByCode)
	app.Bind(createAccountCode)
	app.Bind(updateAccountCode)
	app.Run()
}
