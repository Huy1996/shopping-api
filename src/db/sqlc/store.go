package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transaction
type Store interface {
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	//AddToCartTx(ctx context.Context, arg AddToCartTxParam) (AddToCartTxResult, error)
	Querier
}

// SQLStore provides all function to execute SQL queries and transaction
type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(queries *Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	queries := New(tx)
	err = fn(queries)
	if err != nil {
		// Try to rollback transaction when failed
		if rbErr := tx.Rollback(); rbErr != nil {
			// Return initial error and rollback error when failed to rollback
			return fmt.Errorf("tx err: %v, rb err %v", err, rbErr)
		}
		// Return initial error
		return err
	}
	return tx.Commit()
}
