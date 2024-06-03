package gapi

import (
	db "GoBank/db/sqlc"
	"GoBank/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// 将db.user转换为pb.user
func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:         user.Username,
		FullName:         user.FullName,
		Email:            user.Email,
		PasswordChangeAt: timestamppb.New(user.PasswordChangeAt.Time),
		CreatedAt:        timestamppb.New(user.CreatedAt.Time),
	}
}
