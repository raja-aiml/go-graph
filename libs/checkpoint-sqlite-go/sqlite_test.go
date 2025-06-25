package checkpoint_sqlite

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestSQLiteSaver(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	s, err := NewSQLiteSaver(db)
	if err != nil {
		t.Fatalf("new saver: %v", err)
	}
	next, err := s.Save([]byte("d"), 0)
	if err != nil || next != 1 {
		t.Fatalf("unexpected result: %v %d", err, next)
	}
}
