-- name: CreatePaymentRecord :one
INSERT INTO payment_detail (
    id,
    amount,
    type,
    status,
    card_number
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetPaymentRecord :one
SELECT * FROM payment_detail
WHERE id = $1
LIMIT 1;

-- name: UpdatePaymentStatus :one
UPDATE payment_detail
SET
    status = $1,
    updated_at = now()
WHERE id = $2
RETURNING *;