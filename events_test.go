package main_test

import (
	. "github.com/ayoo/eventsourcing-go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Event", func() {

	var (
		accId  = "AB928323-2837232832-3283232-232"
		amount = 2000
	)

	Describe("NewCreateAccountEvent", func() {
		It("can create a create account event", func() {
			name := "John Smith"

			event := NewCreateAccountEvent(name)

			Expect(event.AccName).To(Equal(name))
			Expect(event.AccID).NotTo(BeNil())
			Expect(event.Type).To(Equal("CreateEvent"))
		})
	})
	Describe("NewDepositEvent", func() {
		It("can create a deposit event", func() {
			event := NewDepositEvent(accId, amount)

			Expect(event.AccID).To(Equal(accId))
			Expect(event.Amount).To(Equal(amount))
			Expect(event.Type).To(Equal("DepositEvent"))
		})
	})

	Describe("NewWithdrawEvent", func() {
		It("can create a withdrawal event", func() {
			event := NewWithdrawEvent(accId, amount)

			Expect(event.AccID).To(Equal(accId))
			Expect(event.Amount).To(Equal(amount))
			Expect(event.Type).To(Equal("WithdrawEvent"))
		})
	})

	Describe("NewTransferEvent", func() {
		It("can create a transfer event", func() {
			event := NewTransferEvent(accId, "T0", amount)

			Expect(event.AccID).To(Equal(accId))
			Expect(event.Amount).To(Equal(amount))
			Expect(event.TargetID).To(Equal("T0"))
			Expect(event.Type).To(Equal("TransferEvent"))
		})
	})
})
