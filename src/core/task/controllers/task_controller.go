package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jraphaelo/taskmanagement/task/src/core/task/usecases/task"
)

type TaskController struct {
	CreateTaskUseCase *task.CreateTaskUseCase
	GetTaskUseCase    *task.GetTaskUseCase
}

func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var taskRequest task.CreateTaskUseCaseRequest
	err := json.NewDecoder(r.Body).Decode(&taskRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.CreateTaskUseCase.Execute(taskRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *TaskController) GetTask(w http.ResponseWriter, r *http.Request) {
	taskID := r.URL.Query().Get("id")        // get task id from query string
	currentPage := r.URL.Query().Get("page") // get current page from query string
	if currentPage == "" {
		currentPage = "1"
	}

	currentPageInt, err := strconv.Atoi(currentPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	requestData := task.GetTaskRequest{
		ID:          taskID,
		CurrentPage: currentPageInt,
	}
	tasks, err := c.GetTaskUseCase.Execute(requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
