package task_test

import (
	"errors"
	"testing"

	"github.com/jraphaelo/taskmanagement/task/src/core/task/domain/entities"
	"github.com/jraphaelo/taskmanagement/task/src/core/task/useCases/task"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock do repositório de tarefas
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) Save(task entities.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func TestCreateTaskUseCase_Execute(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	uc := &task.CreateTaskUseCase{TaskRepository: mockRepo}

	tests := []struct {
		name          string
		request       task.CreateTaskUseCaseRequest
		mockSaveError error
		expectedError error
	}{
		{
			name: "successfully create task",
			request: task.CreateTaskUseCaseRequest{
				Title:         "Test Task",
				PrevisionDate: "2024-07-27",
				Description:   "Task description",
			},
			mockSaveError: nil,
			expectedError: nil,
		},
		{
			name: "invalid prevision date format",
			request: task.CreateTaskUseCaseRequest{
				Title:         "Test Task",
				PrevisionDate: "invalid-date",
				Description:   "Task description",
			},
			mockSaveError: nil,
			expectedError: errors.New("invalid prevision date format"),
		},
		{
			name: "invalid prevision date format",
			request: task.CreateTaskUseCaseRequest{
				Title:         "Test Task",
				PrevisionDate: "122",
				Description:   "Task description",
			},
			mockSaveError: errors.New("invalid prevision date format"),
			expectedError: errors.New("invalid prevision date format"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockSaveError != nil {
				mockRepo.On("Save", mock.AnythingOfType("entities.Task")).Return(tt.mockSaveError)
			} else {
				mockRepo.On("Save", mock.AnythingOfType("entities.Task")).Return(nil)
			}

			err := uc.Execute(tt.request)

			assert.Equal(t, tt.expectedError, err)
			mockRepo.AssertExpectations(t)
		})
	}
}
