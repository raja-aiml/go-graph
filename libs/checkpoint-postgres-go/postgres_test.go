package checkpoint_postgres

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestPostgresSaver(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("mock init: %v", err)
	}
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS checkpoints").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("INSERT INTO checkpoints").WithArgs(1, []byte("d")).WillReturnResult(sqlmock.NewResult(1, 1))
	s, err := NewPostgresSaver(db)
	if err != nil {
		t.Fatalf("new saver: %v", err)
	}
	next, err := s.Save([]byte("d"), 0)
	if err != nil || next != 1 {
		t.Fatalf("unexpected result: %v %d", err, next)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}
