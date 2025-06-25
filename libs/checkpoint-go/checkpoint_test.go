package checkpoint

import "testing"

type nopSaver struct{}

func (nopSaver) Save(data []byte, version int) (int, error) {
	return version + 1, nil
}

func TestSaver(t *testing.T) {
	var s Saver = nopSaver{}
	next, err := s.Save([]byte("data"), 1)
	if err != nil || next != 2 {
		t.Fatalf("unexpected result: %v %d", err, next)
	}
}
