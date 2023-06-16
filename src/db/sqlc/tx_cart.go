package db

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

// Add To Cart Section

type AddToCartTxParam struct {
	CartID    uuid.UUID `json:"cart_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int32     `json:"quality"`
}

type AddToCartTxResult struct {
	Total    float64
	CartItem CartItem
}

// AddToCartTx function uses
func (store *SQLStore) AddToCartTx(ctx context.Context, arg AddToCartTxParam) (AddToCartTxResult, error) {
	var result AddToCartTxResult

	err := store.execTx(ctx, func(queries *Queries) error {
		cartItemId, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		product, err := queries.GetProductDetail(ctx, arg.ProductID)
		if err != nil {
			return err
		}

		if arg.Quantity <= 0 {
			return fmt.Errorf("Insuffice quantity. Cannot be less than 1.")
		}

		if product.Quantity.Int32 < arg.Quantity {
			return fmt.Errorf("Insuffice quantity. Only %v available", product.Quantity.Int32)
		}

		result.CartItem, err = queries.AddToCart(ctx, AddToCartParams{
			ID:        cartItemId,
			CartID:    arg.CartID,
			ProductID: arg.ProductID,
			Quantity:  arg.Quantity,
		})
		if err != nil {
			return err
		}

		result.Total, err = queries.GetTotal(ctx, result.CartItem.CartID)

		return err
	})

	return result, err
}

// Remove From Cart Section

type RemoveFromCartTxParam struct {
	CartItemID uuid.UUID `json:"cart_item_id"`
}

type RemoveFromCartTxResult struct {
	Total float64
}

func (store *SQLStore) RemoveFromCartTx(ctx context.Context, arg RemoveFromCartTxParam) (RemoveFromCartTxResult, error) {
	var result RemoveFromCartTxResult

	err := store.execTx(ctx, func(queries *Queries) error {
		var err error

		cartItem, err := queries.GetCartItemDetail(ctx, arg.CartItemID)
		if err != nil {
			return err
		}

		err = queries.RemoveItem(ctx, arg.CartItemID)
		if err != nil {
			return err
		}

		result.Total, err = queries.GetTotal(ctx, cartItem.CartID)

		return err
	})

	return result, err
}
