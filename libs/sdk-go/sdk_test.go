package langgraph

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Go")
	want := "Hello, Go!"
	if got != want {
		t.Errorf("Hello('Go') = %q, want %q", got, want)
	}
}
