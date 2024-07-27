package memory

import (
	"errors"
	"sync"

	"github.com/jraphaelo/taskmanagement/task/src/core/_shared/domain"
	"github.com/jraphaelo/taskmanagement/task/src/core/task/domain/entities"
	repositories "github.com/jraphaelo/taskmanagement/task/src/core/task/domain/repository"
)

type InMemoryTaskRepository struct {
	mu    sync.Mutex
	tasks map[string]entities.Task
}

func NewInMemoryTaskRepository() repositories.TaskRepository {
	return &InMemoryTaskRepository{
		tasks: map[string]entities.Task{},
	}
}

func (r *InMemoryTaskRepository) Save(task entities.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tasks[task.GetID()] = task
	return nil
}

func (r *InMemoryTaskRepository) GetAll(pagination domain.Pagination) ([]entities.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	tasks := []entities.Task{}
	if pagination.CurrentPage == 0 {
		pagination.CurrentPage = 1
	}

	if pagination.PerPage == 0 {
		pagination.PerPage = 10
	}

	start := (pagination.CurrentPage - 1) * pagination.PerPage
	end := start + pagination.PerPage

	// Convert map to slice
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}

	if start > len(tasks) {
		return []entities.Task{}, errors.New("no tasks found")
	}

	if end > len(tasks) {
		end = len(tasks)
	}

	tasks = tasks[start:end]

	return tasks, nil
}

func (r *InMemoryTaskRepository) GetID(id string) (entities.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.tasks[id]; !ok {
		return entities.Task{}, errors.New("task not found")
	}

	return r.tasks[id], nil
}
