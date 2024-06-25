package db

import (
	"GoBank/util"
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateRandomUser(t *testing.T) User {
	//hashedPassword, err := util.HashPassword(util.RandomString(6))
	hashedPassword, err := util.HashPassword("secret")
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	// 调用函数进行测试
	user, err := testQueries.CreateUser(context.Background(), arg)
	// 使用 testify 判断操作是否成功
	// 判断err
	require.NoError(t, err)

	// 判断结果不为空
	require.NotEmpty(t, user)

	// 判断输入和查询相等
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	//require.True(t, user.PasswordChangeAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

// 测试CreateUser
func TestQueries_CreateUser(t *testing.T) {
	CreateRandomUser(t)
}

// 测试GetUser
func TestQueries_GetUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)

	// 比较时间是否相同，最后一个参数为允许的误差
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
	require.WithinDuration(t, user1.PasswordChangeAt.Time, user2.PasswordChangeAt.Time, time.Second)
}

func TestQueries_UpdateUserWithFullName(t *testing.T) {
	oldUser := CreateRandomUser(t)
	newFullName := util.RandomOwner()
	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		FullName: pgtype.Text{String: newFullName, Valid: true},
		Username: oldUser.Username,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)
	require.NotEqual(t, oldUser.FullName, updatedUser.FullName)
	require.Equal(t, updatedUser.FullName, newFullName)
	require.Equal(t, updatedUser.HashedPassword, oldUser.HashedPassword)
	require.Equal(t, updatedUser.Email, oldUser.Email)
}
