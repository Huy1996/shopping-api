package db

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
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

	var price float64
	if productDetail.DiscountActive.Bool {
		price = float64(productDetail.Quantity.Int32) * product.Price * (1 - productDetail.DiscountPercent.Float64/100)
	} else {
		price = float64(productDetail.Quantity.Int32) * product.Price
	}

	require.Equal(t, cart.ID, result.CartItem.CartID)
	require.Equal(t, productDetail.ID, result.CartItem.ProductID)
	require.Equal(t, productDetail.Quantity.Int32, result.CartItem.Quantity)
	require.True(t, util.WithinTolerance(price, result.Total, util.CurrencyTolerance))

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

		if productDetail.DiscountActive.Bool {
			price = float64(productDetail.Quantity.Int32) * product.Price * (1 - productDetail.DiscountPercent.Float64/100)
		} else {
			price = float64(productDetail.Quantity.Int32) * product.Price
		}

		total += price

		require.NoError(t, err)
		require.NotEmpty(t, result)
		require.Equal(t, cart.ID, result.CartItem.CartID)
		require.True(t, util.WithinTolerance(total, result.Total, util.CurrencyTolerance))
		require.Equal(t, productDetail.ID, result.CartItem.ProductID)
		require.Equal(t, productDetail.Quantity.Int32, result.CartItem.Quantity)
	}

	// Non-exist product ID
	randomID, err := uuid.NewRandom()
	require.NoError(t, err)
	require.NotEmpty(t, randomID)

	result4, err := store.AddToCartTx(ctx, AddToCartTxParam{
		CartID:    cart.ID,
		ProductID: randomID,
		Quantity:  1,
	})
	require.Error(t, err)
	require.Empty(t, result4)
	require.ErrorIs(t, sql.ErrNoRows, err)
}

func TestRemoveFromCartTx(t *testing.T) {
	store := NewStore(testDB)
	cart := CreateCart(t)

	ctx := context.Background()
	var lastProduct Product
	var lastCartId uuid.UUID
	var total float64

	for i := 0; i < 10; i++ {
		lastProduct = CreateProduct(t)

		productDetail, err := store.GetProductDetail(ctx, lastProduct.ID)
		require.NoError(t, err)
		require.NotEmpty(t, productDetail)

		result, err := store.AddToCartTx(ctx, AddToCartTxParam{
			CartID:    cart.ID,
			ProductID: productDetail.ID,
			Quantity:  productDetail.Quantity.Int32,
		})

		if productDetail.DiscountActive.Bool {
			total += float64(productDetail.Quantity.Int32) * lastProduct.Price * (1 - productDetail.DiscountPercent.Float64/100)
		} else {
			total += float64(productDetail.Quantity.Int32) * lastProduct.Price
		}
		lastCartId = result.CartItem.ID
		require.NoError(t, err)
		require.NotEmpty(t, result)
		require.Equal(t, cart.ID, result.CartItem.CartID)
		require.True(t, util.WithinTolerance(total, result.Total, util.CurrencyTolerance))
		require.Equal(t, productDetail.ID, result.CartItem.ProductID)
		require.Equal(t, productDetail.Quantity.Int32, result.CartItem.Quantity)
	}

	productDetail, err := store.GetProductDetail(ctx, lastProduct.ID)
	require.NoError(t, err)
	require.NotEmpty(t, productDetail)
	if productDetail.DiscountActive.Bool {
		total -= float64(productDetail.Quantity.Int32) * lastProduct.Price * (1 - productDetail.DiscountPercent.Float64/100)
	} else {
		total -= float64(productDetail.Quantity.Int32) * lastProduct.Price
	}

	// Successfully
	result, err := store.RemoveFromCartTx(ctx, RemoveFromCartTxParam{
		CartItemID: lastCartId,
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.True(t, util.WithinTolerance(total, result.Total, util.CurrencyTolerance))

	// Item not exist
	result1, err := store.RemoveFromCartTx(ctx, RemoveFromCartTxParam{
		CartItemID: lastCartId,
	})
	require.Error(t, err)
	require.Empty(t, result1)

	require.ErrorIs(t, sql.ErrNoRows, err)
}

func TestChangeQtyTx(t *testing.T) {
	store := NewStore(testDB)
	cart := CreateCart(t)
	product := CreateProduct(t)

	ctx := context.Background()

	productDetail, err := store.GetProductDetail(ctx, product.ID)
	require.NoError(t, err)
	require.NotEmpty(t, productDetail)

	addResult, err := store.AddToCartTx(ctx, AddToCartTxParam{
		CartID:    cart.ID,
		ProductID: productDetail.ID,
		Quantity:  productDetail.Quantity.Int32,
	})
	require.NoError(t, err)
	require.NotEmpty(t, addResult)

	// success
	result, err := store.ChangeQtyTx(ctx, ChangeQtyTxParam{
		CartItemID: addResult.CartItem.ID,
		Quantity:   productDetail.Quantity.Int32 - 1,
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
	var total float64

	if productDetail.DiscountActive.Bool {
		total = float64(productDetail.Quantity.Int32-1) * productDetail.Price * (1 - productDetail.DiscountPercent.Float64/100)
	} else {
		total = float64(productDetail.Quantity.Int32-1) * productDetail.Price
	}
	require.True(t, util.WithinTolerance(total, result.Total, util.CurrencyTolerance))

	// fail (over available qty)
	result1, err := store.ChangeQtyTx(ctx, ChangeQtyTxParam{
		CartItemID: addResult.CartItem.ID,
		Quantity:   productDetail.Quantity.Int32 + 1,
	})
	require.Error(t, err)
	require.Empty(t, result1)

	// fail (less than 1)
	result2, err := store.ChangeQtyTx(ctx, ChangeQtyTxParam{
		CartItemID: addResult.CartItem.ID,
		Quantity:   0,
	})
	require.Error(t, err)
	require.Empty(t, result2)

	// Error not found
	randomID, err := uuid.NewRandom()
	result3, err := store.ChangeQtyTx(ctx, ChangeQtyTxParam{
		CartItemID: randomID,
		Quantity:   0,
	})
	require.Error(t, err)
	require.Empty(t, result3)
	require.ErrorIs(t, sql.ErrNoRows, err)
}
