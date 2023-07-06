package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	mockdb "shopping-cart/src/db/mock"
	db "shopping-cart/src/db/sqlc"
	"shopping-cart/src/util"
	"testing"
)

type eqCreateUserParamsMatcher struct {
	arg      db.CreateUserTxParams
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	// In case some value is nil
	arg, ok := x.(db.CreateUserTxParams)
	if !ok {
		return false
	}

	err := util.CheckPassword(e.password, arg.HashedPassword)
	if err != nil {
		return false
	}

	e.arg.HashedPassword = arg.HashedPassword
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserTxParam(arg db.CreateUserTxParams, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

func TestCreateUserAPI(t *testing.T) {
	userInfo, userCredential, password := randomAccount(t)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username":     userCredential.Username,
				"password":     password,
				"email":        userCredential.Email,
				"phone_number": userInfo.PhoneNumber,
				"first_name":   userInfo.FirstName,
				"last_name":    userInfo.LastName,
				"middle_name":  userInfo.MiddleName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateUserTxParams{
					Username:       userCredential.Username,
					HashedPassword: userCredential.HashedPassword,
					Email:          userCredential.Email,
					PhoneNumber:    userInfo.PhoneNumber,
					FirstName:      userInfo.FirstName,
					LastName:       userInfo.LastName,
					MiddleName:     userInfo.MiddleName,
				}

				store.EXPECT().
					CreateUserTx(gomock.Any(), EqCreateUserTxParam(arg, password)).
					Times(1).
					Return(db.CreateUserTxResult{
						UserCredential: userCredential,
						UserInfo:       userInfo,
					}, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchCreateUser(t, recorder.Body, userCredential, userInfo)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"username":     userCredential.Username,
				"password":     password,
				"email":        userCredential.Email,
				"phone_number": userInfo.PhoneNumber,
				"first_name":   userInfo.FirstName,
				"last_name":    userInfo.LastName,
				"middle_name":  userInfo.MiddleName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.CreateUserTxResult{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "DuplicateUserName",
			body: gin.H{
				"username":     userCredential.Username,
				"password":     password,
				"email":        userCredential.Email,
				"phone_number": userInfo.PhoneNumber,
				"first_name":   userInfo.FirstName,
				"last_name":    userInfo.LastName,
				"middle_name":  userInfo.MiddleName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.CreateUserTxResult{}, &pq.Error{Code: "23505"})
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
			},
		},
		{
			name: "InvalidUserName",
			body: gin.H{
				"username":     "invalid-user#1",
				"password":     password,
				"email":        userCredential.Email,
				"phone_number": userInfo.PhoneNumber,
				"first_name":   userInfo.FirstName,
				"last_name":    userInfo.LastName,
				"middle_name":  userInfo.MiddleName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "PasswordTooShort",
			body: gin.H{
				"username":     userCredential.Username,
				"password":     "123",
				"email":        userCredential.Email,
				"phone_number": userInfo.PhoneNumber,
				"first_name":   userInfo.FirstName,
				"last_name":    userInfo.LastName,
				"middle_name":  userInfo.MiddleName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidEmail",
			body: gin.H{
				"username":     userCredential.Username,
				"password":     password,
				"email":        "invalid-email",
				"phone_number": userInfo.PhoneNumber,
				"first_name":   userInfo.FirstName,
				"last_name":    userInfo.LastName,
				"middle_name":  userInfo.MiddleName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
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

			// Marshal body data to JSON
			data, err := json.Marshal(testCase.body)
			require.NoError(t, err)

			url := fmt.Sprintf("/users")
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			testCase.checkResponse(recorder)
		})
	}

}

func randomAccount(t *testing.T) (userInfo db.UserInfo, userCredential db.UserCredential, password string) {
	password = util.RandomString(10)

	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	userCredentialID, err := uuid.NewRandom()
	require.NoError(t, err)

	userInfoID, err := uuid.NewRandom()
	require.NoError(t, err)

	userCredential = db.UserCredential{
		ID:             userCredentialID,
		Username:       util.RandomName(),
		HashedPassword: hashedPassword,
		Email:          util.RandomEmail(),
	}

	userInfo = db.UserInfo{
		ID:          userInfoID,
		UserID:      userCredentialID,
		PhoneNumber: util.RandomPhoneNumber(),
		FirstName:   util.RandomName(),
		LastName:    util.RandomName(),
		MiddleName:  util.RandomName(),
	}

	return
}

func requireBodyMatchCreateUser(t *testing.T, body *bytes.Buffer, userCredential db.UserCredential, userInfo db.UserInfo) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotInfo db.UserInfo
	var gotCredential db.UserCredential

	err = json.Unmarshal(data, &gotInfo)
	require.NoError(t, err)

	err = json.Unmarshal(data, &gotCredential)
	require.NoError(t, err)

	require.Equal(t, userCredential.Username, gotCredential.Username)
	require.Equal(t, userCredential.Email, gotCredential.Email)
	require.Empty(t, gotCredential.HashedPassword)

	require.Equal(t, userInfo.ID, gotInfo.ID)
	require.Equal(t, userInfo.FirstName, gotInfo.FirstName)
	require.Equal(t, userInfo.LastName, gotInfo.LastName)
	require.Equal(t, userInfo.MiddleName, gotInfo.MiddleName)
	require.Equal(t, userInfo.PhoneNumber, gotInfo.PhoneNumber)
}

func TestLoginUserAPI(t *testing.T) {
	userInfo, userCredential, password := randomAccount(t)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username": userCredential.Username,
				"password": password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserCredential(gomock.Any(), gomock.Eq(userCredential.Username)).
					Times(1).
					Return(userCredential, nil)
				store.EXPECT().
					GetUserInfoByUserID(gomock.Any(), gomock.Eq(userCredential.ID)).
					Times(1).
					Return(userInfo, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "UserNotFound",
			body: gin.H{
				"username": "NotFound",
				"password": password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserCredential(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.UserCredential{}, sql.ErrNoRows)
				store.EXPECT().
					GetUserInfoByUserID(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "IncorrectPassword",
			body: gin.H{
				"username": userCredential.Username,
				"password": "IncorrectPassword",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserCredential(gomock.Any(), gomock.Any()).
					Times(1).
					Return(userCredential, nil)
				store.EXPECT().
					GetUserInfoByUserID(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"username": userCredential.Username,
				"password": password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserCredential(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.UserCredential{}, sql.ErrConnDone)
				store.EXPECT().
					GetUserInfoByUserID(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidUsername",
			body: gin.H{
				"username": "invalid-user#1",
				"password": password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserCredential(gomock.Any(), gomock.Any()).
					Times(0)
				store.EXPECT().
					GetUserInfoByUserID(gomock.Any(), gomock.Any()).
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

			// Marshal body data to JSON
			data, err := json.Marshal(testCase.body)
			require.NoError(t, err)

			url := "/users/login"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			testCase.checkResponse(recorder)
		})
	}
}
