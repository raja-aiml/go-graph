package prebuilt

import "testing"

func TestHelloAgent(t *testing.T) {
	if HelloAgent() != "hello from prebuilt" {
		t.Fatalf("unexpected output")
	}
}
