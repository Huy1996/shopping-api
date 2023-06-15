// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: cart.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const addToCart = `-- name: AddToCart :one
INSERT INTO cart_item (
    id,
    cart_id,
    product_id,
    quantity
) VALUES (
    $1, $2, $3, $4
) RETURNING id, cart_id, product_id, quantity
`

type AddToCartParams struct {
	ID        uuid.UUID `json:"id"`
	CartID    uuid.UUID `json:"cart_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int32     `json:"quantity"`
}

func (q *Queries) AddToCart(ctx context.Context, arg AddToCartParams) (CartItem, error) {
	row := q.queryRow(ctx, q.addToCartStmt, addToCart,
		arg.ID,
		arg.CartID,
		arg.ProductID,
		arg.Quantity,
	)
	var i CartItem
	err := row.Scan(
		&i.ID,
		&i.CartID,
		&i.ProductID,
		&i.Quantity,
	)
	return i, err
}

const createCart = `-- name: CreateCart :one
INSERT INTO user_cart (
    id,
    owner
) VALUES (
    $1, $2
) RETURNING id, owner, total
`

type CreateCartParams struct {
	ID    uuid.UUID `json:"id"`
	Owner uuid.UUID `json:"owner"`
}

func (q *Queries) CreateCart(ctx context.Context, arg CreateCartParams) (UserCart, error) {
	row := q.queryRow(ctx, q.createCartStmt, createCart, arg.ID, arg.Owner)
	var i UserCart
	err := row.Scan(&i.ID, &i.Owner, &i.Total)
	return i, err
}

const deleteCart = `-- name: DeleteCart :exec
DELETE FROM user_cart
WHERE owner = $1
`

func (q *Queries) DeleteCart(ctx context.Context, owner uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteCartStmt, deleteCart, owner)
	return err
}

const getCartByID = `-- name: GetCartByID :one
SELECT id, owner, total FROM user_cart
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetCartByID(ctx context.Context, id uuid.UUID) (UserCart, error) {
	row := q.queryRow(ctx, q.getCartByIDStmt, getCartByID, id)
	var i UserCart
	err := row.Scan(&i.ID, &i.Owner, &i.Total)
	return i, err
}

const getCartProductList = `-- name: GetCartProductList :many
SELECT id, cart_id, product_id, quantity FROM cart_item
WHERE cart_id = $1
LIMIT $2
OFFSET $3
`

type GetCartProductListParams struct {
	CartID uuid.UUID `json:"cart_id"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

func (q *Queries) GetCartProductList(ctx context.Context, arg GetCartProductListParams) ([]CartItem, error) {
	rows, err := q.query(ctx, q.getCartProductListStmt, getCartProductList, arg.CartID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CartItem{}
	for rows.Next() {
		var i CartItem
		if err := rows.Scan(
			&i.ID,
			&i.CartID,
			&i.ProductID,
			&i.Quantity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeItem = `-- name: RemoveItem :exec
DELETE FROM cart_item
WHERE id = $1
`

func (q *Queries) RemoveItem(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.removeItemStmt, removeItem, id)
	return err
}

const updateCartItemQty = `-- name: UpdateCartItemQty :one
UPDATE cart_item
SET quantity = $1
WHERE id = $2
RETURNING id, cart_id, product_id, quantity
`

type UpdateCartItemQtyParams struct {
	Quantity int32     `json:"quantity"`
	ID       uuid.UUID `json:"id"`
}

func (q *Queries) UpdateCartItemQty(ctx context.Context, arg UpdateCartItemQtyParams) (CartItem, error) {
	row := q.queryRow(ctx, q.updateCartItemQtyStmt, updateCartItemQty, arg.Quantity, arg.ID)
	var i CartItem
	err := row.Scan(
		&i.ID,
		&i.CartID,
		&i.ProductID,
		&i.Quantity,
	)
	return i, err
}

const updateTotal = `-- name: UpdateTotal :one
UPDATE user_cart
SET total = total + $1
WHERE id = $2
RETURNING id, owner, total
`

type UpdateTotalParams struct {
	Amount float64   `json:"amount"`
	ID     uuid.UUID `json:"id"`
}

func (q *Queries) UpdateTotal(ctx context.Context, arg UpdateTotalParams) (UserCart, error) {
	row := q.queryRow(ctx, q.updateTotalStmt, updateTotal, arg.Amount, arg.ID)
	var i UserCart
	err := row.Scan(&i.ID, &i.Owner, &i.Total)
	return i, err
}
