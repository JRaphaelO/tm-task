package memory_test

import (
	"testing"
	"time"

	"github.com/jraphaelo/taskmanagement/task/src/core/task/domain/entities"
	"github.com/jraphaelo/taskmanagement/task/src/core/task/infrastructure/database/memory"
	"github.com/stretchr/testify/require"
)

func TestInMemoryTaskRepository_Create(t *testing.T) {
	t.Parallel()

	t.Run("Should create a task", func(t *testing.T) {
		t.Parallel()

		// Arrange
		repository := memory.NewInMemoryTaskRepository()
		parsedDate, _ := time.Parse("2006-01-02", "2021-12-31")
		description := "Description"
		task, _ := entities.NewTask("Task 1", parsedDate, &description)

		// Act
		err := repository.Save(*task)

		// Assert
		require.Nil(t, err)
	})

	t.Run("Should create a task without description", func(t *testing.T) {
		t.Parallel()

		// Arrange
		repository := memory.NewInMemoryTaskRepository()
		parsedDate, _ := time.Parse("2006-01-02", "2021-12-31")
		task, _ := entities.NewTask("Task 1", parsedDate, nil)

		// Act
		err := repository.Save(*task)

		// Assert
		require.Nil(t, err)
	})

}

func TestInMemoryTaskRepository_GetByID(t *testing.T) {
	t.Parallel()

	t.Run("Should get a task by ID", func(t *testing.T) {
		t.Parallel()

		// Arrange
		repository := memory.NewInMemoryTaskRepository()
		parsedDate, _ := time.Parse("2006-01-02", "2021-12-31")
		description := "Description"
		task, _ := entities.NewTask("Task 1", parsedDate, &description)
		repository.Save(*task)

		// Act
		taskFound, err := repository.GetID(task.GetID())

		// Assert
		require.Nil(t, err)
		require.Equal(t, task.GetID(), taskFound.GetID())
		require.Equal(t, task.GetTitle(), taskFound.GetTitle())
		require.Equal(t, task.GetPrevisionDate(), taskFound.GetPrevisionDate())
		require.Equal(t, task.GetDescription(), taskFound.GetDescription())
	})

	t.Run("Should return an error when task not found", func(t *testing.T) {
		t.Parallel()

		// Arrange
		repository := memory.NewInMemoryTaskRepository()

		// Act
		_, err := repository.GetID("non-existent-id")

		// Assert
		require.NotNil(t, err)
		require.Equal(t, "task not found", err.Error())
	})
}


