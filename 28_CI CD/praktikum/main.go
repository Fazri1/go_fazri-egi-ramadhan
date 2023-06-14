package main

import (
	"restful_api_testing/config"
	"restful_api_testing/routes"
)

func main() {
  config.InitDB()
  e := routes.New()
  // start the server, and log if it fails
  e.Logger.Fatal(e.Start(":8080"))
}