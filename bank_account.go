package main

import (
	"errors"
	"strconv"
)

type BankAccount struct {
	ID      string
	Name    string
	Balance int
}

// GetAccount - fetch bank account by id from Redis
// returns BankAccount, error
func GetAccount(id string) (*BankAccount, error) {
	cmd := Redis.HGetAll(id)
	if err := cmd.Err(); err != nil {
		return nil, err
	}

	data := cmd.Val()
	if len(data) == 0 {
		return nil, nil
	}

	return ToAccount(data)
}

// saveAccount - save a map of account data to Redis
func saveAccount(id string, data map[string]interface{}) (*BankAccount, error) {
	cmd := Redis.HMSet(id, data)
	if err := cmd.Err(); err != nil {
		return nil, err
	}
	return GetAccount(id)
}

// ToAccount - converts a data map to BankAccount type
func ToAccount(m map[string]string) (*BankAccount, error) {
	if _, ok := m["ID"]; !ok {
		return nil, errors.New("Missing account ID")
	}

	balance, err := strconv.Atoi(m["Balance"])
	if err != nil {
		return nil, err
	}

	return &BankAccount{
		ID:      m["ID"],
		Name:    m["Name"],
		Balance: balance,
	}, nil
}
