package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}

	usecase := NewUseCase(db)

	product, err := usecase.GetProduct(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(product.Name)



}

