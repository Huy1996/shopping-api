// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.addDiscountStmt, err = db.PrepareContext(ctx, addDiscount); err != nil {
		return nil, fmt.Errorf("error preparing query AddDiscount: %w", err)
	}
	if q.addToCartStmt, err = db.PrepareContext(ctx, addToCart); err != nil {
		return nil, fmt.Errorf("error preparing query AddToCart: %w", err)
	}
	if q.createCartStmt, err = db.PrepareContext(ctx, createCart); err != nil {
		return nil, fmt.Errorf("error preparing query CreateCart: %w", err)
	}
	if q.createCategoryStmt, err = db.PrepareContext(ctx, createCategory); err != nil {
		return nil, fmt.Errorf("error preparing query CreateCategory: %w", err)
	}
	if q.createOrderItemStmt, err = db.PrepareContext(ctx, createOrderItem); err != nil {
		return nil, fmt.Errorf("error preparing query CreateOrderItem: %w", err)
	}
	if q.createOrderRecordStmt, err = db.PrepareContext(ctx, createOrderRecord); err != nil {
		return nil, fmt.Errorf("error preparing query CreateOrderRecord: %w", err)
	}
	if q.createPaymentRecordStmt, err = db.PrepareContext(ctx, createPaymentRecord); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePaymentRecord: %w", err)
	}
	if q.createProductStmt, err = db.PrepareContext(ctx, createProduct); err != nil {
		return nil, fmt.Errorf("error preparing query CreateProduct: %w", err)
	}
	if q.createProductDiscountStmt, err = db.PrepareContext(ctx, createProductDiscount); err != nil {
		return nil, fmt.Errorf("error preparing query CreateProductDiscount: %w", err)
	}
	if q.createProductInventoryStmt, err = db.PrepareContext(ctx, createProductInventory); err != nil {
		return nil, fmt.Errorf("error preparing query CreateProductInventory: %w", err)
	}
	if q.createUserAddressStmt, err = db.PrepareContext(ctx, createUserAddress); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUserAddress: %w", err)
	}
	if q.createUserCredentialStmt, err = db.PrepareContext(ctx, createUserCredential); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUserCredential: %w", err)
	}
	if q.createUserInfoStmt, err = db.PrepareContext(ctx, createUserInfo); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUserInfo: %w", err)
	}
	if q.deleteCartStmt, err = db.PrepareContext(ctx, deleteCart); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteCart: %w", err)
	}
	if q.getAddressStmt, err = db.PrepareContext(ctx, getAddress); err != nil {
		return nil, fmt.Errorf("error preparing query GetAddress: %w", err)
	}
	if q.getCartByIDStmt, err = db.PrepareContext(ctx, getCartByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetCartByID: %w", err)
	}
	if q.getCartByOwnerStmt, err = db.PrepareContext(ctx, getCartByOwner); err != nil {
		return nil, fmt.Errorf("error preparing query GetCartByOwner: %w", err)
	}
	if q.getCartItemDetailStmt, err = db.PrepareContext(ctx, getCartItemDetail); err != nil {
		return nil, fmt.Errorf("error preparing query GetCartItemDetail: %w", err)
	}
	if q.getCartProductDetailListStmt, err = db.PrepareContext(ctx, getCartProductDetailList); err != nil {
		return nil, fmt.Errorf("error preparing query GetCartProductDetailList: %w", err)
	}
	if q.getCartProductListStmt, err = db.PrepareContext(ctx, getCartProductList); err != nil {
		return nil, fmt.Errorf("error preparing query GetCartProductList: %w", err)
	}
	if q.getCategoryDetailStmt, err = db.PrepareContext(ctx, getCategoryDetail); err != nil {
		return nil, fmt.Errorf("error preparing query GetCategoryDetail: %w", err)
	}
	if q.getCategoryForUpdateStmt, err = db.PrepareContext(ctx, getCategoryForUpdate); err != nil {
		return nil, fmt.Errorf("error preparing query GetCategoryForUpdate: %w", err)
	}
	if q.getDiscountDetailStmt, err = db.PrepareContext(ctx, getDiscountDetail); err != nil {
		return nil, fmt.Errorf("error preparing query GetDiscountDetail: %w", err)
	}
	if q.getInventoryDetailStmt, err = db.PrepareContext(ctx, getInventoryDetail); err != nil {
		return nil, fmt.Errorf("error preparing query GetInventoryDetail: %w", err)
	}
	if q.getListAddressesStmt, err = db.PrepareContext(ctx, getListAddresses); err != nil {
		return nil, fmt.Errorf("error preparing query GetListAddresses: %w", err)
	}
	if q.getListCategoriesStmt, err = db.PrepareContext(ctx, getListCategories); err != nil {
		return nil, fmt.Errorf("error preparing query GetListCategories: %w", err)
	}
	if q.getNumberAddressesStmt, err = db.PrepareContext(ctx, getNumberAddresses); err != nil {
		return nil, fmt.Errorf("error preparing query GetNumberAddresses: %w", err)
	}
	if q.getOrderItemListStmt, err = db.PrepareContext(ctx, getOrderItemList); err != nil {
		return nil, fmt.Errorf("error preparing query GetOrderItemList: %w", err)
	}
	if q.getPaymentRecordStmt, err = db.PrepareContext(ctx, getPaymentRecord); err != nil {
		return nil, fmt.Errorf("error preparing query GetPaymentRecord: %w", err)
	}
	if q.getProductDetailStmt, err = db.PrepareContext(ctx, getProductDetail); err != nil {
		return nil, fmt.Errorf("error preparing query GetProductDetail: %w", err)
	}
	if q.getProductListStmt, err = db.PrepareContext(ctx, getProductList); err != nil {
		return nil, fmt.Errorf("error preparing query GetProductList: %w", err)
	}
	if q.getTotalStmt, err = db.PrepareContext(ctx, getTotal); err != nil {
		return nil, fmt.Errorf("error preparing query GetTotal: %w", err)
	}
	if q.getUserCredentialStmt, err = db.PrepareContext(ctx, getUserCredential); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserCredential: %w", err)
	}
	if q.getUserInfoByIDStmt, err = db.PrepareContext(ctx, getUserInfoByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserInfoByID: %w", err)
	}
	if q.getUserInfoByUserIDStmt, err = db.PrepareContext(ctx, getUserInfoByUserID); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserInfoByUserID: %w", err)
	}
	if q.getUserOrderSummaryStmt, err = db.PrepareContext(ctx, getUserOrderSummary); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserOrderSummary: %w", err)
	}
	if q.removeDiscountStmt, err = db.PrepareContext(ctx, removeDiscount); err != nil {
		return nil, fmt.Errorf("error preparing query RemoveDiscount: %w", err)
	}
	if q.removeItemStmt, err = db.PrepareContext(ctx, removeItem); err != nil {
		return nil, fmt.Errorf("error preparing query RemoveItem: %w", err)
	}
	if q.updateCartItemQtyStmt, err = db.PrepareContext(ctx, updateCartItemQty); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateCartItemQty: %w", err)
	}
	if q.updateCategoryStmt, err = db.PrepareContext(ctx, updateCategory); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateCategory: %w", err)
	}
	if q.updateDiscountStmt, err = db.PrepareContext(ctx, updateDiscount); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateDiscount: %w", err)
	}
	if q.updatePaymentStatusStmt, err = db.PrepareContext(ctx, updatePaymentStatus); err != nil {
		return nil, fmt.Errorf("error preparing query UpdatePaymentStatus: %w", err)
	}
	if q.updateProductInventoryStmt, err = db.PrepareContext(ctx, updateProductInventory); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateProductInventory: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.addDiscountStmt != nil {
		if cerr := q.addDiscountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addDiscountStmt: %w", cerr)
		}
	}
	if q.addToCartStmt != nil {
		if cerr := q.addToCartStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addToCartStmt: %w", cerr)
		}
	}
	if q.createCartStmt != nil {
		if cerr := q.createCartStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createCartStmt: %w", cerr)
		}
	}
	if q.createCategoryStmt != nil {
		if cerr := q.createCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createCategoryStmt: %w", cerr)
		}
	}
	if q.createOrderItemStmt != nil {
		if cerr := q.createOrderItemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createOrderItemStmt: %w", cerr)
		}
	}
	if q.createOrderRecordStmt != nil {
		if cerr := q.createOrderRecordStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createOrderRecordStmt: %w", cerr)
		}
	}
	if q.createPaymentRecordStmt != nil {
		if cerr := q.createPaymentRecordStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPaymentRecordStmt: %w", cerr)
		}
	}
	if q.createProductStmt != nil {
		if cerr := q.createProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createProductStmt: %w", cerr)
		}
	}
	if q.createProductDiscountStmt != nil {
		if cerr := q.createProductDiscountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createProductDiscountStmt: %w", cerr)
		}
	}
	if q.createProductInventoryStmt != nil {
		if cerr := q.createProductInventoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createProductInventoryStmt: %w", cerr)
		}
	}
	if q.createUserAddressStmt != nil {
		if cerr := q.createUserAddressStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserAddressStmt: %w", cerr)
		}
	}
	if q.createUserCredentialStmt != nil {
		if cerr := q.createUserCredentialStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserCredentialStmt: %w", cerr)
		}
	}
	if q.createUserInfoStmt != nil {
		if cerr := q.createUserInfoStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserInfoStmt: %w", cerr)
		}
	}
	if q.deleteCartStmt != nil {
		if cerr := q.deleteCartStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteCartStmt: %w", cerr)
		}
	}
	if q.getAddressStmt != nil {
		if cerr := q.getAddressStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAddressStmt: %w", cerr)
		}
	}
	if q.getCartByIDStmt != nil {
		if cerr := q.getCartByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCartByIDStmt: %w", cerr)
		}
	}
	if q.getCartByOwnerStmt != nil {
		if cerr := q.getCartByOwnerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCartByOwnerStmt: %w", cerr)
		}
	}
	if q.getCartItemDetailStmt != nil {
		if cerr := q.getCartItemDetailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCartItemDetailStmt: %w", cerr)
		}
	}
	if q.getCartProductDetailListStmt != nil {
		if cerr := q.getCartProductDetailListStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCartProductDetailListStmt: %w", cerr)
		}
	}
	if q.getCartProductListStmt != nil {
		if cerr := q.getCartProductListStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCartProductListStmt: %w", cerr)
		}
	}
	if q.getCategoryDetailStmt != nil {
		if cerr := q.getCategoryDetailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCategoryDetailStmt: %w", cerr)
		}
	}
	if q.getCategoryForUpdateStmt != nil {
		if cerr := q.getCategoryForUpdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCategoryForUpdateStmt: %w", cerr)
		}
	}
	if q.getDiscountDetailStmt != nil {
		if cerr := q.getDiscountDetailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getDiscountDetailStmt: %w", cerr)
		}
	}
	if q.getInventoryDetailStmt != nil {
		if cerr := q.getInventoryDetailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getInventoryDetailStmt: %w", cerr)
		}
	}
	if q.getListAddressesStmt != nil {
		if cerr := q.getListAddressesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getListAddressesStmt: %w", cerr)
		}
	}
	if q.getListCategoriesStmt != nil {
		if cerr := q.getListCategoriesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getListCategoriesStmt: %w", cerr)
		}
	}
	if q.getNumberAddressesStmt != nil {
		if cerr := q.getNumberAddressesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getNumberAddressesStmt: %w", cerr)
		}
	}
	if q.getOrderItemListStmt != nil {
		if cerr := q.getOrderItemListStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getOrderItemListStmt: %w", cerr)
		}
	}
	if q.getPaymentRecordStmt != nil {
		if cerr := q.getPaymentRecordStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPaymentRecordStmt: %w", cerr)
		}
	}
	if q.getProductDetailStmt != nil {
		if cerr := q.getProductDetailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProductDetailStmt: %w", cerr)
		}
	}
	if q.getProductListStmt != nil {
		if cerr := q.getProductListStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProductListStmt: %w", cerr)
		}
	}
	if q.getTotalStmt != nil {
		if cerr := q.getTotalStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTotalStmt: %w", cerr)
		}
	}
	if q.getUserCredentialStmt != nil {
		if cerr := q.getUserCredentialStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserCredentialStmt: %w", cerr)
		}
	}
	if q.getUserInfoByIDStmt != nil {
		if cerr := q.getUserInfoByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserInfoByIDStmt: %w", cerr)
		}
	}
	if q.getUserInfoByUserIDStmt != nil {
		if cerr := q.getUserInfoByUserIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserInfoByUserIDStmt: %w", cerr)
		}
	}
	if q.getUserOrderSummaryStmt != nil {
		if cerr := q.getUserOrderSummaryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserOrderSummaryStmt: %w", cerr)
		}
	}
	if q.removeDiscountStmt != nil {
		if cerr := q.removeDiscountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removeDiscountStmt: %w", cerr)
		}
	}
	if q.removeItemStmt != nil {
		if cerr := q.removeItemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removeItemStmt: %w", cerr)
		}
	}
	if q.updateCartItemQtyStmt != nil {
		if cerr := q.updateCartItemQtyStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateCartItemQtyStmt: %w", cerr)
		}
	}
	if q.updateCategoryStmt != nil {
		if cerr := q.updateCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateCategoryStmt: %w", cerr)
		}
	}
	if q.updateDiscountStmt != nil {
		if cerr := q.updateDiscountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateDiscountStmt: %w", cerr)
		}
	}
	if q.updatePaymentStatusStmt != nil {
		if cerr := q.updatePaymentStatusStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updatePaymentStatusStmt: %w", cerr)
		}
	}
	if q.updateProductInventoryStmt != nil {
		if cerr := q.updateProductInventoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateProductInventoryStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                           DBTX
	tx                           *sql.Tx
	addDiscountStmt              *sql.Stmt
	addToCartStmt                *sql.Stmt
	createCartStmt               *sql.Stmt
	createCategoryStmt           *sql.Stmt
	createOrderItemStmt          *sql.Stmt
	createOrderRecordStmt        *sql.Stmt
	createPaymentRecordStmt      *sql.Stmt
	createProductStmt            *sql.Stmt
	createProductDiscountStmt    *sql.Stmt
	createProductInventoryStmt   *sql.Stmt
	createUserAddressStmt        *sql.Stmt
	createUserCredentialStmt     *sql.Stmt
	createUserInfoStmt           *sql.Stmt
	deleteCartStmt               *sql.Stmt
	getAddressStmt               *sql.Stmt
	getCartByIDStmt              *sql.Stmt
	getCartByOwnerStmt           *sql.Stmt
	getCartItemDetailStmt        *sql.Stmt
	getCartProductDetailListStmt *sql.Stmt
	getCartProductListStmt       *sql.Stmt
	getCategoryDetailStmt        *sql.Stmt
	getCategoryForUpdateStmt     *sql.Stmt
	getDiscountDetailStmt        *sql.Stmt
	getInventoryDetailStmt       *sql.Stmt
	getListAddressesStmt         *sql.Stmt
	getListCategoriesStmt        *sql.Stmt
	getNumberAddressesStmt       *sql.Stmt
	getOrderItemListStmt         *sql.Stmt
	getPaymentRecordStmt         *sql.Stmt
	getProductDetailStmt         *sql.Stmt
	getProductListStmt           *sql.Stmt
	getTotalStmt                 *sql.Stmt
	getUserCredentialStmt        *sql.Stmt
	getUserInfoByIDStmt          *sql.Stmt
	getUserInfoByUserIDStmt      *sql.Stmt
	getUserOrderSummaryStmt      *sql.Stmt
	removeDiscountStmt           *sql.Stmt
	removeItemStmt               *sql.Stmt
	updateCartItemQtyStmt        *sql.Stmt
	updateCategoryStmt           *sql.Stmt
	updateDiscountStmt           *sql.Stmt
	updatePaymentStatusStmt      *sql.Stmt
	updateProductInventoryStmt   *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                           tx,
		tx:                           tx,
		addDiscountStmt:              q.addDiscountStmt,
		addToCartStmt:                q.addToCartStmt,
		createCartStmt:               q.createCartStmt,
		createCategoryStmt:           q.createCategoryStmt,
		createOrderItemStmt:          q.createOrderItemStmt,
		createOrderRecordStmt:        q.createOrderRecordStmt,
		createPaymentRecordStmt:      q.createPaymentRecordStmt,
		createProductStmt:            q.createProductStmt,
		createProductDiscountStmt:    q.createProductDiscountStmt,
		createProductInventoryStmt:   q.createProductInventoryStmt,
		createUserAddressStmt:        q.createUserAddressStmt,
		createUserCredentialStmt:     q.createUserCredentialStmt,
		createUserInfoStmt:           q.createUserInfoStmt,
		deleteCartStmt:               q.deleteCartStmt,
		getAddressStmt:               q.getAddressStmt,
		getCartByIDStmt:              q.getCartByIDStmt,
		getCartByOwnerStmt:           q.getCartByOwnerStmt,
		getCartItemDetailStmt:        q.getCartItemDetailStmt,
		getCartProductDetailListStmt: q.getCartProductDetailListStmt,
		getCartProductListStmt:       q.getCartProductListStmt,
		getCategoryDetailStmt:        q.getCategoryDetailStmt,
		getCategoryForUpdateStmt:     q.getCategoryForUpdateStmt,
		getDiscountDetailStmt:        q.getDiscountDetailStmt,
		getInventoryDetailStmt:       q.getInventoryDetailStmt,
		getListAddressesStmt:         q.getListAddressesStmt,
		getListCategoriesStmt:        q.getListCategoriesStmt,
		getNumberAddressesStmt:       q.getNumberAddressesStmt,
		getOrderItemListStmt:         q.getOrderItemListStmt,
		getPaymentRecordStmt:         q.getPaymentRecordStmt,
		getProductDetailStmt:         q.getProductDetailStmt,
		getProductListStmt:           q.getProductListStmt,
		getTotalStmt:                 q.getTotalStmt,
		getUserCredentialStmt:        q.getUserCredentialStmt,
		getUserInfoByIDStmt:          q.getUserInfoByIDStmt,
		getUserInfoByUserIDStmt:      q.getUserInfoByUserIDStmt,
		getUserOrderSummaryStmt:      q.getUserOrderSummaryStmt,
		removeDiscountStmt:           q.removeDiscountStmt,
		removeItemStmt:               q.removeItemStmt,
		updateCartItemQtyStmt:        q.updateCartItemQtyStmt,
		updateCategoryStmt:           q.updateCategoryStmt,
		updateDiscountStmt:           q.updateDiscountStmt,
		updatePaymentStatusStmt:      q.updatePaymentStatusStmt,
		updateProductInventoryStmt:   q.updateProductInventoryStmt,
	}
}
