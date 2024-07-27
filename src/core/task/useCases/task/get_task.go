package task

import (
	"errors"

	"github.com/jraphaelo/taskmanagement/task/src/core/_shared/domain"
	repositories "github.com/jraphaelo/taskmanagement/task/src/core/task/domain/repository"
)

type TaskEntityResponse struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	PrevisionDate string `json:"prevision_date"`
	StartedDate   string `json:"started_date"`
	FinishedDate  string `json:"finished_date"`
}

type GetTaskRequest struct {
	ID          string
	CurrentPage int
}

type GetTaskResponse struct {
	Data       []TaskEntityResponse
	Pagination domain.Pagination
}

type GetTaskUseCase struct {
	TaskRepository repositories.TaskRepository
}

func (uc *GetTaskUseCase) Execute(requestData GetTaskRequest) (GetTaskResponse, error) {
	if requestData.ID != "" {
		task, err := uc.TaskRepository.GetID(requestData.ID)
		if err != nil {
			return GetTaskResponse{}, errors.New("task not found")
		}

		return GetTaskResponse{
			Data: []TaskEntityResponse{
				{
					ID:            task.GetID(),
					Title:         task.GetTitle(),
					Description:   task.GetDescription(),
					PrevisionDate: task.GetPrevisionDate().Format("2006-01-02"),
					StartedDate:   task.GetStartedDate().Format("2006-01-02"),
					FinishedDate:  task.GetFinishedDate().Format("2006-01-02"),
				},
			},
			Pagination: domain.Pagination{
				Total:        1,
				PerPage:      1,
				CurrentPage:  1,
				TotalPages:   1,
				NextPage:     0,
				PreviousPage: 0,
			},
		}, nil
	}

	tasks, err := uc.TaskRepository.GetAll(
		domain.Pagination{
			PerPage:     2,
			CurrentPage: requestData.CurrentPage,
		},
	)
	if err != nil {
		return GetTaskResponse{}, errors.New("tasks not found 2")
	}

	data := make([]TaskEntityResponse, 0, len(tasks))
	for _, task := range tasks {
		data = append(data, TaskEntityResponse{
			ID:            task.GetID(),
			Title:         task.GetTitle(),
			Description:   task.GetDescription(),
			PrevisionDate: task.GetPrevisionDate().Format("2006-01-02"),
			StartedDate:   task.GetStartedDate().Format("2006-01-02"),
			FinishedDate:  task.GetFinishedDate().Format("2006-01-02"),
		})
	}

	response := GetTaskResponse{
		Data: data,
		Pagination: domain.Pagination{
			Total:        len(data),
			PerPage:      2,
			CurrentPage:  1,
			TotalPages:   1,
			NextPage:     2,
			PreviousPage: 0,
		},
	}

	return response, nil
}
