package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	Database "github.com/rashid642/banking/Database/sqlc"
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
	if !server.validAccount(ctx, req.FromAccountID, req.Currency) {
		return 
	}
	if !server.validAccount(ctx, req.ToAccountId, req.Currency) {
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

func (server *Server) validAccount(ctx *gin.Context, accountId int64, currecy string) bool {
	account, err := server.store.GetAccount(ctx, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false 
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false 
	}

	if account.Currency != currecy {
		err := fmt.Errorf("Account [%d] currecy mismatch %s vs %s", accountId, account.Currency, currecy) 
		ctx.JSON(http.StatusBadRequest, errorResponse(err)) 
		return false 
	}
	return true 
}