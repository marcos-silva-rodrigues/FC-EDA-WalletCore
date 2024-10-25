package createtransaction

import (
	"context"
	"testing"

	"github.com/marcos-silva-rodrigues/wallet-ms/internal/entity"
	"github.com/marcos-silva-rodrigues/wallet-ms/internal/event"
	"github.com/marcos-silva-rodrigues/wallet-ms/internal/usecase/mocks"
	"github.com/marcos-silva-rodrigues/wallet-ms/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransactionUseCaseExecute(t *testing.T) {
	client1, _ := entity.NewClient("client1", "j@j.com")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entity.NewClient("client2", "j@j2.com")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000)

	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	inputDto := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	}

	dispatcher := events.NewEventDispatcher()
	eventTransaction := event.NewTransactionCreated()
	evetnBalance := event.NewBalanceUpdated()

	ctx := context.Background()
	uc := NewCreateTransactionUseCase(dispatcher, eventTransaction, evetnBalance, mockUow)

	output, err := uc.Execute(ctx, inputDto)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)
}
