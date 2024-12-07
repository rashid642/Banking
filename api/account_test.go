package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mock_sqlc "github.com/rashid642/banking/Database/mock"
	Database "github.com/rashid642/banking/Database/sqlc"
	"github.com/rashid642/banking/utils"
	"github.com/stretchr/testify/require"
)

func TestGetaccountAPI(t *testing.T) {
	account := randomAccount() 

	testCases := []struct {
		name string 
		accountId int64 
		buildStubs func (store *mock_sqlc.MockStore) 
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name : "OK",
			accountId: account.ID,
			buildStubs: func (store *mock_sqlc.MockStore) {
				store.EXPECT().
				GetAccount(gomock.Any(), gomock.Eq(account.ID)).
				Times(1). // expect to call atleast once 
				Return(account, nil) 
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
		{
			name : "NotFound",
			accountId: account.ID,
			buildStubs: func (store *mock_sqlc.MockStore) {
				store.EXPECT().
				GetAccount(gomock.Any(), gomock.Eq(account.ID)).
				Times(1). // expect to call atleast once 
				Return(Database.Account{}, sql.ErrNoRows) 
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name : "InternalError",
			accountId: account.ID,
			buildStubs: func (store *mock_sqlc.MockStore) {
				store.EXPECT().
				GetAccount(gomock.Any(), gomock.Eq(account.ID)).
				Times(1). // expect to call atleast once 
				Return(Database.Account{}, sql.ErrConnDone) 
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name : "InvalidId",
			accountId: 0,
			buildStubs: func (store *mock_sqlc.MockStore) {
				store.EXPECT().
				GetAccount(gomock.Any(), gomock.Any()).
				Times(0) // expect to call atleast once 
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i] 

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t) 
			defer ctrl.Finish()
			store := mock_sqlc.NewMockStore(ctrl) 
			tc.buildStubs(store) 
			
			server := newTestServer(t, store) 
			recorder := httptest.NewRecorder() 
	
			url := fmt.Sprintf("/accounts/%d", tc.accountId) 
			req, err := http.NewRequest(http.MethodGet, url, nil) 
			require.NoError(t, err)
	
			server.router.ServeHTTP(recorder, req) 

			tc.checkResponse(t, recorder)
		})

	}
}

func randomAccount() Database.Account {
	return Database.Account{
		ID : utils.RandomInt(1, 1000),
		Owner: utils.RandomOwner(),
		Balance: utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account Database.Account) {
	data, err := ioutil.ReadAll(body) 
	require.NoError(t, err)

	var getAccount Database.Account
	err = json.Unmarshal(data, &getAccount) 
	require.NoError(t, err) 
	require.Equal(t, account, getAccount)
}