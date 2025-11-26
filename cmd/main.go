package main

import (
	"fmt"
	"metabeat/internal/db"

	"metabeat/cmd/server"

	"github.com/go-playground/validator/v10"
)

func main() {
	DB := db.InitDB()
	validator := validator.New()
	err := server.Main(DB, validator)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
