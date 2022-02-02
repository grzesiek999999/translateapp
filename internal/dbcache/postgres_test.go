package dbcache_test

import (
	"context"
	"errors"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
	"translateapp/internal/dbcache"
	"translateapp/internal/logging"
)

func TestSetPostgres(t *testing.T) {
	const key = "key"
	const value = "klucz"

	var ValueSet = errors.New("Value Set")
	tt := time.Now().Add(30 * time.Second)

	mock, err := pgxmock.NewConn()
	require.NoError(t, err)
	defer mock.Close(context.Background())

	mock.ExpectExec("INSERT INTO cache").WithArgs(key, value, tt).WillReturnResult(pgxmock.NewResult("INSERT", 1))

	repo := dbcache.NewRepo(mock, logging.DefaultLogger().Desugar())
	err = repo.Set(context.Background(), key, value, tt)
	
	require.Error(t, err)
	require.Equal(t, ValueSet, err)
	require.NoError(t, mock.ExpectationsWereMet())

}

func TestPostgresDelete(t *testing.T) {
	const key = "key"
	const expected = "value"

	var ValueDeleted = errors.New("Deleted")

	mock, err := pgxmock.NewConn()
	require.NoError(t, err)
	defer mock.Close(context.Background())

	mock.ExpectExec("DELETE FROM cache WHERE").WithArgs(key).WillReturnResult(pgxmock.NewResult("DELETE", 1))

	repo := dbcache.NewRepo(mock, logging.DefaultLogger().Desugar())
	err = repo.Delete(context.Background(), key)

	require.Error(t, err)
	require.Equal(t, ValueDeleted, err)
	require.NoError(t, mock.ExpectationsWereMet())
}
