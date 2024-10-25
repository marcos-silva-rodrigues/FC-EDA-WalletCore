package handler

import (
	"fmt"
	"sync"

	"github.com/marcos-silva-rodrigues/wallet-ms/pkg/events"
	"github.com/marcos-silva-rodrigues/wallet-ms/pkg/kafka"
)

type TransactionCreatedKakfaHandler struct {
	Kafka *kafka.Producer
}

func NewTransactionCreatedKakfaHandler(kafka *kafka.Producer) *TransactionCreatedKakfaHandler {
	return &TransactionCreatedKakfaHandler{
		Kafka: kafka,
	}
}

func (h *TransactionCreatedKakfaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()

	h.Kafka.Publish(message, nil, "transactions")
	fmt.Println("TransactionCreatedHandler: ", message.GetPayload())
}
