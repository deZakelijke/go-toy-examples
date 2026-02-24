package sortingapi

import (
	"fmt"
	"github.com/deZakelijke/go-toy-examples/sorting"
	"github.com/deZakelijke/go-toy-examples/todo_database"
	"github.com/gin-gonic/gin"
	"net/http"
)

type inputData struct {
	UnsortedData []float64 `form:"unsorted_data" json:"unsorted_data" binding:"required"`
}

type outputData struct {
	SortedData []float64 `json:"sorted_data"`
}

func SortData(c *gin.Context) {
	var newInput inputData
	var data outputData

	if err := c.BindJSON(&newInput); err != nil {
		fmt.Printf("Error: %q\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data.SortedData = sorting.MergeSort(newInput.UnsortedData)

	c.IndentedJSON(http.StatusOK, data)
}

func GetTodoItems(c *gin.Context) {
	items, err := tododatabase.Get(tododatabase.DB)
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, items)

}
func InsertTodoItem(c *gin.Context) {
	var newItem tododatabase.TodoItem

	if err := c.BindJSON(&newItem); err != nil {
		return
	}
	id, err := tododatabase.Insert(tododatabase.DB, newItem)
	if err != nil {
		return
	}

	c.String(http.StatusCreated, fmt.Sprintf("New Id: %d", id))
}

func UpdateTodoItem(c *gin.Context) {
	var updateItem tododatabase.TodoItem

	if err := c.BindJSON(&updateItem); err != nil {
		fmt.Printf("Update item: %q\n", err)
		return
	}

	id, err := tododatabase.Update(tododatabase.DB, updateItem.Id, updateItem.Done)
	if err != nil {
		fmt.Printf("Update item: %q\n", err)
	}

	c.String(http.StatusOK, fmt.Sprintf("Update Id: %d", id))
}
