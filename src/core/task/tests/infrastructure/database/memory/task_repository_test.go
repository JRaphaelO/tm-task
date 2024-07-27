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

func TestInMemoryTaskRepository_GetAll(t *testing.T) {
	t.Parallel()

	t.Run("Should get all tasks", func(t *testing.T) {
		t.Parallel()

		// Arrange
		repository := memory.NewInMemoryTaskRepository()
		parsedDate, _ := time.Parse("2006-01-02", "2021-12-31")
		description := "Description"
		task1, _ := entities.NewTask("Task 1", parsedDate, &description)
		task2, _ := entities.NewTask("Task 2", parsedDate, &description)
		repository.Save(*task1)
		repository.Save(*task2)

		// Act
		tasks, err := repository.GetAll()

		// Assert
		require.Nil(t, err)
		require.Len(t, tasks, 2)
		require.Equal(t, task1.GetID(), tasks[0].GetID())
		require.Equal(t, task1.GetTitle(), tasks[0].GetTitle())
		require.Equal(t, task1.GetPrevisionDate(), tasks[0].GetPrevisionDate())
		require.Equal(t, task1.GetDescription(), tasks[0].GetDescription())
		require.Equal(t, task2.GetID(), tasks[1].GetID())
		require.Equal(t, task2.GetTitle(), tasks[1].GetTitle())
		require.Equal(t, task2.GetPrevisionDate(), tasks[1].GetPrevisionDate())
		require.Equal(t, task2.GetDescription(), tasks[1].GetDescription())
	})
}
