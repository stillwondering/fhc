package sqlite

import (
	"database/sql"

	"github.com/pkg/errors"
	"github.com/stillwondering/fhc"
)

// Interface guard to make sure that we really implement the given interface.
var _ fhc.FirehoseRepository = (*FirehoseRepository)(nil)

// FirehoseRepository manages a collection of Firehose objects using an
// underlying sqlite3 database.
type FirehoseRepository struct {
	DB *sql.DB
}

// FindAll returns all objects that are currently stored in the database.
func (r *FirehoseRepository) FindAll() (*fhc.Firehoses, error) {
	rows, err := r.DB.Query(`
		SELECT
			id,
			code,
			type,
			length
		FROM
			hose
	`)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	var collection fhc.Firehoses

	defer rows.Close()
	for rows.Next() {
		var id, length int
		var code, typ string

		err = rows.Scan(&id, &code, &typ, &length)
		if err != nil {
			return nil, errors.Wrap(err, "scan query results")
		}

		hose := fhc.Firehose{
			ID:     id,
			Code:   code,
			Type:   typ,
			Length: length,
		}

		collection = append(collection, hose)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "after iterating over rows")
	}

	return &collection, nil
}

func (r *FirehoseRepository) Find(id int) (*fhc.Firehose, error) {
	row := r.DB.QueryRow(`
		SELECT
			id,
			code,
			type,
			length
		FROM
			hose
		WHERE
			id = ?
	`, id)

	var rowid, length int
	var code, typ string

	err := row.Scan(&rowid, &code, &typ, &length)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "execute query")
	}

	return &fhc.Firehose{
		ID:     rowid,
		Code:   code,
		Type:   typ,
		Length: length,
	}, nil
}

func (r *FirehoseRepository) FindByCode(code string) (*fhc.Firehose, error) {
	row := r.DB.QueryRow(`
		SELECT
			id,
			code,
			type,
			length
		FROM
			hose
		WHERE
			code = ?
	`, code)

	var rowid, length int
	var c, typ string

	err := row.Scan(&rowid, &c, &typ, &length)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "execute query")
	}

	return &fhc.Firehose{
		ID:     rowid,
		Code:   c,
		Type:   typ,
		Length: length,
	}, nil
}

func (r *FirehoseRepository) Create(data fhc.CreateFirehoseData) (*fhc.Firehose, error) {
	if data.Type != fhc.B && data.Type != fhc.C {
		return nil, fhc.ErrInvalidFirehoseType{
			InvalidType: data.Type,
		}
	}

	stmt, err := r.DB.Prepare(`
		INSERT INTO hose (
			code,
			type,
			length,
			created_at,
			updated_at
		) VALUES (
			?,
			?,
			?,
			NOW(),
			NOW()
		)
	`)
	if err != nil {
		return nil, errors.Wrap(err, "prepare query")
	}

	res, err := stmt.Exec(data.Code, data.Type, data.Length)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	newID, err := res.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, "get last inserted ID")
	}

	return r.Find(int(newID))
}
