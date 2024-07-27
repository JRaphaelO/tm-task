package repositories

import "github.com/jraphaelo/taskmanagement/task/src/core/task/domain/entities"

type TaskRepository interface {
	Save(task entities.Task) error
	GetID(id string) (entities.Task, error)
}
