package gateway

import "github.com/marcos-silva-rodrigues/wallet-ms/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
