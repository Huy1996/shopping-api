-- name: CreateCart :one
INSERT INTO user_cart (
    id,
    owner
) VALUES (
    $1, $2
) RETURNING *;

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

-- name: GetCartItemDetail :one
SELECT
	cart_item.id,
	cart_item.cart_id,
	cart_item.quantity,
	product.price AS price,
	product_discount.discount_percent AS discount_percent,
	product_discount.active AS discount_active
FROM cart_item
LEFT JOIN product ON cart_item.product_id = product.id
LEFT JOIN product_discount ON product.discount_id = product_discount.id
WHERE cart_item.id = $1;

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

-- name: GetTotal :one
SELECT
	float8(sum(a.price)) AS total
FROM (
	SELECT
		CASE
			WHEN product_discount.active THEN (product.price * (1 + product_discount.discount_percent / 100) * cart_item.quantity)
			ELSE (product.price * cart_item.quantity )
		END AS price
	FROM cart_item
	LEFT JOIN product ON cart_item.product_id = product.id
	LEFT JOIN product_discount ON product.discount_id = product_discount.id
	WHERE cart_item.cart_id = $1
) AS a;