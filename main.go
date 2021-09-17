package main

import (
	"github.com/sigitprd/simple-crud-rest/database"
	"github.com/sigitprd/simple-crud-rest/routes"
)

func main() {
	database.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8000"))
}
