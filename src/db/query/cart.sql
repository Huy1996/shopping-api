-- name: CreateCart :one
INSERT INTO user_cart (
    id,
    owner
) VALUES (
    $1, $2
) RETURNING *;

-- name: UpdateTotal :one
UPDATE user_cart
SET total = total + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: GetCartByID :one
SELECT * FROM user_cart
WHERE id = $1
LIMIT 1;

-- name: DeleteCart :exec
DELETE FROM user_cart
WHERE owner = $1;

-- name: AddToCart :one
INSERT INTO cart_item (
    id,
    cart_id,
    product_id,
    quantity
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetCartProductList :many
SELECT * FROM cart_item
WHERE cart_id = $1
LIMIT $2
OFFSET $3;

-- name: RemoveItem :exec
DELETE FROM cart_item
WHERE id = $1;

-- name: UpdateCartItemQty :one
UPDATE cart_item
SET quantity = $1
WHERE id = $2
RETURNING *;