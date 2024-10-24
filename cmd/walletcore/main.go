package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/marcos-silva-rodrigues/wallet-ms/internal/database"
	"github.com/marcos-silva-rodrigues/wallet-ms/internal/event"
	createaccount "github.com/marcos-silva-rodrigues/wallet-ms/internal/usecase/create_account"
	createclient "github.com/marcos-silva-rodrigues/wallet-ms/internal/usecase/create_client"
	createtransaction "github.com/marcos-silva-rodrigues/wallet-ms/internal/usecase/create_transaction"
	"github.com/marcos-silva-rodrigues/wallet-ms/internal/web"
	"github.com/marcos-silva-rodrigues/wallet-ms/internal/web/webserver"
	"github.com/marcos-silva-rodrigues/wallet-ms/pkg/events"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionEventCreated := event.NewTransactionCreated()
	// eventDispatcher.Register("TransactionCreated", handler)

	clientDB := database.NewClientDB(db)
	accountDB := database.NewAccountDB(db)
	transactionDB := database.NewTransactionDB(db)

	createClientUseCase := createclient.NewCreateClientUseCase(clientDB)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDB, clientDB)

	createTransactionUseCase := createtransaction.NewCreateTransactionUseCase(
		transactionDB, accountDB, eventDispatcher, transactionEventCreated)

	webserver := webserver.NewWebServer(":9000")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	webserver.Start()
}
