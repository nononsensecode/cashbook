package main

import (
	"cashbook/model"
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
	app.Run()
}
