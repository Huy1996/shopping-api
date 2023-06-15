package db

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

type AddToCartTxParam struct {
	CartID    uuid.UUID `json:"cart_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int32     `json:"quality"`
}

type AddToCartTxResult struct {
	Cart     UserCart
	CartItem CartItem
}

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

		if product.Quantity.Int32 < arg.Quantity {
			return fmt.Errorf("insuffice quantity. Only %v available", product.Quantity.Int32)
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

		result.Cart, err = queries.UpdateTotal(ctx, UpdateTotalParams{
			ID:     arg.CartID,
			Amount: float64(arg.Quantity) * (product.Price * product.DiscountPercent.Float64 / 100),
		})

		return err
	})

	return result, err
}
