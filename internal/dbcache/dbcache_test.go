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

func TestSet(t *testing.T) {
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

func TestSetError(t *testing.T) {
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
