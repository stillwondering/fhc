package fhc

import (
	"encoding/json"
)

const (
	C string = "C"
	B string = "B"
)

// The Firehose object contains data to identify one specific firehose.
type Firehose struct {
	ID     int
	Code   string
	Type   string
	Length int
}

// Firehoses is a collection of Firehose objects.
type Firehoses []Firehose

// FirehoseRepository represents an object that manages a collection of
// Firehose objects.
type FirehoseRepository interface {
	FindAll() (*Firehoses, error)
	Find(id int) (*Firehose, error)
	FindByCode(string) (*Firehose, error)
	Create(CreateFirehoseData) (*Firehose, error)
}

type CreateFirehoseData struct {
	Code   string
	Type   string
	Length int
}

func (f Firehose) String() string {
	rep := struct {
		ID     int    `json:"id"`
		Code   string `json:"code"`
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
