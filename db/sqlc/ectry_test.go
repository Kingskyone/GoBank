package db

import (
	"GoBank/util"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateRandomEctry(t *testing.T) Ectry {
	arg := CreateEntryParams{
		AccountID: 1,
		Amount:    float64(util.RandomAmount()),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.Amount, entry.Amount)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestQueries_CreateEctry(t *testing.T) {
	CreateRandomEctry(t)
}

func TestQueries_GetEctry(t *testing.T) {
	entry1 := CreateRandomEctry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.WithinDuration(t, entry1.CreatedAt.Time, entry2.CreatedAt.Time, time.Second)
}

func TestQueries_UpdateEctry(t *testing.T) {
	entry1 := CreateRandomEctry(t)
	fmt.Println(entry1)

	arg := UpdateEntryParams{
		ID:     entry1.ID,
		Amount: float64(util.RandomAmount()),
	}

	entry2, err := testQueries.UpdateEntry(context.Background(), arg)
	fmt.Println(entry2)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, arg.Amount, entry2.Amount)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.WithinDuration(t, entry1.CreatedAt.Time, entry2.CreatedAt.Time, time.Second)
}

func TestQueries_DeleteEctry(t *testing.T) {
	entry1 := CreateRandomEctry(t)
	err := testQueries.DeleteEntry(context.Background(), entry1.ID)
	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, entry2)

}

func TestQueries_ListEctryWithID(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomEctry(t)
	}

	arg := ListEntryWithIDParams{
		Limit:  5,
		Offset: 5,
	}

	entrys, err := testQueries.ListEntryWithID(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entrys, 5)
	for _, entry := range entrys {
		require.NotEmpty(t, entry)
	}
}

func TestQueries_ListEctryWithAccountID(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomEctry(t)
	}

	arg := ListEntryWithAccountIDParams{
		AccountID: 1,
		Limit:     5,
		Offset:    5,
	}

	entrys, err := testQueries.ListEntryWithAccountID(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entrys, 5)
	for _, entry := range entrys {
		require.NotEmpty(t, entry)
	}
}
