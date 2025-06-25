package checkpoint_sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	checkpoint "github.com/langchain-ai/checkpoint-go"
)

// SQLiteSaver persists checkpoints in a SQLite database.
type SQLiteSaver struct {
	DB *sql.DB
}

// NewSQLiteSaver initializes the saver and ensures required tables exist.
func NewSQLiteSaver(db *sql.DB) (*SQLiteSaver, error) {
	s := &SQLiteSaver{DB: db}
	if err := s.setup(); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *SQLiteSaver) setup() error {
	_, err := s.DB.Exec(`CREATE TABLE IF NOT EXISTS checkpoints(
        version INTEGER PRIMARY KEY,
        data BLOB
    )`)
	return err
}

// Save stores the data blob and returns the next version.
func (s *SQLiteSaver) Save(data []byte, version int) (int, error) {
	next := version + 1
	if _, err := s.DB.Exec(`INSERT INTO checkpoints(version, data) VALUES(?, ?)`, next, data); err != nil {
		return version, err
	}
	return next, nil
}

var _ checkpoint.Saver = (*SQLiteSaver)(nil)
