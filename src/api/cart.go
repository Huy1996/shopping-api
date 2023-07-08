package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"net/http"
	db "shopping-cart/src/db/sqlc"
	"shopping-cart/src/token"
	"strconv"
)

// Add Item To Cart Request
type addItemToCartRequest struct {
	CartID    uuid.UUID `json:"cart_id" binding:"required"`
	ProductID uuid.UUID `json:"product_id" binding:"required"`
	Quantity  int32     `json:"quantity" binding:"required,gte=1"`
}

func (server *Server) addItemToCart(ctx *gin.Context) {
	var req addItemToCartRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.AddToCartTxParam{
		CartID:    req.CartID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}
	result, err := server.store.AddToCartTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, result)
}

type getCartItemListResponse struct {
	Items []db.GetCartProductDetailListRow `json:"items"`
	Total float64                          `json:"total"`
}

func (server *Server) getCartItemList(ctx *gin.Context) {
	payload := ctx.Keys[authorizationPayloadKey].(*token.Payload)

	page, err := strconv.ParseInt(ctx.DefaultQuery("page", "0"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(fmt.Errorf("page must be number")))
		return
	}

	arg := db.GetCartProductDetailListParams{
		CartID: payload.CartID,
		Limit:  server.config.LimitItemDisplay,
		Offset: int32(page) * server.config.LimitItemDisplay,
	}
	result, err := server.store.GetCartProductDetailList(ctx, arg)
	if err != nil {
		if pqerr, ok := err.(*pq.Error); ok {
			if pqerr.Code == pq.ErrorCode("2201X") {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
				return
			}
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	total, err := server.store.GetTotal(ctx, payload.CartID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := getCartItemListResponse{
		Items: result,
		Total: total,
	}

	ctx.JSON(http.StatusOK, res)
}
