package checkpoint_postgres

import "testing"

func TestPostgresSaver(t *testing.T) {
	s := PostgresSaver{}
	next, err := s.Save([]byte("d"), 0)
	if err != nil || next != 1 {
		t.Fatalf("unexpected result: %v %d", err, next)
	}
}
