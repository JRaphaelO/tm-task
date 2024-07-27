package task_test

import (
	"testing"
	"time"

	"github.com/jraphaelo/taskmanagement/task/src/core/task/domain/entities"
	"github.com/jraphaelo/taskmanagement/task/src/core/task/infrastructure/database/memory"
	"github.com/jraphaelo/taskmanagement/task/src/core/task/usecases/task"
	"github.com/stretchr/testify/require"
)

func TestGetTaskUseCase_Execute(t *testing.T) {
	t.Parallel()

	mockRepo := memory.NewInMemoryTaskRepository()
	uc := &task.GetTaskUseCase{TaskRepository: mockRepo}

	t.Run("successfully get task", func(t *testing.T) {
		parsedDate, _ := time.Parse("2006-01-02", "2024-07-27")
		description := "Task description"
		taskModel, _ := entities.NewTask("Test Task", parsedDate, &description)
		mockRepo.Save(*taskModel)

		taskID := taskModel.GetID()
		request := task.GetTaskRequest{
			ID:          taskID,
			CurrentPage: 1,
		}
		response, err := uc.Execute(request)

		require.Nil(t, err)
		require.Equal(t, 1, len(response.Data))
		require.Equal(t, taskID, response.Data[0].ID)
		require.Equal(t, 1, response.Pagination.Total)
	})

	t.Run("task not found", func(t *testing.T) {
		taskID := "invalid-id"
		request := task.GetTaskRequest{
			ID:          taskID,
			CurrentPage: 1,
		}
		_, err := uc.Execute(request)

		require.Error(t, err)
		require.Equal(t, "task not found", err.Error())
	})

	t.Run("get all tasks", func(t *testing.T) {
		parsedDate, _ := time.Parse("2006-01-02", "2024-07-27")
		description := "Task description"
		taskModel, _ := entities.NewTask("Test Task", parsedDate, &description)
		mockRepo.Save(*taskModel)

		request := task.GetTaskRequest{
			CurrentPage: 1,
		}
		response, err := uc.Execute(request)

		require.Nil(t, err)
		require.Equal(t, 1, len(response.Data))
		require.Equal(t, 1, response.Pagination.Total)
	})

	t.Run("get all tasks with pagination", func(t *testing.T) {
		parsedDate, _ := time.Parse("2006-01-02", "2024-07-27")
		description := "Task description"
		taskModel, _ := entities.NewTask("Test Task", parsedDate, &description)
		mockRepo.Save(*taskModel)

		taskModel, _ = entities.NewTask("Test Task 2", parsedDate, &description)
		mockRepo.Save(*taskModel)

		taskModel, _ = entities.NewTask("Test Task 3", parsedDate, &description)
		mockRepo.Save(*taskModel)

		request := task.GetTaskRequest{
			CurrentPage: 1,
		}
		response, err := uc.Execute(request)

		require.Nil(t, err)
		require.Equal(t, 2, len(response.Data))
		require.Equal(t, 2, response.Pagination.Total)
	})

	t.Run("tasks not found", func(t *testing.T) {
		request := task.GetTaskRequest{
			CurrentPage: 3,
		}
		_, err := uc.Execute(request)

		require.Error(t, err)
		require.Equal(t, "tasks not found", err.Error())
	})

	t.Run("tasks for other page", func(t *testing.T) {
		parsedDate, _ := time.Parse("2006-01-02", "2024-07-27")
		description := "Task description"
		taskModel, _ := entities.NewTask("Test Task", parsedDate, &description)
		mockRepo.Save(*taskModel)

		taskModel, _ = entities.NewTask("Test Task 2", parsedDate, &description)
		mockRepo.Save(*taskModel)

		taskModel, _ = entities.NewTask("Test Task 3", parsedDate, &description)
		mockRepo.Save(*taskModel)

		request := task.GetTaskRequest{
			CurrentPage: 2,
		}
		response, err := uc.Execute(request)

		require.Nil(t, err)
		require.Equal(t, 1, len(response.Data))
		require.Equal(t, 1, response.Pagination.Total)
	})
}
