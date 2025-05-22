package task

import (
	"errors"
	"strings"
)

var (
	ErrEmptyTitle                   = errors.New("task title cannot be empty")
	ErrTitleTooLong                 = errors.New("task title cannot exceed 255 characters")
	ErrInvalidStatus                = errors.New("invalid task status")
	ErrInvalidPriority              = errors.New("invalid task priority")
	ErrNegativeEstimate             = errors.New("estimate cannot be negative")
	ErrNegativeActual               = errors.New("actual time cannot be negative")
	ErrCompletedAtWithoutDoneStatus = errors.New(
		"completed_at should only be set if status is 'done'",
	)
)

// Validate checks whether the Task fields are valid according to business rules.
func (t *Task) Validate() error {
	if err := t.ValidateTitle(); err != nil {
		return err
	}
	if err := t.ValidateStatus(); err != nil {
		return err
	}
	if err := t.ValidatePriority(); err != nil {
		return err
	}
	if err := t.ValidateEstimate(); err != nil {
		return err
	}
	if err := t.ValidateActual(); err != nil {
		return err
	}
	if err := t.ValidateCompletion(); err != nil {
		return err
	}
	if err := t.ValidateTimestamps(); err != nil {
		return err
	}
	return nil
}

func (t *Task) ValidateTitle() error {
	if strings.TrimSpace(t.Title) == "" {
		return ErrEmptyTitle
	}
	if len(t.Title) > 255 {
		return ErrTitleTooLong
	}
	return nil
}

func (t *Task) ValidateStatus() error {
	switch t.Status {
	case Todo, InProgress, Done, Aborted:
		return nil
	default:
		return ErrInvalidStatus
	}
}

func (t *Task) ValidatePriority() error {
	switch t.Priority {
	case Low, Medium, High:
		return nil
	default:
		return ErrInvalidPriority
	}
}

func (t *Task) ValidateEstimate() error {
	if t.EstimateMin < 0 {
		return ErrNegativeEstimate
	}
	return nil
}

func (t *Task) ValidateActual() error {
	if t.ActualMin < 0 {
		return ErrNegativeActual
	}
	return nil
}

func (t *Task) ValidateCompletion() error {
	if t.CompletedAt != nil && t.Status != Done {
		return ErrCompletedAtWithoutDoneStatus
	}
	if t.Status == Done && t.CompletedAt == nil {
		return errors.New("completed_at must be set if status is 'done'")
	}
	return nil
}

func (t *Task) ValidateTimestamps() error {
	if t.CreatedAt.IsZero() || t.UpdatedAt.IsZero() {
		return errors.New("created_at and updated_at must be set")
	}
	if t.UpdatedAt.Before(t.CreatedAt) {
		return errors.New("updated_at cannot be before created_at")
	}
	return nil
}
