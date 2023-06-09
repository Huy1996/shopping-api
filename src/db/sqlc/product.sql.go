// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: product.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const addDiscount = `-- name: AddDiscount :one
UPDATE product
SET discount_id = $1
WHERE id = $2
RETURNING id, name, description, "SKU", price, category_id, inventory_id, discount_id
`

type AddDiscountParams struct {
	DiscountID uuid.NullUUID `json:"discount_id"`
	ID         uuid.UUID     `json:"id"`
}

func (q *Queries) AddDiscount(ctx context.Context, arg AddDiscountParams) (Product, error) {
	row := q.queryRow(ctx, q.addDiscountStmt, addDiscount, arg.DiscountID, arg.ID)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.SKU,
		&i.Price,
		&i.CategoryID,
		&i.InventoryID,
		&i.DiscountID,
	)
	return i, err
}

const createCategory = `-- name: CreateCategory :one
INSERT INTO product_category (
    id,
    name,
    description
) VALUES (
    $1, $2, $3
) RETURNING id, name, description, created_at, updated_at
`

type CreateCategoryParams struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (ProductCategory, error) {
	row := q.queryRow(ctx, q.createCategoryStmt, createCategory, arg.ID, arg.Name, arg.Description)
	var i ProductCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createProduct = `-- name: CreateProduct :one
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
) RETURNING id, name, description, "SKU", price, category_id, inventory_id, discount_id
`

type CreateProductParams struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SKU         string    `json:"SKU"`
	Price       float64   `json:"price"`
	CategoryID  uuid.UUID `json:"category_id"`
	InventoryID uuid.UUID `json:"inventory_id"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.queryRow(ctx, q.createProductStmt, createProduct,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.SKU,
		arg.Price,
		arg.CategoryID,
		arg.InventoryID,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.SKU,
		&i.Price,
		&i.CategoryID,
		&i.InventoryID,
		&i.DiscountID,
	)
	return i, err
}

const createProductDiscount = `-- name: CreateProductDiscount :one
INSERT INTO product_discount (
    id,
    name,
    description,
    discount_percent
) VALUES (
    $1, $2, $3, $4
) RETURNING id, name, description, discount_percent, active, created_at, updated_at
`

type CreateProductDiscountParams struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	DiscountPercent float64   `json:"discount_percent"`
}

func (q *Queries) CreateProductDiscount(ctx context.Context, arg CreateProductDiscountParams) (ProductDiscount, error) {
	row := q.queryRow(ctx, q.createProductDiscountStmt, createProductDiscount,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.DiscountPercent,
	)
	var i ProductDiscount
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.DiscountPercent,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createProductInventory = `-- name: CreateProductInventory :one
INSERT INTO product_inventory (
    id,
    quantity
) VALUES (
    $1, $2
) RETURNING id, quantity, created_at, updated_at
`

type CreateProductInventoryParams struct {
	ID       uuid.UUID `json:"id"`
	Quantity int32     `json:"quantity"`
}

func (q *Queries) CreateProductInventory(ctx context.Context, arg CreateProductInventoryParams) (ProductInventory, error) {
	row := q.queryRow(ctx, q.createProductInventoryStmt, createProductInventory, arg.ID, arg.Quantity)
	var i ProductInventory
	err := row.Scan(
		&i.ID,
		&i.Quantity,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCategoryDetail = `-- name: GetCategoryDetail :one
SELECT id, name, description, created_at, updated_at FROM product_category
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetCategoryDetail(ctx context.Context, id uuid.UUID) (ProductCategory, error) {
	row := q.queryRow(ctx, q.getCategoryDetailStmt, getCategoryDetail, id)
	var i ProductCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCategoryForUpdate = `-- name: GetCategoryForUpdate :one
SELECT id, name, description, created_at, updated_at FROM product_category
WHERE id = $1
LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetCategoryForUpdate(ctx context.Context, id uuid.UUID) (ProductCategory, error) {
	row := q.queryRow(ctx, q.getCategoryForUpdateStmt, getCategoryForUpdate, id)
	var i ProductCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getDiscountDetail = `-- name: GetDiscountDetail :one
SELECT id, name, description, discount_percent, active, created_at, updated_at FROM product_discount
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetDiscountDetail(ctx context.Context, id uuid.UUID) (ProductDiscount, error) {
	row := q.queryRow(ctx, q.getDiscountDetailStmt, getDiscountDetail, id)
	var i ProductDiscount
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.DiscountPercent,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getInventoryDetail = `-- name: GetInventoryDetail :one
SELECT id, quantity, created_at, updated_at FROM product_inventory
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetInventoryDetail(ctx context.Context, id uuid.UUID) (ProductInventory, error) {
	row := q.queryRow(ctx, q.getInventoryDetailStmt, getInventoryDetail, id)
	var i ProductInventory
	err := row.Scan(
		&i.ID,
		&i.Quantity,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getListCategories = `-- name: GetListCategories :many
SELECT id, name, description, created_at, updated_at FROM product_category
`

func (q *Queries) GetListCategories(ctx context.Context) ([]ProductCategory, error) {
	rows, err := q.query(ctx, q.getListCategoriesStmt, getListCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProductCategory{}
	for rows.Next() {
		var i ProductCategory
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getProductDetail = `-- name: GetProductDetail :one
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
LIMIT 1
`

type GetProductDetailRow struct {
	ID                  uuid.UUID       `json:"id"`
	Name                string          `json:"name"`
	Price               float64         `json:"price"`
	SKU                 string          `json:"SKU"`
	Description         string          `json:"description"`
	Category            sql.NullString  `json:"category"`
	CategoryDescription sql.NullString  `json:"category_description"`
	DiscountName        sql.NullString  `json:"discount_name"`
	DiscountPercent     sql.NullFloat64 `json:"discount_percent"`
	DiscountActive      sql.NullBool    `json:"discount_active"`
	DiscountDescription sql.NullString  `json:"discount_description"`
	Quantity            sql.NullInt32   `json:"quantity"`
}

func (q *Queries) GetProductDetail(ctx context.Context, id uuid.UUID) (GetProductDetailRow, error) {
	row := q.queryRow(ctx, q.getProductDetailStmt, getProductDetail, id)
	var i GetProductDetailRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.SKU,
		&i.Description,
		&i.Category,
		&i.CategoryDescription,
		&i.DiscountName,
		&i.DiscountPercent,
		&i.DiscountActive,
		&i.DiscountDescription,
		&i.Quantity,
	)
	return i, err
}

const getProductList = `-- name: GetProductList :many
SELECT id, name, description, "SKU", price, category_id, inventory_id, discount_id FROM product
ORDER BY id
LIMIT $1
OFFSET $2
`

type GetProductListParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetProductList(ctx context.Context, arg GetProductListParams) ([]Product, error) {
	rows, err := q.query(ctx, q.getProductListStmt, getProductList, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.SKU,
			&i.Price,
			&i.CategoryID,
			&i.InventoryID,
			&i.DiscountID,
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

const removeDiscount = `-- name: RemoveDiscount :one
UPDATE product
SET discount_id = NULL
WHERE id = $1
RETURNING id, name, description, "SKU", price, category_id, inventory_id, discount_id
`

func (q *Queries) RemoveDiscount(ctx context.Context, id uuid.UUID) (Product, error) {
	row := q.queryRow(ctx, q.removeDiscountStmt, removeDiscount, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.SKU,
		&i.Price,
		&i.CategoryID,
		&i.InventoryID,
		&i.DiscountID,
	)
	return i, err
}

const updateCategory = `-- name: UpdateCategory :one
UPDATE product_category
SET
    name = COALESCE($1, name),
    description = COALESCE($2, description),
    updated_at = now()
WHERE
    id = $3
RETURNING id, name, description, created_at, updated_at
`

type UpdateCategoryParams struct {
	Name        sql.NullString `json:"name"`
	Description sql.NullString `json:"description"`
	ID          uuid.UUID      `json:"id"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (ProductCategory, error) {
	row := q.queryRow(ctx, q.updateCategoryStmt, updateCategory, arg.Name, arg.Description, arg.ID)
	var i ProductCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateDiscount = `-- name: UpdateDiscount :one
UPDATE product_discount
SET
    name = COALESCE($1, name),
    description = COALESCE($2, description),
    discount_percent = COALESCE($3, discount_percent),
    updated_at = now()
WHERE id = $4
RETURNING id, name, description, discount_percent, active, created_at, updated_at
`

type UpdateDiscountParams struct {
	Name            sql.NullString  `json:"name"`
	Description     sql.NullString  `json:"description"`
	DiscountPercent sql.NullFloat64 `json:"discount_percent"`
	ID              uuid.UUID       `json:"id"`
}

func (q *Queries) UpdateDiscount(ctx context.Context, arg UpdateDiscountParams) (ProductDiscount, error) {
	row := q.queryRow(ctx, q.updateDiscountStmt, updateDiscount,
		arg.Name,
		arg.Description,
		arg.DiscountPercent,
		arg.ID,
	)
	var i ProductDiscount
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.DiscountPercent,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateProductInventory = `-- name: UpdateProductInventory :one
UPDATE product_inventory
SET
    quantity = quantity + $1,
    updated_at = now()
WHERE id = $2
RETURNING id, quantity, created_at, updated_at
`

type UpdateProductInventoryParams struct {
	Amount int32     `json:"amount"`
	ID     uuid.UUID `json:"id"`
}

func (q *Queries) UpdateProductInventory(ctx context.Context, arg UpdateProductInventoryParams) (ProductInventory, error) {
	row := q.queryRow(ctx, q.updateProductInventoryStmt, updateProductInventory, arg.Amount, arg.ID)
	var i ProductInventory
	err := row.Scan(
		&i.ID,
		&i.Quantity,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
