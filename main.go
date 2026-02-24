package main

import (
	"fmt"
	"github.com/deZakelijke/go-toy-examples/sorting_api"
	"github.com/deZakelijke/go-toy-examples/todo_database"
	"github.com/gin-gonic/gin"
)

func main() {

	err := tododatabase.GetOrCreateDB("./local_db.db")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := tododatabase.DB.Close()
		if err != nil {
			fmt.Println("Error on closing db: ", err)
		}
	}()
	fmt.Printf("Database connection set up.\n")

	router := gin.Default()
	router.POST("/sort", sortingapi.SortData)

	router.GET("/todo", sortingapi.GetTodoItems)
	router.POST("/todo", sortingapi.InsertTodoItem)
	router.PUT("/todo", sortingapi.UpdateTodoItem)

	go func() {
		if err := router.Run("localhost:8601"); err != nil {
			fmt.Println("Failed to start server: ", err)
		}
	}()
}
