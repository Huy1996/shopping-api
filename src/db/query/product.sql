-- name: CreateCategory :one
INSERT INTO product_category (
    id,
    name,
    description
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetCategoryForUpdate :one
SELECT * FROM product_category
WHERE id = $1
LIMIT 1
FOR NO KEY UPDATE;

-- name: UpdateCategory :one
UPDATE product_category
SET
    name = COALESCE(sqlc.narg(name), name),
    description = COALESCE(sqlc.narg(description), description),
    updated_at = now()
WHERE
    id = sqlc.arg(id)
RETURNING *;

-- name: GetListCategories :many
SELECT * FROM product_category;

-- name: CreateProductInventory :one
INSERT INTO product_inventory (
    id,
    quantity
) VALUES (
    $1, $2
) RETURNING *;

-- name: UpdateProductInventory :one
UPDATE product_inventory
SET
    quantity = quantity + sqlc.arg(amount),
    updated_at = now()
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: CreateProduct :one
INSERT INTO product (
    id,
    name,
    description,
    "SKU",
    price,
    category_id,
    inventory_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetProductList :many
SELECT * FROM product
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: AddDiscount :one
UPDATE product
SET discount_id = $1
WHERE id = $2
RETURNING *;

-- name: RemoveDiscount :one
UPDATE product
SET discount_id = NULL
WHERE id = $1
RETURNING *;

-- name: CreateProductDiscount :one
INSERT INTO product_discount (
    id,
    name,
    description,
    discount_percent
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: UpdateDiscount :one
UPDATE product_discount
SET
    name = COALESCE(sqlc.narg(name), name),
    description = COALESCE(sqlc.narg(description), description),
    discount_percent = COALESCE(sqlc.narg(discount_percent), discount_percent),
    updated_at = now()
WHERE id = sqlc.arg(id)
RETURNING *;