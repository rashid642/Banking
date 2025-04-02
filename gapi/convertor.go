package gapi

import (
	Database "github.com/rashid642/banking/Database/sqlc"
	"github.com/rashid642/banking/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user Database.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		Email:             user.Email,
		FullName:          user.FullName,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt.Time),
	}
}


