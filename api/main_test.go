package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// Gin is running in debug mode by default so we need to do this

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run()) 
}