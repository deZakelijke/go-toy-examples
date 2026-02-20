package main

import (
	"fmt"
	"github.com/deZakelijke/go-toy-examples/sorting"
	"github.com/gin-gonic/gin"
	"net/http"
)

type inputData struct {
	UnsortedData []float64 `form:"unsorted_data" json:"unsorted_data" binding:"required"`
}

type outputData struct {
	SortedData []float64 `json:"sorted_data"`
}

func main() {
	router := gin.Default()
	router.POST("/sort", sortData)

	router.Run("localhost:8601")
}

func sortData(c *gin.Context) {
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
