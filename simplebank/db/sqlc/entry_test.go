package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.con/khaikenly/golang-postgresql/util"
)

func RandomCreateEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomBalance(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.Amount, entry.Amount)
	require.Equal(t, arg.AccountID, entry.AccountID)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := RamdomCreateAccount(t)
	RandomCreateEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := RamdomCreateAccount(t)
	entry1 := RandomCreateEntry(t, account)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntry(t *testing.T) {
	account := RamdomCreateAccount(t)
	for i := 0; i < 10; i++ {
		RandomCreateEntry(t, account)
	}

	arg := ListEntriesParams{
		Limit:     5,
		Offset:    5,
		AccountID: account.ID,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}
}
