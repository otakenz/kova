package task

import (
	"fmt"

	"github.com/qmuntal/stateless"
)

// Trigger represents actions that cause status transitions
type Trigger string

const (
	TriggerStart    Trigger = "start"
	TriggerComplete Trigger = "complete"
	TriggerAbort    Trigger = "abort"
	TriggerResume   Trigger = "resume"
)

// NewStateMachine returns a state machine configured for task transitions
func NewStateMachine(current Status) *stateless.StateMachine {
	sm := stateless.NewStateMachine(string(current))

	sm.Configure(string(Todo)).
		Permit(string(TriggerStart), string(InProgress)).
		Permit(string(TriggerAbort), string(Aborted))

	sm.Configure(string(InProgress)).
		Permit(string(TriggerComplete), string(Done)).
		Permit(string(TriggerAbort), string(Aborted))

	sm.Configure(string(Aborted)).
		Permit(string(TriggerResume), string(Todo))

	return sm
}

func ParseTrigger(s string) (Trigger, error) {
	switch s {
	case string(TriggerStart):
		return TriggerStart, nil
	case string(TriggerComplete):
		return TriggerComplete, nil
	case string(TriggerAbort):
		return TriggerAbort, nil
	case string(TriggerResume):
		return TriggerResume, nil
	default:
		return "", fmt.Errorf("invalid trigger: %q", s)
	}
}
