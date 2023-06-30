package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "shopping-cart/src/db/sqlc"
	"shopping-cart/src/util"
	"time"
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
	UserName    string    `json:"user_name"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	MiddleName  string    `json:"middle_name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
}

func newUserResponse(userCredential db.UserCredential, userInfo db.UserInfo) UserResponse {
	return UserResponse{
		UserName:    userCredential.Username,
		FirstName:   userInfo.FirstName,
		LastName:    userInfo.LastName,
		MiddleName:  userInfo.MiddleName,
		PhoneNumber: userInfo.PhoneNumber,
		Email:       userCredential.Email,
		CreatedAt:   userCredential.CreatedAt,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	// TODO: Add new route function
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserTxParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		Email:          req.Email,
		PhoneNumber:    req.PhoneNumber,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		MiddleName:     req.MiddleName,
	}

	result, err := server.store.CreateUserTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newUserResponse(result.UserCredential, result.UserInfo))
}
