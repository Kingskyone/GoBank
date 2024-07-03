package gapi

import (
	db "GoBank/db/sqlc"
	"GoBank/pb"
	"GoBank/util"
	"GoBank/val"
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (server *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	//验证用户权限
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	// 进行数据验证
	if violations := validateUpdateUserRequest(req); violations != nil {
		return nil, invalidArgumentError(violations)
	}
	// 从req中获取字段时可以使用前面加 Get 的函数

	//验证用户和修改的用户是否对应
	if authPayload.Username != req.GetUsername() {
		return nil, status.Errorf(codes.PermissionDenied, "无法更新当前用户，用户名不匹配")
	}

	arg := db.UpdateUserParams{
		Username: req.GetUsername(),
		//HashedPassword: pgtype.Text{
		//	String: req.GetFullName(),
		//	Valid:  req.FullName != nil,
		//},
		FullName: pgtype.Text{
			String: req.GetFullName(),
			Valid:  req.FullName != nil,
		},
		Email: pgtype.Text{
			String: req.GetEmail(),
			Valid:  req.Email != nil,
		},
	}
	if req.Password != nil {
		hashedPassword, err := util.HashPassword(req.GetPassword())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "密码hash失败: %s", err)
		}
		arg.HashedPassword = pgtype.Text{
			String: hashedPassword,
			Valid:  true,
		}
		arg.PasswordChangeAt = pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		}
	}

	// 插入数据
	user, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return nil, status.Errorf(codes.AlreadyExists, "用户不存在", err)
		}
		return nil, status.Errorf(codes.Internal, "更新用户失败", err)
	}
	rsp := &pb.UpdateUserResponse{
		User: convertUser(user),
	}
	return rsp, nil
}

// 验证UpdateUserRequest中的参数，返回一个错误信息表
func validateUpdateUserRequest(req *pb.UpdateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if req.FullName != nil {
		if err := val.ValidateFullName(req.GetFullName()); err != nil {
			violations = append(violations, fieldViolation("full_name", err))
		}
	}

	if req.Email != nil {
		if err := val.ValidateEmail(req.GetEmail()); err != nil {
			violations = append(violations, fieldViolation("email", err))
		}
	}

	if req.Password != nil {
		if err := val.ValidatePassword(req.GetPassword()); err != nil {
			violations = append(violations, fieldViolation("password", err))
		}
	}

	return violations
}
