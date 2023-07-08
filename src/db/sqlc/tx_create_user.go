package db

import (
	"context"
	"github.com/google/uuid"
)

type CreateUserTxParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	MiddleName     string `json:"middle_name"`
}

type CreateUserTxResult struct {
	UserCredential UserCredential
	UserInfo       UserInfo
	UserCart       UserCart
}

func (store *SQLStore) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult

	err := store.execTx(ctx, func(queries *Queries) error {
		var err error

		credentialId, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		result.UserCredential, err = queries.CreateUserCredential(ctx, CreateUserCredentialParams{
			ID:             credentialId,
			Username:       arg.Username,
			HashedPassword: arg.HashedPassword,
			Email:          arg.Email,
		})
		if err != nil {
			return err
		}

		infoId, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		result.UserInfo, err = queries.CreateUserInfo(ctx, CreateUserInfoParams{
			ID:          infoId,
			UserID:      credentialId,
			PhoneNumber: arg.PhoneNumber,
			FirstName:   arg.FirstName,
			LastName:    arg.LastName,
			MiddleName:  arg.MiddleName,
		})
		if err != nil {
			return err
		}

		cartId, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		result.UserCart, err = queries.CreateCart(ctx, CreateCartParams{
			ID:    cartId,
			Owner: infoId,
		})
		return err
	})

	return result, err
}
