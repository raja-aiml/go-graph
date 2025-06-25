package checkpoint_postgres

import (
	"database/sql"

	_ "github.com/lib/pq"

	checkpoint "github.com/langchain-ai/checkpoint-go"
)

// PostgresSaver persists checkpoints in a Postgres database.
type PostgresSaver struct {
	DB *sql.DB
}

// NewPostgresSaver initializes the saver and ensures required tables exist.
func NewPostgresSaver(db *sql.DB) (*PostgresSaver, error) {
	p := &PostgresSaver{DB: db}
	if err := p.setup(); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *PostgresSaver) setup() error {
	_, err := p.DB.Exec(`CREATE TABLE IF NOT EXISTS checkpoints(
        version INTEGER PRIMARY KEY,
        data BYTEA
    )`)
	return err
}

// Save stores the data blob and returns the next version.
func (p *PostgresSaver) Save(data []byte, version int) (int, error) {
	next := version + 1
	if _, err := p.DB.Exec(`INSERT INTO checkpoints(version, data) VALUES($1, $2)`, next, data); err != nil {
		return version, err
	}
	return next, nil
}

var _ checkpoint.Saver = (*PostgresSaver)(nil)
