package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/rashid642/banking/Database/sqlc"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store db.Store 
	router *gin.Engine
}

// New Server Instance
func NewServer(store db.Store) *Server {
	server := &Server{store: store} 
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.getAccountList)

	router.POST("/transfer", server.createTransfer)

	server.router = router
	return server 
}

// Start runs to HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"Error" : err.Error()}
}