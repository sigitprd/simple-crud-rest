package main

import "github.com/sigitprd/simple-crud-rest/routes"

func main() {
	e := routes.Init()

	e.Logger.Fatal(e.Start(":8000"))
}
