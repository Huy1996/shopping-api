// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
	"shopping-cart/src/util"
)

type OrderStatus string

const (
	OrderStatusPreparing     OrderStatus = "Preparing"
	OrderStatusShipped       OrderStatus = "Shipped"
	OrderStatusDelivered     OrderStatus = "Delivered"
	OrderStatusPendingCancel OrderStatus = "Pending Cancel"
	OrderStatusCanceled      OrderStatus = "Canceled"
)

func (e *OrderStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = OrderStatus(s)
	case string:
		*e = OrderStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for OrderStatus: %T", src)
	}
	return nil
}

type NullOrderStatus struct {
	OrderStatus OrderStatus
	Valid       bool // Valid is true if OrderStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullOrderStatus) Scan(value interface{}) error {
	if value == nil {
		ns.OrderStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.OrderStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullOrderStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.OrderStatus), nil
}

func AllOrderStatusValues() []OrderStatus {
	return []OrderStatus{
		OrderStatusPreparing,
		OrderStatusShipped,
		OrderStatusDelivered,
		OrderStatusPendingCancel,
		OrderStatusCanceled,
	}
}

type PaymentStatus string

const (
	PaymentStatusPending  PaymentStatus = "Pending"
	PaymentStatusSucceed  PaymentStatus = "Succeed"
	PaymentStatusRejected PaymentStatus = "Rejected"
)

func (e *PaymentStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PaymentStatus(s)
	case string:
		*e = PaymentStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for PaymentStatus: %T", src)
	}
	return nil
}

type NullPaymentStatus struct {
	PaymentStatus PaymentStatus
	Valid         bool // Valid is true if PaymentStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPaymentStatus) Scan(value interface{}) error {
	if value == nil {
		ns.PaymentStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.PaymentStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPaymentStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.PaymentStatus), nil
}

func AllPaymentStatusValues() []PaymentStatus {
	return []PaymentStatus{
		PaymentStatusPending,
		PaymentStatusSucceed,
		PaymentStatusRejected,
	}
}

type PaymentType string

const (
	PaymentTypeVISA            PaymentType = "VISA"
	PaymentTypeMASTERCARD      PaymentType = "MASTER CARD"
	PaymentTypeAMERICANEXPRESS PaymentType = "AMERICAN EXPRESS"
)

func (e *PaymentType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PaymentType(s)
	case string:
		*e = PaymentType(s)
	default:
		return fmt.Errorf("unsupported scan type for PaymentType: %T", src)
	}
	return nil
}

type NullPaymentType struct {
	PaymentType PaymentType
	Valid       bool // Valid is true if PaymentType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPaymentType) Scan(value interface{}) error {
	if value == nil {
		ns.PaymentType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.PaymentType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPaymentType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.PaymentType), nil
}

func AllPaymentTypeValues() []PaymentType {
	return []PaymentType{
		PaymentTypeVISA,
		PaymentTypeMASTERCARD,
		PaymentTypeAMERICANEXPRESS,
	}
}

type CartItem struct {
	ID        uuid.UUID `json:"id"`
	CartID    uuid.UUID `json:"cart_id"`
	ProductID uuid.UUID `json:"product_id"`
	// Cannot be less than 1
	Quantity int32 `json:"quantity"`
}

type OrderDetail struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Total     float64   `json:"total"`
	PaymentID uuid.UUID `json:"payment_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderItem struct {
	ID        uuid.UUID        `json:"id"`
	OrderID   uuid.UUID        `json:"order_id"`
	ProductID uuid.UUID        `json:"product_id"`
	Quantity  int32            `json:"quantity"`
	Status    util.OrderStatus `json:"status"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

type PaymentDetail struct {
	ID         uuid.UUID          `json:"id"`
	Amount     float64            `json:"amount"`
	Type       util.PaymentType   `json:"type"`
	Status     util.PaymentStatus `json:"status"`
	CardNumber string             `json:"card_number"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SKU         string    `json:"SKU"`
	// Cannot be negative
	Price       float64       `json:"price"`
	CategoryID  uuid.UUID     `json:"category_id"`
	InventoryID uuid.UUID     `json:"inventory_id"`
	DiscountID  uuid.NullUUID `json:"discount_id"`
}

type ProductCategory struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductDiscount struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	DiscountPercent float64   `json:"discount_percent"`
	Active          bool      `json:"active"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type ProductInventory struct {
	ID        uuid.UUID `json:"id"`
	Quantity  int32     `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserAddress struct {
	ID          uuid.UUID `json:"id"`
	Owner       uuid.UUID `json:"owner"`
	AddressName string    `json:"address_name"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Zipcode     int32     `json:"zipcode"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserCart struct {
	ID    uuid.UUID `json:"id"`
	Owner uuid.UUID `json:"owner"`
}

type UserCredential struct {
	ID                uuid.UUID `json:"id"`
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type UserInfo struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	PhoneNumber string    `json:"phone_number"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	MiddleName  string    `json:"middle_name"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}
