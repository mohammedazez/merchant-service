package main

import (
	"merchant-service/app"
	DataRes "merchant-service/datasources/mysql"
)

func main() {
	DataRes.Init_db()
	app.StartApplication()
}
