package main

import (
	"github.com/arie0703/go-headless-cms-api/config"
	"github.com/arie0703/go-headless-cms-api/routes"
)

func main() {
	db := config.InitDB()
	r := routes.SetupRouter(db)
	r.Run(":8080")
}
