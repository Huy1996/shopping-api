package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"shopping-cart/src/util"
	"testing"
	"time"
)

func CreateRandomUserCredential(t *testing.T) UserCredential {
	hashedPassword, err := util.HashPassword(util.RandomString(10))
	require.NoError(t, err)

	id, err := uuid.NewRandom()
	require.NoError(t, err)
	require.NotEmpty(t, id)

	arg := CreateUserCredentialParams{
		ID:             id,
		Username:       util.RandomName(),
		HashedPassword: hashedPassword,
		Email:          util.RandomEmail(),
	}

	userCredential, err := testQueries.CreateUserCredential(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userCredential)

	require.Equal(t, arg.ID, userCredential.ID)
	require.Equal(t, arg.Email, userCredential.Email)
	require.Equal(t, arg.Username, userCredential.Username)
	require.Equal(t, arg.HashedPassword, userCredential.HashedPassword)

	require.False(t, userCredential.IsAdmin)
	require.True(t, userCredential.PasswordChangedAt.IsZero())
	require.NotZero(t, userCredential.CreatedAt)

	return userCredential
}

func CreateRandomUserInfo(t *testing.T, uc UserCredential) UserInfo {
	id, err := uuid.NewRandom()
	require.NoError(t, err)
	require.NotEmpty(t, id)

	arg := CreateUserInfoParams{
		ID:          id,
		UserID:      uc.ID,
		PhoneNumber: util.RandomPhoneNumber(),
		FirstName:   util.RandomName(),
		LastName:    util.RandomName(),
		MiddleName:  util.RandomName(),
	}

	userInfo, err := testQueries.CreateUserInfo(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userInfo)

	require.Equal(t, arg.ID, userInfo.ID)
	require.Equal(t, arg.UserID, userInfo.UserID)
	require.Equal(t, arg.PhoneNumber, userInfo.PhoneNumber)
	require.Equal(t, arg.FirstName, userInfo.FirstName)
	require.Equal(t, arg.MiddleName, userInfo.MiddleName)

	require.True(t, userInfo.UpdatedAt.IsZero())
	require.NotZero(t, userInfo.CreatedAt)

	return userInfo
}

func CreateRandomUserAddress(t *testing.T, ui UserInfo) AddressBook {
	id, err := uuid.NewRandom()
	require.NoError(t, err)
	require.NotEmpty(t, id)

	arg := CreateAddressBookParams{
		ID:          id,
		Owner:       ui.ID,
		AddressName: util.RandomName(),
		Address:     util.RandomName(),
		City:        util.RandomCity(),
		State:       util.RandomState(),
		Zipcode:     int32(util.RandomInt(0, 99999)),
	}

	addressBook, err := testQueries.CreateAddressBook(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, addressBook)

	require.Equal(t, arg.ID, addressBook.ID)
	require.Equal(t, arg.Owner, addressBook.Owner)
	require.Equal(t, arg.AddressName, addressBook.AddressName)
	require.Equal(t, arg.Address, addressBook.Address)
	require.Equal(t, arg.City, addressBook.City)
	require.Equal(t, arg.State, addressBook.State)
	require.Equal(t, arg.Zipcode, addressBook.Zipcode)

	require.NotZero(t, addressBook.AddedAt)
	return addressBook
}

func TestGetUserCredential(t *testing.T) {
	userCredential1 := CreateRandomUserCredential(t)
	userCredential2, err := testQueries.GetUserCredential(context.Background(), userCredential1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, userCredential2)

	require.Equal(t, userCredential1.ID, userCredential2.ID)
	require.Equal(t, userCredential1.Username, userCredential2.Username)
	require.Equal(t, userCredential1.HashedPassword, userCredential2.HashedPassword)
	require.Equal(t, userCredential1.Email, userCredential2.Email)
	require.Equal(t, userCredential1.IsAdmin, userCredential2.IsAdmin)

	require.WithinDuration(t, userCredential1.PasswordChangedAt, userCredential2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, userCredential1.CreatedAt, userCredential2.CreatedAt, time.Second)
}

func TestGetUserInfo(t *testing.T) {
	userCredential := CreateRandomUserCredential(t)
	userInfo1 := CreateRandomUserInfo(t, userCredential)
	userInfo2, err := testQueries.GetUserInfoByID(context.Background(), userInfo1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, userInfo2)

	require.Equal(t, userInfo1.ID, userInfo2.ID)
	require.Equal(t, userInfo1.UserID, userInfo2.UserID)
	require.Equal(t, userInfo1.FirstName, userInfo2.FirstName)
	require.Equal(t, userInfo1.MiddleName, userInfo2.MiddleName)
	require.Equal(t, userInfo1.LastName, userInfo2.LastName)
	require.Equal(t, userInfo1.PhoneNumber, userInfo2.PhoneNumber)

	require.WithinDuration(t, userInfo1.UpdatedAt, userInfo2.UpdatedAt, time.Second)
	require.WithinDuration(t, userInfo1.CreatedAt, userInfo2.CreatedAt, time.Second)

	userInfo3, err := testQueries.GetUserInfoByUserID(context.Background(), userCredential.ID)
	require.NoError(t, err)
	require.NotEmpty(t, userInfo3)

	require.Equal(t, userInfo3.ID, userInfo2.ID)
	require.Equal(t, userInfo3.UserID, userInfo2.UserID)
	require.Equal(t, userInfo3.FirstName, userInfo2.FirstName)
	require.Equal(t, userInfo3.MiddleName, userInfo2.MiddleName)
	require.Equal(t, userInfo3.LastName, userInfo2.LastName)
	require.Equal(t, userInfo3.PhoneNumber, userInfo2.PhoneNumber)

	require.WithinDuration(t, userInfo3.UpdatedAt, userInfo2.UpdatedAt, time.Second)
	require.WithinDuration(t, userInfo3.CreatedAt, userInfo2.CreatedAt, time.Second)
}

func TestGetUserAddressBook(t *testing.T) {
	userCredential := CreateRandomUserCredential(t)
	userInfo := CreateRandomUserInfo(t, userCredential)
	userAddress1 := CreateRandomUserAddress(t, userInfo)
	userAddress2, err := testQueries.GetAddress(context.Background(), userAddress1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, userAddress2)

	require.Equal(t, userAddress1.ID, userAddress2.ID)
	require.Equal(t, userAddress1.Owner, userAddress2.Owner)
	require.Equal(t, userAddress1.AddressName, userAddress2.AddressName)
	require.Equal(t, userAddress1.Address, userAddress2.Address)
	require.Equal(t, userAddress1.City, userAddress2.City)
	require.Equal(t, userAddress1.Zipcode, userAddress2.Zipcode)

	require.WithinDuration(t, userAddress1.AddedAt, userAddress2.AddedAt, time.Second)
}

func TestGetListAddresses(t *testing.T) {
	userCredential := CreateRandomUserCredential(t)
	userInfo := CreateRandomUserInfo(t, userCredential)

	var lastAddress AddressBook
	const numAdr = 10
	for i := 0; i < numAdr; i++ {
		lastAddress = CreateRandomUserAddress(t, userInfo)
	}

	arg := GetListAddressesParams{
		Owner:  userInfo.ID,
		Limit:  5,
		Offset: 0,
	}

	addressList, err := testQueries.GetListAddresses(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, addressList)

	require.Equal(t, len(addressList), int(arg.Limit))

	for _, address := range addressList {
		require.NotEmpty(t, address)
		require.Equal(t, address.Owner, lastAddress.Owner)
	}

	count, err := testQueries.GetNumberAddresses(context.Background(), userInfo.ID)
	require.NoError(t, err)
	require.NotEmpty(t, count)

	require.Equal(t, int64(numAdr), count)

}
