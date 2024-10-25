package handler

import (
	"fmt"
	"sync"

	"github.com/marcos-silva-rodrigues/wallet-ms/pkg/events"
	"github.com/marcos-silva-rodrigues/wallet-ms/pkg/kafka"
)

type BalanceUpdatedKakfaHandler struct {
	Kafka *kafka.Producer
}

func NewBalanceUpdatedKakfaHandler(kafka *kafka.Producer) *BalanceUpdatedKakfaHandler {
	return &BalanceUpdatedKakfaHandler{
		Kafka: kafka,
	}
}

func (h *BalanceUpdatedKakfaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()

	h.Kafka.Publish(message, nil, "balances")
	fmt.Println("BalanceUpdatedHandler: ", message.GetPayload())
}
