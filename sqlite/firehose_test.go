package sqlite_test

import (
	"database/sql"
	"io/ioutil"
	"testing"

	"github.com/stillwondering/fhc"

	"github.com/stillwondering/fhc/sqlite"

	_ "github.com/mattn/go-sqlite3"
)

func TestFindAll(t *testing.T) {
	testcases := []struct {
		name         string
		testdataFile string
		expected     fhc.Firehoses
	}{
		{
			name:         "Empty collection",
			testdataFile: "testdata/empty.sql",
			expected:     fhc.Firehoses{},
		},
		{
			name:         "One firehose in collection",
			testdataFile: "testdata/one.sql",
			expected: fhc.Firehoses{
				{
					ID:     1,
					Code:   1,
					Type:   "C",
					Length: 15,
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			db := setupDatabase(t)
			insertTestdata(t, db, tc.testdataFile)

			repo := sqlite.FirehoseRepository{
				DB: db,
			}

			res, err := repo.FindAll()
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}

			if len(*res) != len(tc.expected) {
				t.Errorf("expected %v, got %v", *res, tc.expected)
			}
		})
	}
}

func setupDatabase(t *testing.T) *sql.DB {
	path := t.TempDir() + "/test.db"
	t.Logf("opening DB file in %s", path)
	db, err := sql.Open("sqlite3", t.TempDir()+"/test.db")
	if err != nil {
		t.Fatalf("expected to open DB connection, got %v", err)
	}

	schema, err := ioutil.ReadFile("./../dist/sql/schema.sql")
	if err != nil {
		t.Fatalf("expected to read schema file, got %v", err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		t.Fatalf("expected to create DB schema, got %v", err)
	}

	return db
}

func insertTestdata(t *testing.T, db *sql.DB, file string) {
	testdata, err := ioutil.ReadFile(file)
	if err != nil {
		t.Fatalf("expected to read testdata file, got %v", err)
	}

	_, err = db.Exec(string(testdata))
	if err != nil {
		t.Fatalf("expected to insert testdata, got %v", err)
	}
}
