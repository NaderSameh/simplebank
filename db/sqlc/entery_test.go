package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, acc Account) Entry {

	var i int64 = 10
	arg := CreateEntryParams{
		AccountID: acc.ID,
		Amount:    i,
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, acc.ID)
	require.Equal(t, entry.Amount, i)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry

}
func TestCreateEntry(t *testing.T) {
	acc := createRandomAccount(t)
	createRandomEntry(t, acc)

}

func TestGetEntry(t *testing.T) {
	acc := createRandomAccount(t)
	entery := createRandomEntry(t, acc)
	entery2, err := testQueries.GetEntry(context.Background(), entery.ID)

	require.NoError(t, err)
	require.Equal(t, entery.ID, entery2.ID)
	require.Equal(t, entery.Amount, entery2.Amount)
	require.Equal(t, entery.AccountID, entery2.AccountID)
	require.WithinDuration(t, entery.CreatedAt, entery2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	acc := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, acc)
	}

	arg := ListEntriesParams{
		AccountID: acc.ID,
		Limit:     5,
		Offset:    5,
	}
	enteries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)

	require.Len(t, enteries, 5)

	for _, entery := range enteries {
		require.NotEmpty(t, entery)
	}

}
