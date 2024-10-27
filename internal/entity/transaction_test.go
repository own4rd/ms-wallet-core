package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@j")
	client2, _ := NewClient("John Doe2", "j2@j")
	accountFrom := NewAccount(client1)
	accountTo := NewAccount(client2)

	accountFrom.Credit(1000)
	accountTo.Credit(1000)
	transaction, err := NewTransaction(accountFrom, accountTo, 500)

	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 1500.0, accountTo.Balance)
	assert.Equal(t, 500.0, accountFrom.Balance)
}

func TestCreateInsuficientFundsTransaction(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@j")
	client2, _ := NewClient("John Doe2", "j2@j")
	accountFrom := NewAccount(client1)
	accountTo := NewAccount(client2)

	accountFrom.Credit(1000)
	accountTo.Credit(1000)
	transaction, err := NewTransaction(accountFrom, accountTo, 1500)

	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Error(t, err, "Insufficient funds")
	assert.Equal(t, 1000.0, accountTo.Balance)
	assert.Equal(t, 1000.0, accountFrom.Balance)
}
