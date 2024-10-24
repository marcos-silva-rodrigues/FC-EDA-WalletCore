package createtransaction

import (
	"context"

	"github.com/marcos-silva-rodrigues/wallet-ms/internal/entity"
	"github.com/marcos-silva-rodrigues/wallet-ms/internal/gateway"
	"github.com/marcos-silva-rodrigues/wallet-ms/pkg/events"
	"github.com/marcos-silva-rodrigues/wallet-ms/pkg/uow"
)

type CreateTransactionInputDTO struct {
	AccountIDFrom string `json:"account_id_from"`
	AccountIDTo   string `json:"account_id_to"`
	Amount        float64
}

type CreateTransactionOutputDTO struct {
	ID            string
	AccountIDFrom string `json:"account_id_from"`
	AccountIDTo   string `json:"account_id_to"`
	Amount        float64
}

type CreateTransactionUseCase struct {
	Uow                uow.UowInterface
	EventDispatcher    events.EventDispatcherInterface
	TransactionCreated events.EventInterface
}

func NewCreateTransactionUseCase(
	eventDispatcher events.EventDispatcherInterface,
	transactionCreated events.EventInterface,
	uow uow.UowInterface,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		Uow:                uow,
		EventDispatcher:    eventDispatcher,
		TransactionCreated: transactionCreated,
	}
}

func (uc *CreateTransactionUseCase) Execute(
	ctx context.Context,
	input CreateTransactionInputDTO,
) (*CreateTransactionOutputDTO, error) {
	output := &CreateTransactionOutputDTO{}

	err := uc.Uow.Do(ctx, func(_ *uow.Uow) error {
		accountGateway := uc.getAccountGateway(ctx)
		transactionGateway := uc.getTransactionGateway(ctx)
		accountFrom, err := accountGateway.FindByID(input.AccountIDFrom)

		if err != nil {
			return err
		}

		accountTo, err := accountGateway.FindByID(input.AccountIDTo)

		if err != nil {
			return err
		}

		transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
		if err != nil {
			return err
		}

		err = accountGateway.UpdateBalance(accountFrom)
		if err != nil {
			return err
		}

		err = accountGateway.UpdateBalance(accountTo)
		if err != nil {
			return err
		}

		err = transactionGateway.Create(transaction)

		if err != nil {
			return err
		}

		output = &CreateTransactionOutputDTO{
			ID: transaction.ID,
		}

		output.ID = transaction.ID
		output.AccountIDFrom = accountFrom.ID
		output.AccountIDTo = accountTo.ID
		output.Amount = input.Amount
		return nil
	})

	if err != nil {
		return nil, err
	}

	uc.TransactionCreated.SetPayload(output)
	uc.EventDispatcher.Dispatch(uc.TransactionCreated)

	return output, nil
}

func (uc *CreateTransactionUseCase) getAccountGateway(ctx context.Context) gateway.AccountGateway {
	repo, err := uc.Uow.GetRepository(ctx, "AccountGateway")

	if err != nil {
		panic(err)
	}

	return repo.(gateway.AccountGateway)
}

func (uc *CreateTransactionUseCase) getTransactionGateway(ctx context.Context) gateway.TransactionGateway {
	repo, err := uc.Uow.GetRepository(ctx, "TransactionGateway")

	if err != nil {
		panic(err)
	}

	return repo.(gateway.TransactionGateway)
}
