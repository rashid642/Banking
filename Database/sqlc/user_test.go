package Database 

import (
	"context"
	"testing"

	"github.com/rashid642/banking/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := utils.HashPassowrd(utils.RandomString(6))
	require.NoError(t, err) 

	arg := CreateUserParams{
		Username: utils.RandomOwner(),
		HashedPassowrd: hashedPassword,
		FullName: utils.RandomOwner(),
		Email: utils.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg) 
	
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassowrd, user.HashedPassowrd)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t) 
	user2, err := testQueries.GetUser(context.Background(), user.Username)

	require.NoError(t, err)

	require.NotEmpty(t, user2) 
	require.Equal(t, user.Username, user2.Username)
	require.Equal(t, user.HashedPassowrd, user2.HashedPassowrd)
	require.Equal(t, user.Email, user2.Email) 
	require.Equal(t, user.FullName, user2.FullName)
	require.Equal(t, user.CreatedAt, user2.CreatedAt) 
}
