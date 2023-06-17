package db

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"shopping-cart/src/util"
	"testing"
)

func CreateCart(t *testing.T) UserCart {
	userCredential := CreateRandomUserCredential(t)
	userInfo := CreateRandomUserInfo(t, userCredential)

	id, err := uuid.NewRandom()
	require.NoError(t, err)
	require.NotEmpty(t, id)

	arg := CreateCartParams{
		ID:    id,
		Owner: userInfo.ID,
	}

	cart, err := testQueries.CreateCart(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cart)

	require.Equal(t, arg.ID, cart.ID)
	require.Equal(t, arg.Owner, cart.Owner)

	return cart
}

func CreateProduct(t *testing.T) Product {
	product := CreateRandomProduct(
		t,
		CreateRandomInventory(t),
		CreateRandomCategory(t),
	)

	productDiscount := CreateRandomDiscount(t)
	product, err := testQueries.AddDiscount(context.Background(), AddDiscountParams{
		ID: product.ID,
		DiscountID: uuid.NullUUID{
			UUID:  productDiscount.ID,
			Valid: true,
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, product)

	return product
}

func CreateCartItem(t *testing.T, cart UserCart, item Product, qty int32) CartItem {
	id, err := uuid.NewRandom()
	require.NoError(t, err)
	require.NotEmpty(t, id)

	arg := AddToCartParams{
		ID:        id,
		CartID:    cart.ID,
		ProductID: item.ID,
		Quantity:  qty,
	}

	cartItem, err := testQueries.AddToCart(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cartItem)

	require.Equal(t, id, cartItem.ID)
	require.Equal(t, cart.ID, cartItem.CartID)
	require.Equal(t, item.ID, cartItem.ProductID)
	require.Equal(t, qty, cartItem.Quantity)

	return cartItem
}

func TestAddToCart(t *testing.T) {
	cart := CreateCart(t)
	product := CreateProduct(t)

	id, err := uuid.NewRandom()
	require.NoError(t, err)
	require.NotEmpty(t, id)
	ctx := context.Background()
	arg := AddToCartParams{
		ID:        id,
		CartID:    cart.ID,
		ProductID: product.ID,
		Quantity:  int32(util.RandomInt(1, 100)),
	}

	cartItem, err := testQueries.AddToCart(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, cartItem)

	require.Equal(t, arg.ID, cartItem.ID)
	require.Equal(t, arg.CartID, cartItem.CartID)
	require.Equal(t, arg.Quantity, cartItem.Quantity)

	total, err := testQueries.GetTotal(ctx, cartItem.CartID)
	require.NoError(t, err)
	require.NotEmpty(t, total)

	cartItemDetail, err := testQueries.GetCartItemDetail(ctx, cartItem.ID)
	require.NoError(t, err)
	require.NotEmpty(t, cartItemDetail)

	var expectedTotal float64
	if cartItemDetail.DiscountActive.Bool {
		expectedTotal = float64(cartItemDetail.Quantity) * cartItemDetail.Price.Float64 * (1 + cartItemDetail.DiscountPercent.Float64/100)
	} else {
		expectedTotal = float64(cartItemDetail.Quantity) * cartItemDetail.Price.Float64
	}
	require.True(t, util.WithinTolerance(expectedTotal, total, util.CurrencyTolerance))
	require.True(t, util.WithinTolerance(expectedTotal, cartItemDetail.Total, util.CurrencyTolerance))
}

func TestDeleteCart(t *testing.T) {
	cart := CreateCart(t)

	err := testQueries.DeleteCart(context.Background(), cart.Owner)
	require.NoError(t, err)

	deletedCart, err := testQueries.GetCartByID(context.Background(), cart.ID)
	require.Error(t, err)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, deletedCart)
}

func TestAddGetRemoveItem(t *testing.T) {
	cart := CreateCart(t)
	numProduct := 10
	var lastProduct Product
	for i := 0; i < numProduct; i++ {
		lastProduct = CreateProduct(t)
		_ = CreateCartItem(t, cart, lastProduct, 1)
	}

	cartProductList, err := testQueries.GetCartProductList(context.Background(), GetCartProductListParams{
		CartID: cart.ID,
		Limit:  int32(numProduct),
		Offset: 0,
	})
	require.NoError(t, err)
	require.NotEmpty(t, cartProductList)
	require.Equal(t, numProduct, len(cartProductList))

	inCart := false
	var id uuid.UUID
	for _, cartProduct := range cartProductList {
		if cartProduct.ProductID == lastProduct.ID {
			inCart = true
			id = cartProduct.ID
		}
	}
	require.True(t, inCart)

	err = testQueries.RemoveItem(context.Background(), id)
	require.NoError(t, err)

	updatedCartProductList, err := testQueries.GetCartProductList(context.Background(), GetCartProductListParams{
		CartID: cart.ID,
		Limit:  int32(numProduct),
		Offset: 0,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedCartProductList)
	require.Equal(t, numProduct-1, len(updatedCartProductList))

	for _, cartProduct := range updatedCartProductList {
		require.NotEqual(t, lastProduct.ID, cartProduct.ProductID)
	}
}

func TestUpdateItemQty(t *testing.T) {
	cart := CreateCart(t)
	product := CreateProduct(t)
	cartItem := CreateCartItem(t, cart, product, 1)
	newQty := int32(10)

	updatedCartItem, err := testQueries.UpdateCartItemQty(context.Background(), UpdateCartItemQtyParams{
		ID:       cartItem.ID,
		Quantity: newQty,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedCartItem)

	require.Equal(t, cartItem.ID, updatedCartItem.ID)
	require.Equal(t, cart.ID, updatedCartItem.CartID)
	require.Equal(t, product.ID, updatedCartItem.ProductID)
	require.Equal(t, newQty, updatedCartItem.Quantity)
}
