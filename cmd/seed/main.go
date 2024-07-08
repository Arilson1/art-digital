package main

import (
	"fmt"

	"github.com/Arilson1/art-digital/data"
	"github.com/Arilson1/art-digital/db"
)

func main() {
	fmt.Println("Preenchendo o banco de dados")

	db.Init()
	db.DB.AutoMigrate(&data.User{})

	db.DB.Create(&data.User{
		Name:  "Teste",
		Email: "teste@hotmail.com",
	})
}
