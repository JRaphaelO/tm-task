package domain_test

import (
	"testing"
	"time"

	domain "github.com/jraphaelo/taskmanagement/task/src/core/task/domain/entities"
	"github.com/stretchr/testify/require"
)

func TestTask_NewTask(t *testing.T) {
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

func TestTask_UpdateTask(t *testing.T) {
	t.Parallel()

	t.Run("Test Update Task", func(t *testing.T) {
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

		newTitle := "Task 1 updated"
		newDescription := "Description of task 1 updated"
		newParsedTime, err := time.Parse(time.RFC3339, "2021-12-31T23:59:59Z")
		if err != nil {
			t.Fatal(err)
		}

		err = task.UpdateTask(
			&newTitle,
			&newDescription,
			&newParsedTime,
			nil,
			nil,
		)

		require.Nil(t, err)
		require.Equal(t, newTitle, task.Title)
		require.Equal(t, newDescription, task.Description)
		require.Equal(t, newParsedTime, task.PrevisionDate)
	})

	t.Run("Test Update Task with invalid title", func(t *testing.T) {
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

		newTitle := ""
		newDescription := "Description of task 1 updated"
		newParsedTime, err := time.Parse(time.RFC3339, "2021-12-31T23:59:59Z")
		if err != nil {
			t.Fatal(err)
		}

		err = task.UpdateTask(
			&newTitle,
			&newDescription,
			&newParsedTime,
			nil,
			nil,
		)

		require.NotNil(t, err)
		require.Equal(t, "Field 'Title': non zero value required", err.Error())
	})

	t.Run("Test Update Task with invalid prevision date", func(t *testing.T) {
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

		newTitle := "Task 1 updated"
		newDescription := "Description of task 1 updated"
		newParsedTime := time.Time{}

		err = task.UpdateTask(
			&newTitle,
			&newDescription,
			&newParsedTime,
			nil,
			nil,
		)

		require.NotNil(t, err)
		require.Equal(t, "Field 'PrevisionDate': non zero value required", err.Error())
	})

	t.Run("Test Update Task with title so long", func(t *testing.T) {
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

		newTitle := string(make([]byte, 121))
		newDescription := "Description of task 1 updated"
		newParsedTime, err := time.Parse(time.RFC3339, "2021-12-31T23:59:59Z")
		if err != nil {
			t.Fatal(err)
		}

		err = task.UpdateTask(
			&newTitle,
			&newDescription,
			&newParsedTime,
			nil,
			nil,
		)

		require.NotNil(t, err)
		require.Equal(t, "the title must be less than or equal 120 characters", err.Error())
	})
}

func TestTask_StartTask(t *testing.T) {
	t.Parallel()

	t.Run("Test Start Task", func(t *testing.T) {
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

		err = task.StartTask()

		require.Nil(t, err)
		require.Equal(t, domain.IN_PROGRESS, task.Status)
		require.NotEmpty(t, task.StartedDate)
		require.NotEmpty(t, task.UpdatedAt)
	})

	t.Run("Test Start Task already started", func(t *testing.T) {
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

		err = task.StartTask()

		require.Nil(t, err)
		require.Equal(t, domain.IN_PROGRESS, task.Status)
		require.NotEmpty(t, task.StartedDate)
		require.NotEmpty(t, task.UpdatedAt)

		err = task.StartTask()

		require.NotNil(t, err)
		require.Equal(t, "task already started", err.Error())
	})
}

func TestTask_StopTask(t *testing.T) {
	t.Parallel()

	t.Run("Test Stop Task", func(t *testing.T) {
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

		err = task.StartTask()

		require.Nil(t, err)
		require.Equal(t, domain.IN_PROGRESS, task.Status)
		require.NotEmpty(t, task.StartedDate)
		require.NotEmpty(t, task.UpdatedAt)

		err = task.StopTask()

		require.Nil(t, err)
		require.Equal(t, domain.STOPPED, task.Status)
		require.NotEmpty(t, task.UpdatedAt)
	})

	t.Run("Test Stop Task already stopped", func(t *testing.T) {
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

		err = task.StartTask()

		require.Nil(t, err)
		require.Equal(t, domain.IN_PROGRESS, task.Status)
		require.NotEmpty(t, task.StartedDate)
		require.NotEmpty(t, task.UpdatedAt)

		err = task.StopTask()

		require.Nil(t, err)
		require.Equal(t, domain.STOPPED, task.Status)
		require.NotEmpty(t, task.UpdatedAt)

		err = task.StopTask()

		require.NotNil(t, err)
		require.Equal(t, "task already stopped", err.Error())
	})
}

func TestTask_CompleteTask(t *testing.T) {
	t.Parallel()

	t.Run("Test Complete Task", func(t *testing.T) {
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

		err = task.StartTask()

		require.Nil(t, err)
		require.Equal(t, domain.IN_PROGRESS, task.Status)
		require.NotEmpty(t, task.StartedDate)
		require.NotEmpty(t, task.UpdatedAt)

		err = task.CompleteTask()

		require.Nil(t, err)
		require.Equal(t, domain.COMPLETED, task.Status)
		require.NotEmpty(t, task.FinishedDate)
		require.NotEmpty(t, task.UpdatedAt)
	})

	t.Run("Test Complete Task already completed", func(t *testing.T) {
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

		err = task.StartTask()

		require.Nil(t, err)
		require.Equal(t, domain.IN_PROGRESS, task.Status)
		require.NotEmpty(t, task.StartedDate)
		require.NotEmpty(t, task.UpdatedAt)

		err = task.CompleteTask()

		require.Nil(t, err)
		require.Equal(t, domain.COMPLETED, task.Status)
		require.NotEmpty(t, task.FinishedDate)
		require.NotEmpty(t, task.UpdatedAt)

		err = task.CompleteTask()

		require.NotNil(t, err)
		require.Equal(t, "task already completed", err.Error())
	})
}
