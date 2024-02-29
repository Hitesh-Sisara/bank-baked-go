package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Hitesh-Sisara/bank-app-go/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, from, to int64) Transfer {
	arg := CreateTransferParams{
		FromAccountID: from,
		ToAccountID:   to,
		Amount:        util.RandomMoney(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	from := CreateRandomAccount(t)
	to := CreateRandomAccount(t)
	createRandomTransfer(t, from.ID, to.ID)
}

func TestGetTransfer(t *testing.T) {
	from := CreateRandomAccount(t)
	to := CreateRandomAccount(t)
	transfer1 := createRandomTransfer(t, from.ID, to.ID)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	from := CreateRandomAccount(t)
	to := CreateRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomTransfer(t, from.ID, to.ID)
	}
	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}
	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}

func TestDeleteTransfer(t *testing.T) {
	from := CreateRandomAccount(t)
	to := CreateRandomAccount(t)
	transfer1 := createRandomTransfer(t, from.ID, to.ID)
	err := testQueries.DeleteTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, transfer2)
}
