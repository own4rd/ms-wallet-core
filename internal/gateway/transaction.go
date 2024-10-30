package gateway

import "github.com/own4rd/ms-wallet-core/internal/entity"

type TransactionGateway interface {
	Create(transaction entity.Transaction) error
}
