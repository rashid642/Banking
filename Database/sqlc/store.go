package Database

import (
	"context"
	"database/sql"
	"fmt"
)

// store function provide all the function to execute db and transaction
type Store struct {
	*Queries
	db *sql.DB
}

// Create a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx execute a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)

	err = fn(q)
	if err != nil {
		if rollBackError := tx.Rollback(); rollBackError != nil {
			return fmt.Errorf("Transaction Error: %v, RollbackError: %v", err, rollBackError)
		}
		return err
	}

	return tx.Commit()
}

// this contains all the input paramters to transfer money from to account
type TransferTxParams struct {
	FromAccountId int64 `json:"from_account_id"`
	ToAccountId   int64 `json:"to_account_id"`
	Amount        int64         `json:"amount"`
}

// Transafer transaction result
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// Transfer perform money from one account to other
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountId,
			ToAccountID:   arg.ToAccountId,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntires(ctx, CreateEntiresParams{
			AccountID: arg.FromAccountId,
			Amount: -arg.Amount,
		})
		if err != nil {
			return err
		}
		
		result.ToEntry, err = q.CreateEntires(ctx, CreateEntiresParams{
			AccountID: arg.ToAccountId,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}
		
		if arg.FromAccountId < arg.ToAccountId {
			result.FromAccount, result.ToAccount, err = addMoneyToAccount(ctx, q, arg.FromAccountId, -arg.Amount, arg.ToAccountId, arg.Amount)
		}else{
			result.ToAccount, result.FromAccount, err = addMoneyToAccount(ctx, q, arg.ToAccountId, arg.Amount, arg.FromAccountId, -arg.Amount)
		}

		return err 
	})

	return result, err
}

func addMoneyToAccount(
	ctx context.Context, 
	q *Queries, 
	accountID1 int64,
	amount1 int64,
	accountID2 int64, 
	amount2 int64,
) (account1 Account, account2 Account, err error) {
	account1, err = q.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
		ID: accountID1,
		Amount: amount1,
	})
	if err != nil {
		return account1, account2, err 
	}

	account1, err = q.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
		ID: accountID2,
		Amount: amount2,
	})
	return account1, account2, err 
}
