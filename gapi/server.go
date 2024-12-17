package gapi

import (
	"fmt"

	Database "github.com/rashid642/banking/Database/sqlc"
	"github.com/rashid642/banking/pb"
	"github.com/rashid642/banking/token"
	"github.com/rashid642/banking/utils"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store Database.Store 
	// UnimplementedBankingServer - help us to implement rpcs in prallel and not blocking other rpc, 
	// as with out all the interface function implementation we can't start
	pb.UnimplementedBankingServer
	tokenMaker token.Maker
	config utils.Config
}

// New Server Instance
func NewServer(config utils.Config, store Database.Store) (*Server, error) {
	tokeMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("can't create token maker: %v", err)
	}
	server := &Server{
		config: config,
		store: store,
		tokenMaker: tokeMaker,
	} 

	return server, nil 
}


