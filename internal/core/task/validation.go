package task

import (
	"errors"
	"strings"
	"time"
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
	if err := t.ValidateTitle(t.Title); err != nil {
		return err
	}
	if err := validateStatus(t.Status); err != nil {
		return err
	}
	if err := validatePriority(t.Priority); err != nil {
		return err
	}
	if err := validateEstimate(t.EstimateMin); err != nil {
		return err
	}
	if err := validateActual(t.ActualMin); err != nil {
		return err
	}
	if err := validateCompletion(t); err != nil {
		return err
	}
	if err := validateTimestamps(t.CreatedAt, t.UpdatedAt); err != nil {
		return err
	}
	return nil
}

func (t *Task) ValidateTitle(title string) error {
	if strings.TrimSpace(title) == "" {
		return ErrEmptyTitle
	}
	if len(title) > 255 {
		return ErrTitleTooLong
	}
	return nil
}

func validateStatus(s Status) error {
	switch s {
	case Todo, InProgress, Done, Aborted:
		return nil
	default:
		return ErrInvalidStatus
	}
}

func validatePriority(p Priority) error {
	switch p {
	case Low, Medium, High:
		return nil
	default:
		return ErrInvalidPriority
	}
}

func validateEstimate(estimate int) error {
	if estimate < 0 {
		return ErrNegativeEstimate
	}
	return nil
}

func validateActual(actual int) error {
	if actual < 0 {
		return ErrNegativeActual
	}
	return nil
}

func validateCompletion(t *Task) error {
	if t.CompletedAt != nil && t.Status != Done {
		return ErrCompletedAtWithoutDoneStatus
	}
	if t.Status == Done && t.CompletedAt == nil {
		return errors.New("completed_at must be set if status is 'done'")
	}
	return nil
}

func validateTimestamps(created, updated time.Time) error {
	if created.IsZero() || updated.IsZero() {
		return errors.New("created_at and updated_at must be set")
	}
	if updated.Before(created) {
		return errors.New("updated_at cannot be before created_at")
	}
	return nil
}
