package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type TaskInterface interface {
	// Get methods for Task
	GetID() string
	GetTitle() string
	GetDescription() string
	GetStatus() StatusTask
	GetPrevisionDate() time.Time
	GetStartedDate() time.Time
	GetFinishedDate() time.Time
}

type StatusTask string

const (
	CREATED     StatusTask = "CREATED"
	IN_PROGRESS StatusTask = "IN_PROGRESS"
	COMPLETED   StatusTask = "COMPLETED"
	STOPPED     StatusTask = "STOPPED"
)

type Task struct {
	ID            string     `valid:"uuidv4"`
	Title         string     `valid:"required"`
	Description   string     `valid:"optional"`
	Status        StatusTask `valid:"optional"`
	PrevisionDate time.Time  `valid:"required"`
	StartedDate   time.Time  `valid:"-"`
	FinishedDate  time.Time  `valid:"-"`
	CreatedAt     time.Time  `valid:"-"`
	UpdatedAt     time.Time  `valid:"-"`
}

func NewTask(title string, previsionDate time.Time, description *string) (*Task, error) {
	task := &Task{
		ID:            uuid.NewV4().String(),
		Title:         title,
		Description:   "",
		Status:        CREATED,
		PrevisionDate: previsionDate,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if description != nil {
		task.Description = *description
	}

	if err := task.validate(); err != nil {
		return nil, err
	}

	return task, nil
}

func (t *Task) GetID() string {
	return t.ID
}

func (t *Task) GetTitle() string {
	return t.Title
}

func (t *Task) GetDescription() string {
	return t.Description
}

func (t *Task) GetStatus() StatusTask {
	return t.Status
}

func (t *Task) GetPrevisionDate() time.Time {
	return t.PrevisionDate
}

func (t *Task) GetStartedDate() time.Time {
	return t.StartedDate
}

func (t *Task) GetFinishedDate() time.Time {
	return t.FinishedDate
}

func (t *Task) validate() error {
	_, err := govalidator.ValidateStruct(t)
	if err != nil {
		if validationErrors, ok := err.(govalidator.Errors); ok {
			var errorsMessage []string
			for _, validationError := range validationErrors.Errors() {
				if fieldError, ok := validationError.(govalidator.Error); ok {
					errorsMessage = append(errorsMessage, fmt.Sprintf("Field '%s': %s", fieldError.Name, fieldError.Err))
				}
			}
			return errors.New(strings.Join(errorsMessage, ", "))
		}

		return err
	}

	if len(t.Title) > 120 {
		return errors.New("the title must be less than or equal 120 characters")
	}

	return nil
}
