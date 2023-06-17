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
