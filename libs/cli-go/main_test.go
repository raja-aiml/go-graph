package main

import "testing"

func TestMain(t *testing.T) {
	// simply ensure the command builds and runs
	gofunc := func() { main() }
	gofunc()
}
