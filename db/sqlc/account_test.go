package db

import (
	"GoBank/util"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

// 没有Test前缀的不会在测试时运行
// 创建用户 测试之间最好独立运行，所以每个测试最好单独创建测试用例
func CreateRandomAccount(t *testing.T) Account {
	user := CreateRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  float64(util.RandomBalance()),
		Currency: util.RandomCurrency(),
	}
	// 调用函数进行测试
	acc, err := testQueries.CreateAccount(context.Background(), arg)
	// 使用 testify 判断操作是否成功
	// 判断err
	require.NoError(t, err)

	// 判断结果不为空
	require.NotEmpty(t, acc)

	// 判断输入和查询相等
	require.Equal(t, arg.Owner, acc.Owner)
	require.Equal(t, arg.Balance, acc.Balance)
	require.Equal(t, arg.Currency, acc.Currency)

	// 判断其他自动添加的字段非0
	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)

	return acc
}

// 测试创建用户函数CreateAccount
func TestQueries_CreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

// 测试查询用户函数GetAccount
func TestQueries_GetAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)

	// 比较时间是否相同，最后一个参数为允许的误差
	require.WithinDuration(t, account1.CreatedAt.Time, account2.CreatedAt.Time, time.Second)
}

// 测试更新函数UpdateAccount
func TestQueries_UpdateAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: float64(util.RandomBalance()),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt.Time, account2.CreatedAt.Time, time.Second)
}

// 测试删除函数DeleteAccount
func TestQueries_DeleteAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	// 比较错误是否一样
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, account2)
}

// 测试查询函数ListAccount
func TestQueries_ListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomAccount(t)
	}

	arg := ListAccountParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccount(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
