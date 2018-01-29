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

// NewDepositEvent func
func NewDepositEvent(id string, amt int) DepositEvent {
	event := new(DepositEvent)
	event.Type = "DepositEvent"
	event.AccID = id
	event.Amount = amt
	return *event
}

//NewWithdrawEvent func
func NewWithdrawEvent(id string, amt int) WithdrawEvent {
	event := new(WithdrawEvent)
	event.Type = "WithdrawEvent"
	event.AccID = id
	event.Amount = amt
	return *event
}

// NewTransferEvent func
func NewTransferEvent(id string, targetId string, amt int) TransferEvent {
	event := new(TransferEvent)
	event.Type = "TransferEvent"
	event.AccID = id
	event.Amount = amt
	event.TargetID = targetId
	return *event
}
