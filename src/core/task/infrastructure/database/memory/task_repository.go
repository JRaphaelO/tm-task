package memory

import (
	"errors"
	"sync"

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

func (r *InMemoryTaskRepository) GetID(id string) (entities.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.tasks[id]; !ok {
		return entities.Task{}, errors.New("task not found")
	}

	return r.tasks[id], nil
}
