package task

import (
	"errors"
	"time"

	"github.com/jraphaelo/taskmanagement/task/src/core/task/domain/entities"
	repositories "github.com/jraphaelo/taskmanagement/task/src/core/task/domain/repository"
)

type CreateTaskUseCaseRequest struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	PrevisionDate string `json:"prevision_date"`
}

type CreateTaskUseCase struct {
	TaskRepository repositories.TaskRepository
}

func (uc *CreateTaskUseCase) Execute(data CreateTaskUseCaseRequest) error {
	parsedDate, err := time.Parse("2006-01-02", data.PrevisionDate)
	if err != nil {
		return errors.New("invalid prevision date format")
	}

	task, err := entities.NewTask(data.Title, parsedDate, &data.Description)
	if err != nil {
		return err
	}

	err = uc.TaskRepository.Save(*task)
	return err
}
