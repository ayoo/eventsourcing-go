package main

import (
	uuid "github.com/satori/go.uuid"
)

// Event struct
type Event struct {
	AccID string
	Type  string
}

// CreateEvent struct
type CreateEvent struct {
	Event
	AccName string
}

// DepositEvent struct
type DepositEvent struct {
	Event
	Amount int
}

// WithdrawEvent struct
type WithdrawEvent struct {
	Event
	Amount int
}

// TransferEvent struct
type TransferEvent struct {
	Event
	TargetID string
	Amount   int
}

// NewCreateAccountEvent func
func NewCreateAccountEvent(name string) CreateEvent {
	event := &CreateEvent{}
	event.Type = "CreateEvent"
	event.AccID = uuid.NewV4().String()
	event.AccName = name
	return *event
}
