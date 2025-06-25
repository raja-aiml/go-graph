package checkpoint_sqlite

import "testing"

func TestSQLiteSaver(t *testing.T) {
	s := SQLiteSaver{}
	next, err := s.Save([]byte("d"), 0)
	if err != nil || next != 1 {
		t.Fatalf("unexpected result: %v %d", err, next)
	}
}
