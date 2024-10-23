package gateway

import "github.com/marcos-silva-rodrigues/wallet-ms/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
