package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jraphaelo/taskmanagement/task/src/core/task/controllers"
	"github.com/jraphaelo/taskmanagement/task/src/core/task/infrastructure/database/memory"
	"github.com/jraphaelo/taskmanagement/task/src/core/task/usecases/task"
)

func main() {
	taskRepository := memory.NewInMemoryTaskRepository()

	createTaskUseCase := task.CreateTaskUseCase{TaskRepository: taskRepository}
	getTaskUseCase := task.GetTaskUseCase{TaskRepository: taskRepository}

	// Here we would start the server and handle the requests
	createTaskController := controllers.TaskController{CreateTaskUseCase: &createTaskUseCase, GetTaskUseCase: &getTaskUseCase}

	// Here we would start the server and handle the requests
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/task", createTaskController.CreateTask).Methods("POST")
	router.HandleFunc("/api/v1/task", createTaskController.GetTask).Methods("GET")

	// Here we would start the server and handle the requests
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
