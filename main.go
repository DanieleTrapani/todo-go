package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Content: "Buy milk", Completed: false},
	{ID: "2", Content: "Buy eggs", Completed: false},
	{ID: "3", Content: "Buy bread", Completed: false},
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", addTodo)
	router.Run(":3000")
}

func getTodos(request *gin.Context) {
	request.IndentedJSON(http.StatusOK, todos)
}

func addTodo(request *gin.Context) {
	var newTodo todo

	if err := request.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	request.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(request *gin.Context) {
	id := request.Param("id")
	todo, err := findTodo(id)

	if err != nil {
		request.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	request.IndentedJSON(http.StatusOK, todo)
}

func findTodo(id string) (*todo, error) {
	for _, t := range todos {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, errors.New("Todo Not Found")
}
