package fhc

type ErrInvalidFirehoseType struct {
	InvalidType string
}

func (err ErrInvalidFirehoseType) Error() string {
	return err.InvalidType + " is not a valid firehose type"
}
