package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.con/khaikenly/golang-postgresql/util"
)

func RandomCreateTransfer(t *testing.T, fromAccount, toAccount Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        util.RandomBalance(),
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
	fromAccount := RamdomCreateAccount(t)
	toAccount := RamdomCreateAccount(t)
	RandomCreateTransfer(t, fromAccount, toAccount)
}

func TestGetTransfer(t *testing.T) {
	fromAccount := RamdomCreateAccount(t)
	toAccount := RamdomCreateAccount(t)
	transfer1 := RandomCreateTransfer(t, fromAccount, toAccount)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)

}

func TestListTransfer(t *testing.T) {
	fromAccount := RamdomCreateAccount(t)
	toAccount := RamdomCreateAccount(t)

	for i := 0; i < 10; i++ {
		RandomCreateTransfer(t, fromAccount, toAccount)
	}

	arg := ListTransfersParams{
		FromAccountID: fromAccount.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, arg.FromAccountID, fromAccount.ID)
		require.True(t, transfer.FromAccountID == fromAccount.ID || transfer.ToAccountID == toAccount.ID)
	}
}
