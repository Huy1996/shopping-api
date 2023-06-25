package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"shopping-cart/src/util"
	"testing"
	"time"
)

func CreatePayment(t *testing.T) PaymentDetail {
	id, err := uuid.NewRandom()
	require.NoError(t, err)
	require.NotEmpty(t, id)

	arg := CreatePaymentRecordParams{
		ID:         id,
		Amount:     util.RandomFloat(10, 200),
		Type:       util.VISA,
		Status:     util.Pending,
		CardNumber: util.RandomString(12),
	}
	payment, err := testQueries.CreatePaymentRecord(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, payment)

	require.Equal(t, id, payment.ID)
	require.Equal(t, arg.Amount, payment.Amount)
	require.Equal(t, arg.Type, payment.Type)
	require.Equal(t, arg.Status, payment.Status)
	require.Equal(t, arg.CardNumber, payment.CardNumber)
	require.NotZero(t, payment.CreatedAt)
	require.True(t, payment.UpdatedAt.IsZero())

	return payment
}

func CreateOrder(t *testing.T, payment PaymentDetail, userInfo UserInfo) OrderDetail {
	id, err := uuid.NewRandom()
	require.NoError(t, err)
	require.NotEmpty(t, id)

	arg := CreateOrderRecordParams{
		ID:        id,
		UserID:    userInfo.ID,
		Total:     payment.Amount,
		PaymentID: payment.ID,
	}

	order, err := testQueries.CreateOrderRecord(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.Equal(t, arg.ID, order.ID)
	require.Equal(t, arg.UserID, order.UserID)
	require.Equal(t, arg.Total, order.Total)
	require.Equal(t, arg.PaymentID, order.PaymentID)

	require.NotZero(t, order.CreatedAt)
	require.True(t, order.UpdatedAt.IsZero())

	return order
}

func CreateOrderItem(t *testing.T, order OrderDetail) OrderItem {
	id, err := uuid.NewRandom()
	require.NoError(t, err)
	require.NotEmpty(t, id)

	product := CreateProduct(t)

	arg := CreateOrderItemParams{
		ID:        id,
		OrderID:   order.ID,
		ProductID: product.ID,
		Quantity:  int32(util.RandomInt(2, 10)),
		Status:    util.Preparing,
	}

	orderItem, err := testQueries.CreateOrderItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, orderItem)

	require.Equal(t, arg.ID, orderItem.ID)
	require.Equal(t, arg.OrderID, orderItem.OrderID)
	require.Equal(t, arg.ProductID, orderItem.ProductID)
	require.Equal(t, arg.Quantity, orderItem.Quantity)
	require.Equal(t, util.Preparing, orderItem.Status)

	require.NotZero(t, orderItem.CreatedAt)
	require.True(t, orderItem.UpdatedAt.IsZero())

	return orderItem
}

func TestGetPaymentRecord(t *testing.T) {
	randomPayment := CreatePayment(t)

	payment, err := testQueries.GetPaymentRecord(context.Background(), randomPayment.ID)
	require.NoError(t, err)
	require.NotEmpty(t, payment)

	require.Equal(t, randomPayment.ID, payment.ID)
	require.Equal(t, randomPayment.Type, payment.Type)
	require.Equal(t, randomPayment.CardNumber, payment.CardNumber)
	require.Equal(t, randomPayment.Status, payment.Status)

	require.WithinDuration(t, randomPayment.CreatedAt, payment.CreatedAt, time.Second)
	require.True(t, payment.UpdatedAt.IsZero())
}

func TestUpdatePaymentStatus(t *testing.T) {
	randomPayment := CreatePayment(t)

	updatedPayment, err := testQueries.UpdatePaymentStatus(context.Background(), UpdatePaymentStatusParams{
		ID:     randomPayment.ID,
		Status: util.Succeed,
	})

	require.NoError(t, err)
	require.NotEmpty(t, updatedPayment)

	require.Equal(t, randomPayment.ID, updatedPayment.ID)
	require.Equal(t, randomPayment.Amount, updatedPayment.Amount)
	require.Equal(t, randomPayment.Type, updatedPayment.Type)
	require.Equal(t, randomPayment.CardNumber, updatedPayment.CardNumber)
	require.Equal(t, util.Succeed, updatedPayment.Status)

	require.WithinDuration(t, randomPayment.CreatedAt, updatedPayment.CreatedAt, time.Second)
	require.NotZero(t, updatedPayment.UpdatedAt)
}

func TestGetOrderDetail(t *testing.T) {
	userCredential := CreateRandomUserCredential(t)
	userInfo := CreateRandomUserInfo(t, userCredential)
	payment := CreatePayment(t)

	numOrder := 5
	numItem := 10

	var lastOrder OrderDetail
	for i := 0; i < numOrder; i++ {
		lastOrder = CreateOrder(t, payment, userInfo)
		for j := 0; j < numItem; j++ {
			_ = CreateOrderItem(t, lastOrder)
		}
	}

	orderSummary, err := testQueries.GetUserOrderSummary(context.Background(), GetUserOrderSummaryParams{
		UserID: userInfo.ID,
		Limit:  int32(numOrder),
		Offset: 0,
	})
	require.NoError(t, err)
	require.NotEmpty(t, orderSummary)

	require.Equal(t, numOrder, len(orderSummary))

	found := false
	for _, order := range orderSummary {
		require.Equal(t, numItem, int(order.NumberProduct))
		if order.ID == lastOrder.ID {
			found = true
			require.Equal(t, lastOrder.Total, order.Total.Float64)
		}
	}
	require.True(t, found)
}

func TestGetOrderItemList(t *testing.T) {
	userCredential := CreateRandomUserCredential(t)
	userInfo := CreateRandomUserInfo(t, userCredential)
	payment := CreatePayment(t)

	numItem := 10
	order := CreateOrder(t, payment, userInfo)
	var lastItem OrderItem
	for j := 0; j < numItem; j++ {
		lastItem = CreateOrderItem(t, order)
	}

	arg := GetOrderItemListParams{
		OrderID: order.ID,
		Limit:   int32(numItem),
		Offset:  0,
	}
	orderList, err := testQueries.GetOrderItemList(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, orderList)

	require.Equal(t, numItem, len(orderList))

	found := false
	for _, item := range orderList {
		require.Equal(t, order.ID, item.OrderID)
		if item.ID == lastItem.ID {
			found = true
			require.Equal(t, lastItem.Quantity, item.Quantity)
			require.Equal(t, lastItem.Status, item.Status)
			require.Equal(t, lastItem.ProductID, item.ProductID)
			require.WithinDuration(t, lastItem.CreatedAt, item.CreatedAt, time.Second)
		}
	}
	require.True(t, found)
}
