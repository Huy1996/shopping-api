package util

const (
	// State
	CA = "California"
	TX = "Texas"
	NV = "Nevada"

	// City
	SJ = "San Jose"
	SF = "San Francisco"
	GR = "Gilroy"
	MV = "Mountain View"
)

const CurrencyTolerance = 0.00001

type PaymentType string

const (
	VISA     PaymentType = "VISA"
	MASTER   PaymentType = "MASTER CARD"
	AMERICAN PaymentType = "AMERICAN EXPRESS"
)

type PaymentStatus string

const (
	Pending  PaymentStatus = "Pending"
	Succeed  PaymentStatus = "Succeed"
	Rejected PaymentStatus = "Rejected"
)

type OrderStatus string

const (
	Preparing     OrderStatus = "Preparing"
	Shipped       OrderStatus = "Shipped"
	Delivered     OrderStatus = "Delivered"
	PendingCancel OrderStatus = "Pending Cancel"
	Canceled      OrderStatus = "Canceled"
)
