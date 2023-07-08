package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	mockdb "shopping-cart/src/db/mock"
	db "shopping-cart/src/db/sqlc"
	"shopping-cart/src/token"
	"shopping-cart/src/util"
	"testing"
	"time"
)

func TestAddToCartAPI(t *testing.T) {
	userInfo, _, userCart, _ := randomAccount(t)
	product, _, inventory := randomProduct(t)

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"cart_id":    userCart.ID,
				"product_id": product.ID,
				"quantity":   1,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, userInfo.ID, userCart.ID, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				id, err := uuid.NewRandom()
				require.NoError(t, err)

				arg := db.AddToCartTxParam{
					CartID:    userCart.ID,
					ProductID: product.ID,
					Quantity:  1,
				}
				store.EXPECT().
					AddToCartTx(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(db.AddToCartTxResult{
						Total: product.Price,
						CartItem: db.CartItem{
							ID:        id,
							CartID:    userCart.ID,
							ProductID: product.ID,
							Quantity:  1,
						},
					}, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "InvalidQty",
			body: gin.H{
				"cart_id":    userCart.ID,
				"product_id": product.ID,
				"quantity":   -1,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, userInfo.ID, userCart.ID, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					AddToCartTx(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InsufficientAmount",
			body: gin.H{
				"cart_id":    userCart.ID,
				"product_id": product.ID,
				"quantity":   inventory.Quantity + 1,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, userInfo.ID, userCart.ID, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					AddToCartTx(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.AddToCartTxResult{}, fmt.Errorf(""))

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		testCase := testCases[i]

		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			testCase.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(testCase.body)
			require.NoError(t, err)

			url := "/cart"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			testCase.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			testCase.checkResponse(recorder)
		})

	}
}

func TestGetCartItemListAPI(t *testing.T) {
	userInfo, _, userCart, _ := randomAccount(t)
	//_, _, _ := randomProduct(t)

	testCases := []struct {
		name          string
		param         string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, userInfo.ID, userCart.ID, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.GetCartProductDetailListParams{
					CartID: userCart.ID,
					Limit:  5,
					Offset: 0,
				}

				store.EXPECT().
					GetCartProductDetailList(gomock.Any(), gomock.Eq(arg)).
					Times(1)

				store.EXPECT().
					GetTotal(gomock.Any(), gomock.Eq(userCart.ID)).
					Times(1)

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:  "OKWithPage",
			param: "page=1",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, userInfo.ID, userCart.ID, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.GetCartProductDetailListParams{
					CartID: userCart.ID,
					Limit:  5,
					Offset: 5 * 1,
				}

				store.EXPECT().
					GetCartProductDetailList(gomock.Any(), gomock.Eq(arg)).
					Times(1)

				store.EXPECT().
					GetTotal(gomock.Any(), gomock.Eq(userCart.ID)).
					Times(1)

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:  "NegativePage",
			param: "page=-1",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, userInfo.ID, userCart.ID, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.GetCartProductDetailListParams{
					CartID: userCart.ID,
					Limit:  5,
					Offset: 5 * -1,
				}

				store.EXPECT().
					GetCartProductDetailList(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return([]db.GetCartProductDetailListRow{}, &pq.Error{Code: "2201X"})

				store.EXPECT().
					GetTotal(gomock.Any(), gomock.Eq(userCart.ID)).
					Times(0)

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:  "InvalidParams",
			param: "page=invalid",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, userInfo.ID, userCart.ID, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCartProductDetailList(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		testCase := testCases[i]

		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			testCase.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/cart?%v", testCase.param)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			testCase.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			testCase.checkResponse(recorder)
		})

	}
}

func randomProduct(t *testing.T) (db.Product, db.ProductCategory, db.ProductInventory) {
	productID, err := uuid.NewRandom()
	require.NoError(t, err)

	inventoryID, err := uuid.NewRandom()
	require.NoError(t, err)

	categoryID, err := uuid.NewRandom()
	require.NoError(t, err)

	category := db.ProductCategory{
		ID:          categoryID,
		Name:        util.RandomName(),
		Description: util.RandomString(100),
	}

	inventory := db.ProductInventory{
		ID:       inventoryID,
		Quantity: int32(util.RandomInt(1, 100)),
	}

	product := db.Product{
		ID:          productID,
		Name:        util.RandomName(),
		Description: util.RandomString(100),
		SKU:         util.RandomString(20),
		Price:       util.RandomFloat(1, 200),
		CategoryID:  categoryID,
		InventoryID: inventoryID,
	}

	return product, category, inventory
}
