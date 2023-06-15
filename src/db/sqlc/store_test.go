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

func TestAddToCartTx(t *testing.T) {
	store := NewStore(testDB)
	cart := CreateCart(t)
	product := CreateProduct(t)

	ctx := context.Background()

	productDetail, err := store.GetProductDetail(ctx, product.ID)
	require.NoError(t, err)
	require.NotEmpty(t, productDetail)

	// successfully case
	result, err := store.AddToCartTx(ctx, AddToCartTxParam{
		CartID:    cart.ID,
		ProductID: productDetail.ID,
		Quantity:  productDetail.Quantity.Int32,
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
	price := float64(productDetail.Quantity.Int32) * productDetail.Price * (1 + productDetail.DiscountPercent.Float64/100)
	require.Equal(t, price, result.Cart.Total)
	require.Equal(t, cart.ID, result.CartItem.CartID)
	require.Equal(t, productDetail.ID, result.CartItem.ProductID)
	require.Equal(t, productDetail.Quantity.Int32, result.CartItem.Quantity)

	// Negative Qty
	result1, err := store.AddToCartTx(ctx, AddToCartTxParam{
		CartID:    cart.ID,
		ProductID: productDetail.ID,
		Quantity:  0,
	})
	require.Error(t, err)
	require.Empty(t, result1)

	result2, err := store.AddToCartTx(ctx, AddToCartTxParam{
		CartID:    cart.ID,
		ProductID: productDetail.ID,
		Quantity:  -1,
	})
	require.Error(t, err)
	require.Empty(t, result2)

	// Over product qty
	result3, err := store.AddToCartTx(ctx, AddToCartTxParam{
		CartID:    cart.ID,
		ProductID: productDetail.ID,
		Quantity:  productDetail.Quantity.Int32 + 1,
	})
	require.Error(t, err)
	require.Empty(t, result3)

	// Add Multiple Product
	var total float64
	cart = CreateCart(t)
	for i := 0; i < 10; i++ {
		product := CreateProduct(t)
		productDetail, err := store.GetProductDetail(ctx, product.ID)
		require.NoError(t, err)
		require.NotEmpty(t, productDetail)

		result, err := store.AddToCartTx(ctx, AddToCartTxParam{
			CartID:    cart.ID,
			ProductID: productDetail.ID,
			Quantity:  productDetail.Quantity.Int32,
		})
		total += float64(result.CartItem.Quantity) * productDetail.Price * (1 + productDetail.DiscountPercent.Float64/100)

		require.NoError(t, err)
		require.NotEmpty(t, result)
		require.Equal(t, total, result.Cart.Total)
		require.Equal(t, cart.ID, result.CartItem.CartID)
		require.Equal(t, productDetail.ID, result.CartItem.ProductID)
		require.Equal(t, productDetail.Quantity.Int32, result.CartItem.Quantity)
	}
}
