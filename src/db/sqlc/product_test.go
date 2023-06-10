package db

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"shopping-cart/src/util"
	"testing"
	"time"
)

func CreateRandomInventory(t *testing.T) ProductInventory {
	id, err := uuid.NewRandom()
	require.NoError(t, err)
	require.NotEmpty(t, id)

	arg := CreateProductInventoryParams{
		ID:       id,
		Quantity: int32(util.RandomInt(1, 20)),
	}

	productInventory, err := testQueries.CreateProductInventory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, productInventory)

	require.Equal(t, arg.ID, productInventory.ID)
	require.Equal(t, arg.Quantity, productInventory.Quantity)

	require.True(t, productInventory.UpdatedAt.IsZero())
	require.NotZero(t, productInventory.CreatedAt)
	return productInventory
}

func CreateRandomCategory(t *testing.T) ProductCategory {
	id, err := uuid.NewRandom()
	require.NoError(t, err)
	require.NotEmpty(t, id)

	arg := CreateCategoryParams{
		ID:          id,
		Name:        util.RandomName(),
		Description: util.RandomString(100),
	}

	category, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.ID, category.ID)
	require.Equal(t, arg.Name, category.Name)
	require.Equal(t, arg.Description, category.Description)

	require.True(t, category.UpdatedAt.IsZero())
	require.NotZero(t, category.CreatedAt)

	return category
}

func CreateRandomDiscount(t *testing.T) ProductDiscount {
	id, err := uuid.NewRandom()
	require.NoError(t, err)
	require.NotEmpty(t, id)

	arg := CreateProductDiscountParams{
		ID:              id,
		Name:            util.RandomName(),
		Description:     util.RandomString(100),
		DiscountPercent: util.RandomFloat(0, 99.99),
	}

	discount, err := testQueries.CreateProductDiscount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, discount)

	require.Equal(t, arg.ID, discount.ID)
	require.Equal(t, arg.Name, discount.Name)
	require.Equal(t, arg.Description, discount.Description)
	require.Equal(t, arg.DiscountPercent, discount.DiscountPercent)

	require.True(t, discount.UpdatedAt.IsZero())
	require.NotZero(t, discount.CreatedAt)

	return discount
}

func CreateRandomProduct(
	t *testing.T,
	inventory ProductInventory,
	category ProductCategory,
) Product {
	id, err := uuid.NewRandom()
	require.NoError(t, err)
	require.NotEmpty(t, id)

	arg := CreateProductParams{
		ID:          id,
		Name:        util.RandomName(),
		Description: util.RandomString(200),
		SKU:         util.RandomString(12),
		Price:       util.RandomFloat(0.01, 200),
		CategoryID:  category.ID,
		InventoryID: inventory.ID,
	}

	product, err := testQueries.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.ID, product.ID)
	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.Description, product.Description)
	require.Equal(t, arg.SKU, product.SKU)
	require.Equal(t, arg.Price, product.Price)
	require.Equal(t, arg.CategoryID, product.CategoryID)
	require.Equal(t, arg.InventoryID, product.InventoryID)
	require.Empty(t, product.DiscountID)

	return product
}

func TestUpdateCategoryNameOnly(t *testing.T) {
	oldCategory := CreateRandomCategory(t)

	newName := util.RandomName()
	updatedCategory, err := testQueries.UpdateCategory(context.Background(), UpdateCategoryParams{
		ID: oldCategory.ID,
		Name: sql.NullString{
			String: newName,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEmpty(t, updatedCategory)

	require.Equal(t, oldCategory.ID, updatedCategory.ID)
	require.Equal(t, oldCategory.Description, updatedCategory.Description)
	require.Equal(t, newName, updatedCategory.Name)

	require.NotEqual(t, oldCategory.Name, updatedCategory.Name)

	require.WithinDuration(t, oldCategory.CreatedAt, updatedCategory.CreatedAt, time.Second)
	require.NotZero(t, updatedCategory.UpdatedAt)
}

func TestUpdateCategoryDescriptionOnly(t *testing.T) {
	oldCategory := CreateRandomCategory(t)

	newDescription := util.RandomString(200)
	updatedCategory, err := testQueries.UpdateCategory(context.Background(), UpdateCategoryParams{
		ID: oldCategory.ID,
		Description: sql.NullString{
			String: newDescription,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEmpty(t, updatedCategory)

	require.Equal(t, oldCategory.ID, updatedCategory.ID)
	require.Equal(t, oldCategory.Name, updatedCategory.Name)
	require.Equal(t, newDescription, updatedCategory.Description)

	require.NotEqual(t, oldCategory.Description, updatedCategory.Description)

	require.WithinDuration(t, oldCategory.CreatedAt, updatedCategory.CreatedAt, time.Second)
	require.NotZero(t, updatedCategory.UpdatedAt)
}

func TestUpdateCategoryAllField(t *testing.T) {
	oldCategory := CreateRandomCategory(t)

	newName := util.RandomName()
	newDescription := util.RandomString(200)
	updatedCategory, err := testQueries.UpdateCategory(context.Background(), UpdateCategoryParams{
		ID: oldCategory.ID,
		Name: sql.NullString{
			String: newName,
			Valid:  true,
		},
		Description: sql.NullString{
			String: newDescription,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEmpty(t, updatedCategory)

	require.Equal(t, oldCategory.ID, updatedCategory.ID)
	require.Equal(t, newName, updatedCategory.Name)
	require.Equal(t, newDescription, updatedCategory.Description)

	require.NotEqual(t, oldCategory.Name, updatedCategory.Name)
	require.NotEqual(t, oldCategory.Description, updatedCategory.Description)

	require.WithinDuration(t, oldCategory.CreatedAt, updatedCategory.CreatedAt, time.Second)
	require.NotZero(t, updatedCategory.UpdatedAt)
}

func TestUpdateProductInventory(t *testing.T) {
	oldInventory := CreateRandomInventory(t)

	amount := int32(util.RandomInt(-100, 100))
	newInventory, err := testQueries.UpdateProductInventory(context.Background(), UpdateProductInventoryParams{
		ID:     oldInventory.ID,
		Amount: amount,
	})
	require.NoError(t, err)
	require.NotEmpty(t, newInventory)

	require.Equal(t, oldInventory.ID, newInventory.ID)
	require.NotEqual(t, oldInventory.Quantity, newInventory.Quantity)
	require.WithinDuration(t, oldInventory.CreatedAt, newInventory.CreatedAt, time.Second)

	require.NotZero(t, newInventory.UpdatedAt)
	require.Equal(t, amount, newInventory.Quantity-oldInventory.Quantity)
}

func TestUpdateProductDiscountNameOnly(t *testing.T) {
	oldDiscount := CreateRandomDiscount(t)

	newName := util.RandomName()
	newDiscount, err := testQueries.UpdateDiscount(context.Background(), UpdateDiscountParams{
		ID: oldDiscount.ID,
		Name: sql.NullString{
			String: newName,
			Valid:  true,
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, newDiscount)

	require.Equal(t, oldDiscount.ID, newDiscount.ID)
	require.Equal(t, oldDiscount.DiscountPercent, newDiscount.DiscountPercent)
	require.Equal(t, oldDiscount.Description, newDiscount.Description)
	require.Equal(t, oldDiscount.Active, newDiscount.Active)
	require.Equal(t, newName, newDiscount.Name)
	require.WithinDuration(t, oldDiscount.CreatedAt, newDiscount.CreatedAt, time.Second)

	require.NotZero(t, newDiscount.UpdatedAt)
	require.NotEqual(t, oldDiscount.Name, newDiscount.Name)
}

func TestUpdateProductDiscountDescriptionOnly(t *testing.T) {
	oldDiscount := CreateRandomDiscount(t)

	newDescription := util.RandomString(100)
	newDiscount, err := testQueries.UpdateDiscount(context.Background(), UpdateDiscountParams{
		ID: oldDiscount.ID,
		Description: sql.NullString{
			String: newDescription,
			Valid:  true,
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, newDiscount)

	require.Equal(t, oldDiscount.ID, newDiscount.ID)
	require.Equal(t, oldDiscount.DiscountPercent, newDiscount.DiscountPercent)
	require.Equal(t, oldDiscount.Name, newDiscount.Name)
	require.Equal(t, oldDiscount.Active, newDiscount.Active)
	require.Equal(t, newDescription, newDiscount.Description)
	require.WithinDuration(t, oldDiscount.CreatedAt, newDiscount.CreatedAt, time.Second)

	require.NotZero(t, newDiscount.UpdatedAt)
	require.NotEqual(t, oldDiscount.Description, newDiscount.Description)
}

func TestUpdateProductDiscountDiscountPercentOnly(t *testing.T) {
	oldDiscount := CreateRandomDiscount(t)

	newDiscountPercent := util.RandomFloat(0.01, 100)
	newDiscount, err := testQueries.UpdateDiscount(context.Background(), UpdateDiscountParams{
		ID: oldDiscount.ID,
		DiscountPercent: sql.NullString{
			String: newDiscountPercent,
			Valid:  true,
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, newDiscount)

	require.Equal(t, oldDiscount.ID, newDiscount.ID)
	require.Equal(t, oldDiscount.Description, newDiscount.Description)
	require.Equal(t, oldDiscount.Name, newDiscount.Name)
	require.Equal(t, oldDiscount.Active, newDiscount.Active)
	require.Equal(t, newDiscountPercent, newDiscount.DiscountPercent)
	require.WithinDuration(t, oldDiscount.CreatedAt, newDiscount.CreatedAt, time.Second)

	require.NotZero(t, newDiscount.UpdatedAt)
	require.NotEqual(t, oldDiscount.DiscountPercent, newDiscount.DiscountPercent)
}

func TestUpdateProductDiscountAllField(t *testing.T) {
	oldDiscount := CreateRandomDiscount(t)

	newName := util.RandomName()
	newDescription := util.RandomString(100)
	newDiscountPercent := util.RandomFloat(0.01, 100)
	newDiscount, err := testQueries.UpdateDiscount(context.Background(), UpdateDiscountParams{
		ID: oldDiscount.ID,
		Name: sql.NullString{
			String: newName,
			Valid:  true,
		},
		Description: sql.NullString{
			String: newDescription,
			Valid:  true,
		},
		DiscountPercent: sql.NullString{
			String: newDiscountPercent,
			Valid:  true,
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, newDiscount)

	require.Equal(t, oldDiscount.ID, newDiscount.ID)
	require.Equal(t, newDescription, newDiscount.Description)
	require.Equal(t, newName, newDiscount.Name)
	require.Equal(t, oldDiscount.Active, newDiscount.Active)
	require.Equal(t, newDiscountPercent, newDiscount.DiscountPercent)
	require.WithinDuration(t, oldDiscount.CreatedAt, newDiscount.CreatedAt, time.Second)

	require.NotZero(t, newDiscount.UpdatedAt)
	require.NotEqual(t, oldDiscount.DiscountPercent, newDiscount.DiscountPercent)
	require.NotEqual(t, oldDiscount.Description, newDiscount.Description)
	require.NotEqual(t, oldDiscount.Name, newDiscount.Name)
}

func TestAddRemoveDiscount(t *testing.T) {
	inventory := CreateRandomInventory(t)
	category := CreateRandomCategory(t)
	discount := CreateRandomDiscount(t)

	product := CreateRandomProduct(t, inventory, category)

	newProduct, err := testQueries.AddDiscount(context.Background(), AddDiscountParams{
		ID: product.ID,
		DiscountID: uuid.NullUUID{
			UUID:  discount.ID,
			Valid: true,
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, newProduct)

	require.Equal(t, product.ID, newProduct.ID)
	require.Equal(t, discount.ID, newProduct.DiscountID.UUID)

	newProduct2, err := testQueries.RemoveDiscount(context.Background(), newProduct.ID)
	require.NoError(t, err)
	require.NotEmpty(t, newProduct2)

	require.Equal(t, newProduct.ID, newProduct2.ID)
	require.Empty(t, newProduct2.DiscountID)
}
