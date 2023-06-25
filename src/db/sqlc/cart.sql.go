// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: cart.sql

package db

import (
	"context"
	"database/sql"

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
) RETURNING id, owner
`

type CreateCartParams struct {
	ID    uuid.UUID `json:"id"`
	Owner uuid.UUID `json:"owner"`
}

func (q *Queries) CreateCart(ctx context.Context, arg CreateCartParams) (UserCart, error) {
	row := q.queryRow(ctx, q.createCartStmt, createCart, arg.ID, arg.Owner)
	var i UserCart
	err := row.Scan(&i.ID, &i.Owner)
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
SELECT id, owner FROM user_cart
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetCartByID(ctx context.Context, id uuid.UUID) (UserCart, error) {
	row := q.queryRow(ctx, q.getCartByIDStmt, getCartByID, id)
	var i UserCart
	err := row.Scan(&i.ID, &i.Owner)
	return i, err
}

const getCartItemDetail = `-- name: GetCartItemDetail :one
SELECT
	cart_item.id,
	cart_item.cart_id,
	cart_item.quantity,
	product.price AS price,
	product_inventory.quantity AS qty_in_stock,
	float8(
	CASE
		WHEN product_discount.active THEN (product.price * (1 - product_discount.discount_percent / 100) * cart_item.quantity)
		ELSE (product.price * cart_item.quantity )
	END) AS total,
	product_discount.discount_percent AS discount_percent,
	product_discount.active AS discount_active
FROM cart_item
LEFT JOIN product
    ON cart_item.product_id = product.id
LEFT JOIN product_discount
    ON product.discount_id = product_discount.id
LEFT JOIN product_inventory
    ON product.inventory_id = product_inventory.id
WHERE cart_item.id = $1
`

type GetCartItemDetailRow struct {
	ID              uuid.UUID       `json:"id"`
	CartID          uuid.UUID       `json:"cart_id"`
	Quantity        int32           `json:"quantity"`
	Price           sql.NullFloat64 `json:"price"`
	QtyInStock      sql.NullInt32   `json:"qty_in_stock"`
	Total           float64         `json:"total"`
	DiscountPercent sql.NullFloat64 `json:"discount_percent"`
	DiscountActive  sql.NullBool    `json:"discount_active"`
}

func (q *Queries) GetCartItemDetail(ctx context.Context, id uuid.UUID) (GetCartItemDetailRow, error) {
	row := q.queryRow(ctx, q.getCartItemDetailStmt, getCartItemDetail, id)
	var i GetCartItemDetailRow
	err := row.Scan(
		&i.ID,
		&i.CartID,
		&i.Quantity,
		&i.Price,
		&i.QtyInStock,
		&i.Total,
		&i.DiscountPercent,
		&i.DiscountActive,
	)
	return i, err
}

const getCartProductDetailList = `-- name: GetCartProductDetailList :many
SELECT
	cart_item.id,
	cart_item.cart_id,
	cart_item.quantity,
	product.price AS price,
	product_inventory.quantity AS qty_in_stock,
	float8(CASE
		WHEN product_discount.active THEN (product.price * (1 - product_discount.discount_percent / 100) * cart_item.quantity)
		ELSE (product.price * cart_item.quantity )
	END) AS total,
	product_discount.discount_percent AS discount_percent,
	product_discount.active AS discount_active
FROM cart_item
LEFT JOIN product
    ON cart_item.product_id = product.id
LEFT JOIN product_discount
    ON product.discount_id = product_discount.id
LEFT JOIN product_inventory
    ON product.inventory_id = product_inventory.id
WHERE cart_item.cart_id = $1
LIMIT $2
OFFSET $3
`

type GetCartProductDetailListParams struct {
	CartID uuid.UUID `json:"cart_id"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

type GetCartProductDetailListRow struct {
	ID              uuid.UUID       `json:"id"`
	CartID          uuid.UUID       `json:"cart_id"`
	Quantity        int32           `json:"quantity"`
	Price           sql.NullFloat64 `json:"price"`
	QtyInStock      sql.NullInt32   `json:"qty_in_stock"`
	Total           float64         `json:"total"`
	DiscountPercent sql.NullFloat64 `json:"discount_percent"`
	DiscountActive  sql.NullBool    `json:"discount_active"`
}

func (q *Queries) GetCartProductDetailList(ctx context.Context, arg GetCartProductDetailListParams) ([]GetCartProductDetailListRow, error) {
	rows, err := q.query(ctx, q.getCartProductDetailListStmt, getCartProductDetailList, arg.CartID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetCartProductDetailListRow{}
	for rows.Next() {
		var i GetCartProductDetailListRow
		if err := rows.Scan(
			&i.ID,
			&i.CartID,
			&i.Quantity,
			&i.Price,
			&i.QtyInStock,
			&i.Total,
			&i.DiscountPercent,
			&i.DiscountActive,
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

const getTotal = `-- name: GetTotal :one
SELECT
	float8(sum(a.price)) AS total
FROM (
	SELECT
		CASE
			WHEN product_discount.active THEN (product.price * (1 - product_discount.discount_percent / 100) * cart_item.quantity)
			ELSE (product.price * cart_item.quantity )
		END AS price
	FROM cart_item
	LEFT JOIN product
	    ON cart_item.product_id = product.id
	LEFT JOIN product_discount
	    ON product.discount_id = product_discount.id
	WHERE cart_item.cart_id = $1
) AS a
`

func (q *Queries) GetTotal(ctx context.Context, cartID uuid.UUID) (float64, error) {
	row := q.queryRow(ctx, q.getTotalStmt, getTotal, cartID)
	var total float64
	err := row.Scan(&total)
	return total, err
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
