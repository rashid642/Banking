package api

import (
	"errors"
	"net/http"
	"strings"

	token "github.com/rashid642/banking/token"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey = "authorization"
	authorizationTypeBeare = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey) 
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization headers not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return 
		}
		feilds := strings.Fields(authorizationHeader) 
		if len(feilds) < 2 {
			err := errors.New("invalid Authorization headers not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return 
		}

		authorizationType := strings.ToLower(feilds[0]) 
		if authorizationType != authorizationTypeBeare {
			err := errors.New("invalid authorization Type")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return 
		}

		accessToken := feilds[1] 
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return 
		}
		
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}