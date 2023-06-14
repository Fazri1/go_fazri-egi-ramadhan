package main

import (
	"toko-elektronik/config"
	"toko-elektronik/route"
)

func main() {
	config.InitDB()
	e := route.New()

	e.Logger.Fatal(e.Start(":8080"))
}
