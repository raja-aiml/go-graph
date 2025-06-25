package checkpoint_sqlite

import "github.com/langchain-ai/checkpoint-go"

// SQLiteSaver is a placeholder checkpoint saver backed by SQLite.
type SQLiteSaver struct{}

func (SQLiteSaver) Save(data []byte, version int) (int, error) {
	// TODO: implement database persistence
	return version + 1, nil
}

var _ checkpoint.Saver = SQLiteSaver{}
