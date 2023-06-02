// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateAddressBook(ctx context.Context, arg CreateAddressBookParams) (AddressBook, error)
	CreateUserCredential(ctx context.Context, arg CreateUserCredentialParams) (UserCredential, error)
	CreateUserInfo(ctx context.Context, arg CreateUserInfoParams) (UserInfo, error)
	GetListAddresses(ctx context.Context, arg GetListAddressesParams) ([]AddressBook, error)
	GetNumberAddresses(ctx context.Context, owner uuid.UUID) (int64, error)
	GetUserCredential(ctx context.Context, username string) (UserCredential, error)
	GetUserInfoByID(ctx context.Context, id uuid.UUID) (UserInfo, error)
	GetUserInfoByUserID(ctx context.Context, userID uuid.UUID) (UserInfo, error)
}

var _ Querier = (*Queries)(nil)