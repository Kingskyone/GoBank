package gapi

import (
	db "GoBank/db/sqlc"
	"GoBank/pb"
	"GoBank/util"
	"context"
	"database/sql"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	//return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")

	// 获取用户
	user, err := server.store.GetUser(ctx, req.GetUsername())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "用户不存在", err)
		}
		return nil, status.Errorf(codes.Internal, "查询用户失败", err)
	}
	// 检查密码是否正确
	err = util.CheckPassword(req.GetPassword(), user.HashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "密码错误", err)
	}
	// 正确则给予token
	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Token创建失败", err)
	}

	// 创建持久化token
	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "持久化Token创建失败", err)
	}

	ID := pgtype.UUID{
		Bytes: refreshPayload.ID,
		Valid: true,
	}

	var ExpiresAt pgtype.Timestamptz
	err = ExpiresAt.Scan(refreshPayload.ExpiredAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "时间格式转换错误", err)
	}

	_, err = server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    "ctx.Request.UserAgent()",
		ClientIp:     "ctx.ClientIP()",
		IsBlocked:    false,
		ExpiresAt:    ExpiresAt,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "创建会话失败", err)
	}
	rsp := &pb.LoginUserResponse{
		User:                 convertUser(user),
		SessionId:            refreshPayload.ID.String(),
		AccessToken:          accessToken,
		RefreshToken:         refreshToken,
		AccessTokenExpireAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpireAt: timestamppb.New(refreshPayload.ExpiredAt),
	}
	return rsp, nil
}
