package db

import (
	"GoBank/util"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateRandomTransfer(t *testing.T) Transfer {
	arg := CreateTransferParams{
		FromAccountID: 1,
		ToAccountID:   2,
		Money:         float64(util.RandomBalance()),
	}
	tsf, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, tsf)
	require.Equal(t, arg.Money, tsf.Money)
	require.Equal(t, arg.ToAccountID, tsf.ToAccountID)
	require.Equal(t, arg.FromAccountID, tsf.FromAccountID)
	require.NotZero(t, tsf.ID)
	require.NotZero(t, tsf.CreatedAt)

	return tsf
}

func TestQueries_CreateTransfer(t *testing.T) {
	CreateRandomTransfer(t)
}

func TestQueries_GetTransfer(t *testing.T) {
	transfer1 := CreateRandomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.Money, transfer2.Money)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.WithinDuration(t, transfer1.CreatedAt.Time, transfer2.CreatedAt.Time, time.Second)
}

func TestQueries_UpdateTransfer(t *testing.T) {
	transfer1 := CreateRandomTransfer(t)
	arg := UpdateTransferParams{
		ID:    transfer1.ID,
		Money: float64(util.RandomBalance()),
	}
	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, arg.Money, transfer2.Money)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.WithinDuration(t, transfer1.CreatedAt.Time, transfer2.CreatedAt.Time, time.Second)
}

func TestQueries_DeleteTransfer(t *testing.T) {
	transfer1 := CreateRandomTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, transfer2)
}

func TestQueries_ListTransferWithID(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomTransfer(t)
	}

	arg := ListTransferWithIDParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransferWithID(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}

}

func TestQueries_ListTransferWithFromAccountID(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomTransfer(t)
	}

	arg := ListTransferWithFromAccountIDParams{
		FromAccountID: 1,
		Limit:         5,
		Offset:        5,
	}

	accounts, err := testQueries.ListTransferWithFromAccountID(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}

}

func TestQueries_ListTransferWithToAccountID(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomTransfer(t)
	}

	arg := ListTransferWithToAccountIDParams{
		ToAccountID: 2,
		Limit:       5,
		Offset:      5,
	}

	accounts, err := testQueries.ListTransferWithToAccountID(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}

}
