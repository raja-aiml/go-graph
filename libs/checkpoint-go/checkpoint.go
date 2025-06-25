package checkpoint

// Saver defines the interface for persisting checkpoints.
type Saver interface {
	// Save persists the given data and returns a new version identifier.
	Save(data []byte, version int) (int, error)
}
