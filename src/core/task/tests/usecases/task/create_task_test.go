package task_test

import (
	"errors"
	"testing"

	"github.com/jraphaelo/taskmanagement/task/src/core/task/infrastructure/database/memory"
	"github.com/jraphaelo/taskmanagement/task/src/core/task/useCases/task"
	"github.com/stretchr/testify/assert"
)

func TestCreateTaskUseCase_Execute(t *testing.T) {
	mockRepo := memory.NewInMemoryTaskRepository()
	uc := &task.CreateTaskUseCase{TaskRepository: mockRepo}

	tests := []struct {
		name          string
		request       task.CreateTaskUseCaseRequest
		expectedError error
	}{
		{
			name: "successfully create task",
			request: task.CreateTaskUseCaseRequest{
				Title:         "Test Task",
				PrevisionDate: "2024-07-27",
				Description:   "Task description",
			},
			expectedError: nil,
		},
		{
			name: "invalid prevision date format",
			request: task.CreateTaskUseCaseRequest{
				Title:         "Test Task",
				PrevisionDate: "invalid-date",
				Description:   "Task description",
			},
			expectedError: errors.New("invalid prevision date format"),
		},
		{
			name: "invalid prevision date format",
			request: task.CreateTaskUseCaseRequest{
				Title:         "Test Task",
				PrevisionDate: "122",
				Description:   "Task description",
			},
			expectedError: errors.New("invalid prevision date format"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.Execute(tt.request)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
