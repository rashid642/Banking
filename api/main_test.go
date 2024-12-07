package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	Database "github.com/rashid642/banking/Database/sqlc"
	"github.com/rashid642/banking/utils"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store Database.Store) *Server{ 
	config := utils.Config{
		TokenSymmetricKey: utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	server, err := NewServer(config, store) 
	require.NoError(t, err) 
	return server
}

// Gin is running in debug mode by default so we need to do this

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run()) 
}