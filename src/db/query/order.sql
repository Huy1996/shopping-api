-- name: CreatePaymentRecord :one
INSERT INTO payment_detail (
    id,
    amount,
    type,
    status,
    card_number
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetPaymentRecord :one
SELECT * FROM payment_detail
WHERE id = $1
LIMIT 1;

-- name: UpdatePaymentStatus :one
UPDATE payment_detail
SET
    status = $1,
    updated_at = now()
WHERE id = $2
RETURNING *;

-- name: CreateOrderRecord :one
INSERT INTO order_detail (
    id,
    user_id,
    total,
    payment_id
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: CreateOrderItem :one
INSERT INTO order_item (
    id,
    order_id,
    product_id,
    quantity,
    status
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUserOrderSummary :many
SELECT
	order_detail.id,
	amount as total,
	(SELECT COUNT(*) FROM order_item WHERE order_id = order_detail.id) as number_product
FROM order_detail
LEFT JOIN payment_detail
    ON order_detail.payment_id = payment_detail.id
WHERE order_detail.user_id = $1
LIMIT $2
OFFSET $3;

-- name: GetOrderItemList :many
SELECT * FROM order_item
WHERE order_id = $1
LIMIT $2
OFFSET $3;