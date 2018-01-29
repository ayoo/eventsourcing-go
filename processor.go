package main

import "errors"

type EventProcessor interface {
	Process() (*BankAccount, error)
}

func (e CreateEvent) Process() (*BankAccount, error) {
	return saveAccount(e.AccID, map[string]interface{}{
		"ID":      e.AccID,
		"Name":    e.AccName,
		"Balance": 0,
	})
}

func (e DepositEvent) Process() (*BankAccount, error) {
	acc, err := GetAccount(e.AccID)
	if err != nil {
		return nil, err
	}

	newBalance := acc.Balance + e.Amount
	return saveAccount(e.AccID, map[string]interface{}{
		"Balance": newBalance,
	})
}

func (e WithdrawEvent) Process() (*BankAccount, error) {
	acc, err := GetAccount(e.AccID)
	if err != nil {
		return nil, err
	}

	if acc.Balance < e.Amount {
		return nil, errors.New("Insufficient amount")
	}

	newBalance := acc.Balance - e.Amount
	return saveAccount(e.AccID, map[string]interface{}{
		"Balance": newBalance,
	})
}

func (e TransferEvent) Process() (*BankAccount, error) {
	srcAcc, err := GetAccount(e.AccID)
	if err != nil {
		return nil, err
	}

	descAcc, err := GetAccount(e.TargetID)
	if err != nil {
		return nil, err
	}

	if srcAcc.Balance < e.Amount {
		return nil, errors.New("Insufficient amount")
	}

	srcAcc.Balance -= e.Amount
	descAcc.Balance += e.Amount

	_, err = saveAccount(descAcc.ID, map[string]interface{}{
		"Balance": descAcc.Balance,
	})
	if err != nil {
		return nil, err
	}

	return saveAccount(srcAcc.ID, map[string]interface{}{
		"Balance": srcAcc.Balance,
	})
}
