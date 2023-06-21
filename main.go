package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id" `
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

func main() {
	router := gin.Default()
	router.GET("/", getRoot)
	router.GET("/todos", getTodos)

	router.POST("/todos", postTodo)

	router.Run("localhost:8080")
}

func getRoot(context *gin.Context) {
	context.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "Hello Golang"},
	)
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(
		http.StatusOK,
		gin.H{"status": http.StatusOK, "todos": todos},
	)
}

func postTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	if newTodo.ID == "" || newTodo.Item == "" {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"status": http.StatusBadRequest},
		)
	} else {
		n := len(todos)
		exists := false

		for i := 0; i < n; i++ {
			if newTodo.ID == todos[i].ID {
				exists = true
				break
			}
		}

		if exists {
			context.IndentedJSON(
				http.StatusBadRequest,
				gin.H{"status": http.StatusBadRequest},
			)
		} else {
			todos = append(todos, newTodo)
			context.IndentedJSON(
				http.StatusCreated,
				gin.H{"status": http.StatusCreated, "todo": newTodo},
			)
		}
	}
}
