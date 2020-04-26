package csv_conv

import (
	"encoding/csv"
	"errors"
	"io"
)

type Converter struct {
	original [][]string
}

func NewConverter(r io.Reader) (*Converter, error) {
	if r == nil {
		return nil, errors.New("empty source")
	}
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return &Converter{
		original: records,
	}, nil
}

func (c *Converter) ChangeColumnName(newNames map[string]string) ([][]string, error) {
	if c.original == nil || len(c.original) < 1 {
		return nil, errors.New("csv records are empty")
	}

	header := c.original[0]
	for i, org := range header {
		if newName, ok := newNames[org]; ok {
			header[i] = newName
		}
	}
	return c.original, nil
}
