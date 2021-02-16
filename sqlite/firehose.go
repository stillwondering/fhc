package sqlite

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"github.com/stillwondering/fhc"
)

// FirehoseRepository manages a collection of Firehose objects using an
// underlying sqlite3 database.
type FirehoseRepository struct {
	DB *sql.DB
}

// FindAll returns all objects that are currently stored in the database.
func (r *FirehoseRepository) FindAll() (*fhc.Firehoses, error) {
	query := `
		SELECT
			*
		FROM
			hose
	`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	var collection fhc.Firehoses

	for rows.Next() {
		var id, code, length int
		var typ, createdAtText, updatedAtText string

		err = rows.Scan(&id, &code, &typ, &length, &createdAtText, &updatedAtText)
		if err != nil {
			return nil, errors.Wrap(err, "scan query results")
		}

		hose := fhc.Firehose{
			ID:     id,
			Code:   code,
			Type:   typ,
			Length: length,
		}

		createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtText)
		if err != nil {
			return nil, errors.Wrap(err, "parse timestamp")
		}

		updatedAt, err := time.Parse("2006-01-02 15:04:05", updatedAtText)
		if err != nil {
			return nil, errors.Wrap(err, "parse timestamp")
		}

		hose.CreatedAt = createdAt
		hose.UpdatedAt = updatedAt

		collection = append(collection, hose)
	}

	return &collection, nil
}
