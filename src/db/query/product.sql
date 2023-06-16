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

-- name: GetCategoryDetail :one
SELECT * FROM product_category
WHERE id = $1
LIMIT 1;

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

-- name: GetInventoryDetail :one
SELECT * FROM product_inventory
WHERE id = $1
LIMIT 1;

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

-- name: GetProductDetail :one
SELECT
    product.id,
	product.name,
	product.price,
	product."SKU",
	product.description,
	product_category.name AS category,
	product_category.description AS category_description,
	product_discount.name AS discount_name,
	product_discount.discount_percent AS discount_percent,
	product_discount.active AS discount_active,
	product_discount.description AS discount_description,
	product_inventory.quantity AS quantity
FROM product
LEFT JOIN product_discount ON product.discount_id = product_discount.id
LEFT JOIN product_inventory ON product.inventory_id = product_inventory.id
LEFT JOIN product_category ON product.category_id = product_category.id
WHERE product.id = $1
LIMIT 1;

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

-- name: GetDiscountDetail :one
SELECT * FROM product_discount
WHERE id = $1
LIMIT 1;