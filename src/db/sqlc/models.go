// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type AddressBook struct {
	ID          uuid.UUID `json:"id"`
	Owner       uuid.UUID `json:"owner"`
	AddressName string    `json:"address_name"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Zipcode     int32     `json:"zipcode"`
	AddedAt     time.Time `json:"added_at"`
}

type UserCredential struct {
	ID                uuid.UUID `json:"id"`
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	IsAdmin           bool      `json:"is_admin"`
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