package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashedPassword, err := HashPassowrd(password)
	require.NoError(t, err) 
	require.NotEmpty(t, hashedPassword) 

	err = CheckPassowrd(password, hashedPassword) 
	require.NoError(t, err) 

	wrongPassword := RandomString(6)
	err = CheckPassowrd(wrongPassword, hashedPassword) 
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassword2, err := HashPassowrd(password)
	require.NoError(t, err) 
	require.NotEmpty(t, hashedPassword) 
	require.NotEqual(t, hashedPassword, hashedPassword2)
}