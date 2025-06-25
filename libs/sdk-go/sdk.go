package langgraph

// Hello returns a greeting string.
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name + "!"
}
