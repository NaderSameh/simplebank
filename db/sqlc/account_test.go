package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/nadersameh_/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)

}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestGetAccount(t *testing.T) {
	accountDemo := createRandomAccount(t)
	account, err := testQueries.GetAccount(context.Background(), accountDemo.ID)
	require.NoError(t, err)
	require.Equal(t, account, accountDemo)
}

func TestAddAccountBalance(t *testing.T) {
	accountDemo := createRandomAccount(t)
	arg := AddAccountBalanceParams{
		Amount: 40,
		ID:     accountDemo.ID,
	}
	acc, err := testQueries.AddAccountBalance(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, accountDemo.Balance+40, acc.Balance)

	require.Equal(t, accountDemo.Owner, acc.Owner)
	require.Equal(t, accountDemo.Currency, acc.Currency)

	require.Equal(t, accountDemo.ID, acc.ID)
	require.Equal(t, accountDemo.CreatedAt, acc.CreatedAt)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 1,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func TestGetAccountForUpdate(t *testing.T) {
	accountDemo := createRandomAccount(t)

	acc, err := testQueries.GetAccountForUpdate(context.Background(), accountDemo.ID)

	require.NoError(t, err)
	require.Equal(t, acc.ID, accountDemo.ID)
	require.Equal(t, acc.Balance, accountDemo.Balance)
	require.Equal(t, acc.CreatedAt, accountDemo.CreatedAt)
	require.Equal(t, acc.Currency, accountDemo.Currency)
}

func TestUpdateAccount(t *testing.T) {
	accountDemo := createRandomAccount(t)
	X := 20
	arg := UpdateAccountParams{
		ID:      accountDemo.ID,
		Balance: int64(X),
	}

	acc, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, acc.ID, accountDemo.ID)
	require.Equal(t, acc.Owner, accountDemo.Owner)
	require.WithinDuration(t, acc.CreatedAt, accountDemo.CreatedAt, time.Second)
	require.Equal(t, acc.Currency, accountDemo.Currency)
	require.Equal(t, acc.Balance, int64(X))
}
