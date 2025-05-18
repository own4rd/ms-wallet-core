package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/own4rd/ms-wallet-core/internal/database"
	"github.com/own4rd/ms-wallet-core/internal/event"
	"github.com/own4rd/ms-wallet-core/internal/usecase/create_account"
	"github.com/own4rd/ms-wallet-core/internal/usecase/create_client"
	"github.com/own4rd/ms-wallet-core/internal/usecase/create_transaction"
	"github.com/own4rd/ms-wallet-core/internal/web"
	"github.com/own4rd/ms-wallet-core/internal/web/webserver"
	"github.com/own4rd/ms-wallet-core/pkg/events"
	"github.com/own4rd/ms-wallet-core/pkg/uow"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", "root", "root", "localhost", 3306, "testdb"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreateEvent := event.NewTransactionCreated()
	// eventDispatcher.Register("TransactionCreated", handler)
	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(uow, *eventDispatcher, transactionCreateEvent)

	webserver := webserver.NewWebServer(":3000")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandlers("/clients", clientHandler.CreateClient)
	webserver.AddHandlers("/accounts", accountHandler.CreateAccount)
	webserver.AddHandlers("/transactions", transactionHandler.CreateTransaction)

	webserver.Start()
}
