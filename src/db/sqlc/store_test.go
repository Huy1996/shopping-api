package db

import (
	"context"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"shopping-cart/src/util"
	"testing"
)

func TestCreateUserTx(t *testing.T) {
	store := NewStore(testDB)
	randomUsername := util.RandomName()
	randomEmail := util.RandomEmail()

	ctx := context.Background()
	// Successful case
	result, err := store.CreateUserTx(ctx, CreateUserTxParams{
		Username:       randomUsername,
		HashedPassword: util.RandomString(6),
		Email:          randomEmail,
		PhoneNumber:    util.RandomPhoneNumber(),
		FirstName:      util.RandomName(),
		LastName:       util.RandomName(),
		MiddleName:     util.RandomName(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, result.UserCredential.ID, result.UserInfo.UserID)

	// Duplicate username
	result2, err := store.CreateUserTx(ctx, CreateUserTxParams{
		Username:       randomUsername,
		HashedPassword: util.RandomString(6),
		Email:          util.RandomEmail(),
		PhoneNumber:    util.RandomPhoneNumber(),
		FirstName:      util.RandomName(),
		LastName:       util.RandomName(),
		MiddleName:     util.RandomName(),
	})
	require.Error(t, err)
	require.Empty(t, result2)

	require.Equal(t, err.(*pq.Error).Code.Name(), "unique_violation")

	// Duplicate Email
	result3, err := store.CreateUserTx(ctx, CreateUserTxParams{
		Username:       util.RandomName(),
		HashedPassword: util.RandomString(6),
		Email:          randomEmail,
		PhoneNumber:    util.RandomPhoneNumber(),
		FirstName:      util.RandomName(),
		LastName:       util.RandomName(),
		MiddleName:     util.RandomName(),
	})
	require.Error(t, err)
	require.Empty(t, result3)

	require.Equal(t, err.(*pq.Error).Code.Name(), "unique_violation")
}
