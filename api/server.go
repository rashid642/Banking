package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/rashid642/banking/Database/sqlc"
	"github.com/rashid642/banking/token"
	"github.com/rashid642/banking/utils"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store db.Store 
	tokenMaker token.Maker
	router *gin.Engine
	config utils.Config
}

// New Server Instance
func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokeMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("can't create token maker: %v", err)
	}
	server := &Server{
		config: config,
		store: store,
		tokenMaker: tokeMaker,
	} 
	
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil 
}

func (server *Server) setupRouter(){
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.getAccountList)

	authRoutes.POST("/transfer", server.createTransfer)

	server.router = router
}

// Start runs to HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"Error" : err.Error()}
}