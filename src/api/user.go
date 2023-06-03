package api

import (
	"github.com/gin-gonic/gin"
	db "shopping-cart/src/db/sqlc"
)

type createUserRequest struct {
	Username    string `json:"username" binding:"required,alphanum"`
	Password    string `json:"password" binding:"required,min=6"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required,min=10"`
	FirstName   string `json:"first_name" binding:"required,alpha"`
	LastName    string `json:"last_name" binding:"required,alpha"`
	MiddleName  string `json:"middle_name" binding:"required,alpha"`
}

type UserResponse struct {
	UserName    string `json:"user_name"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MiddleName  string `json:"middle_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"Email"`
	IsAdmin     bool   `json:"is_admin"`
}

func newUserResponse(userCredential db.UserCredential, userInfo db.UserInfo) UserResponse {
	return UserResponse{
		UserName:    userCredential.Username,
		FirstName:   userInfo.FirstName,
		LastName:    userInfo.LastName,
		MiddleName:  userInfo.MiddleName,
		PhoneNumber: userInfo.PhoneNumber,
		Email:       userCredential.Email,
		IsAdmin:     userCredential.IsAdmin,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	// TODO: Add new route function
}
