package dbcache

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestSet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectBegin()

}
