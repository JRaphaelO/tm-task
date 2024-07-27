package domain_test

import (
	"testing"
	"time"

	"github.com/jraphaelo/taskmanagement/task/src/core/task/domain"
	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	t.Parallel()

	t.Run("Test New Task", func(t *testing.T) {
		t.Parallel()

		parsedTime, err := time.Parse(time.RFC3339, "2021-12-31T23:59:59Z")
		if err != nil {
			t.Fatal(err)
		}

		task, err := domain.NewTask(
			"Task 1",
			parsedTime,
			nil,
		)

		require.Nil(t, err)
		require.NotEmpty(t, task)

	})

	t.Run("Test New Task with description", func(t *testing.T) {
		t.Parallel()

		parsedTime, err := time.Parse(time.RFC3339, "2021-12-31T23:59:59Z")
		if err != nil {
			t.Fatal(err)
		}

		description := "Description of task 1"
		task, err := domain.NewTask(
			"Task 1",
			parsedTime,
			&description,
		)

		require.Nil(t, err)
		require.NotEmpty(t, task)
		require.Equal(t, description, task.Description)

	})

	t.Run("Test New Task with invalid title", func(t *testing.T) {
		t.Parallel()

		parsedTime, err := time.Parse(time.RFC3339, "2021-12-31T23:59:59Z")
		if err != nil {
			t.Fatal(err)
		}

		task, err := domain.NewTask(
			"",
			parsedTime,
			nil,
		)

		require.NotNil(t, err)
		require.Empty(t, task)
	})

	t.Run("Test New Task with invalid prevision date", func(t *testing.T) {
		t.Parallel()

		task, err := domain.NewTask(
			"Task 1",
			time.Time{},
			nil,
		)

		require.NotNil(t, err)
		require.Empty(t, task)
	})

	t.Run("Test New Task with title so long", func(t *testing.T) {
		t.Parallel()

		parsedTime, err := time.Parse(time.RFC3339, "2021-12-31T23:59:59Z")
		if err != nil {
			t.Fatal(err)
		}

		task, err := domain.NewTask(
			string(make([]byte, 256)),
			parsedTime,
			nil,
		)

		require.NotNil(t, err)
		require.Empty(t, task)
	})
}
