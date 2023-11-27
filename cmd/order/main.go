package main

import (
	"database/sql"
	"fmt"

	"github.com/adrianostankewicz/go-intensivo/internal/infra/database"
	"github.com/adrianostankewicz/go-intensivo/internal/usecase"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := uc.OrderInput{
		ID:    "123",
		Price: 10.0,
		Tax:   1.0,
	}
	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)

}
