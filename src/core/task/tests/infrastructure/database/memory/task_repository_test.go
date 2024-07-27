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
