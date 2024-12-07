package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	Database "github.com/rashid642/banking/Database/sqlc"
	token "github.com/rashid642/banking/token"
)

type transferRequest struct {
	FromAccountID int64 `json:"from_account_id" binding:"required,min=1"`
	ToAccountId   int64 `json:"to_account_id" binding:"required,min=1"`
	Amount        int64 `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest 
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload) 

	fromAccount, valid := server.validAccount(ctx, req.FromAccountID, req.Currency)
	if !valid{
		return 
	}

	if fromAccount.Owner != authPayload.Username {
		err := errors.New("account doesn't belong to this user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return 
	}

	_, valid = server.validAccount(ctx, req.ToAccountId, req.Currency)
	if !valid {
		return 
	}

	arg := Database.TransferTxParams{
		FromAccountId: req.FromAccountID,
		ToAccountId: req.ToAccountId,
		Amount: req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountId int64, currecy string) (Database.Account, bool) {
	account, err := server.store.GetAccount(ctx, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return account, false 
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return account, false 
	}

	if account.Currency != currecy {
		err := fmt.Errorf("account [%d] currecy mismatch %s vs %s", accountId, account.Currency, currecy) 
		ctx.JSON(http.StatusBadRequest, errorResponse(err)) 
		return account, false 
	}
	return account, true 
}