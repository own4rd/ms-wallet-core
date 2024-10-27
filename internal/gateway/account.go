package gateway

import "github.com/own4rd/ms-wallet-core/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	findByID(id string) (*entity.Account, error)
}
