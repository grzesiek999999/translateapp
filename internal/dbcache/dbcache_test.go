package dbcache_test

import (
	"context"
	"errors"
	"testing"
	"time"
	"translateapp/internal/dbcache"
	"translateapp/internal/logging"

	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func TestGetAfterTTL(t *testing.T) {
	const key = "key"
	const expected = "value"

	var ErrDoesNotExist = errors.New("ErrDoesNotExist")

	tt := time.Now().Add(-1 * time.Second)

	mock, err := pgxmock.NewConn()
	require.NoError(t, err)
	defer mock.Close(context.Background())

	rows := mock.NewRows([]string{"value", "timeout"}).AddRow(expected, tt)
	mock.ExpectQuery("SELECT value, timeout from cache WHERE").WithArgs(key).WillReturnRows(rows)
	mock.ExpectExec("DELETE FROM cache WHERE").WithArgs(key).WillReturnResult(pgxmock.NewResult("DELETE", 1))

	repo := dbcache.NewRepo(mock, logging.DefaultLogger().Desugar())
	cache := dbcache.NewDBCache(repo, logging.DefaultLogger().Desugar())

	_, err = cache.Get(key)
	require.Error(t, err)
	require.Equal(t, ErrDoesNotExist, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetBeforeTTL(t *testing.T) {
	const key = "key"
	const expected = "value"

	tt := time.Now().Add(30 * time.Second)

	mock, err := pgxmock.NewConn()
	require.NoError(t, err)
	defer mock.Close(context.Background())

	rows := mock.NewRows([]string{"value", "timeout"}).AddRow(expected, tt)
	mock.ExpectQuery("SELECT value, timeout from cache WHERE").WithArgs(key).WillReturnRows(rows)

	repo := dbcache.NewRepo(mock, logging.DefaultLogger().Desugar())

	cache := dbcache.NewDBCache(repo, logging.DefaultLogger().Desugar())
	val, err := cache.Get(key)
	require.NoError(t, err)
	require.Equal(t, expected, val)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetError(t *testing.T) {
	const key = "key"

	mock, err := pgxmock.NewConn()
	require.NoError(t, err)
	defer mock.Close(context.Background())

	mock.ExpectQuery("SELECT value, timeout from cache WHERE").WithArgs(key).WillReturnError(errors.New("error"))

	repo := dbcache.NewRepo(mock, logging.DefaultLogger().Desugar())

	cache := dbcache.NewDBCache(repo, logging.DefaultLogger().Desugar())
	val, err := cache.Get(key)
	require.Error(t, err)
	require.Equal(t, "", val)

	require.NoError(t, mock.ExpectationsWereMet())
}

//func TestSet(t *testing.T) {
//	const key = "test10"
//	const value = "key"
//
//	expected := true
//
//	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@db:5432/postgres")
//	require.NoError(t, err)
//	defer conn.Close(context.Background())
//	result := conn.QueryRow(context.Background(), "select exists(select 1 from cache where key='test10')")
//	require.NoError(t, err)
//	require.Equal(t, expected, result)
//}
