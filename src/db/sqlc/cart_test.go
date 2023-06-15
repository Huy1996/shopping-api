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
	require.Equal(t, cart.Total, float64(0))

	return cart
}

func CreateProduct(t *testing.T) Product {
	product := CreateRandomProduct(
		t,
		CreateRandomInventory(t),
		CreateRandomCategory(t),
	)

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

	arg := AddToCartParams{
		ID:        id,
		CartID:    cart.ID,
		ProductID: product.ID,
		Quantity:  int32(util.RandomInt(1, 100)),
	}

	cartItem, err := testQueries.AddToCart(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cartItem)

	require.Equal(t, arg.ID, cartItem.ID)
	require.Equal(t, arg.CartID, cartItem.CartID)
	require.Equal(t, arg.Quantity, cartItem.Quantity)
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

func TestUpdateTotal(t *testing.T) {
	cart := CreateCart(t)
	amount := float64(50.0)

	arg := UpdateTotalParams{
		ID:     cart.ID,
		Amount: amount,
	}

	updatedCart, err := testQueries.UpdateTotal(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedCart)

	require.Equal(t, cart.ID, updatedCart.ID)
	require.Equal(t, cart.Owner, updatedCart.Owner)
	require.Equal(t, amount, updatedCart.Total)
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
