package checkpoint_postgres

import "github.com/langchain-ai/checkpoint-go"

// PostgresSaver is a placeholder checkpoint saver backed by Postgres.
type PostgresSaver struct{}

func (PostgresSaver) Save(data []byte, version int) (int, error) {
	// TODO: implement database persistence
	return version + 1, nil
}

var _ checkpoint.Saver = PostgresSaver{}
