package fhc

import (
	"encoding/json"
	"time"
)

// The Firehose object contains data to identify one specific firehose.
type Firehose struct {
	ID        int
	Code      int
	Type      string
	Length    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Firehoses is a collection of Firehose objects.
type Firehoses []Firehose

// FirehoseRepository represents an object that manages a collection of
// Firehose objects.
type FirehoseRepository interface {
	FindAll() (*Firehoses, error)
}

func (f Firehose) String() string {
	rep := struct {
		ID     int    `json:"id"`
		Code   int    `json:"code"`
		Type   string `json:"type"`
		Length int    `json:"length"`
	}{
		ID:     f.ID,
		Code:   f.Code,
		Type:   f.Type,
		Length: f.Length,
	}

	b, _ := json.Marshal(rep)

	return string(b)
}
