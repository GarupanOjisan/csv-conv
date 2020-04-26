package csv_conv

import "errors"

func (c *Converter) AppendColumn(colName, defaultValue string) ([][]string, error) {
	if c.original == nil || len(c.original) < 1 {
		return nil, errors.New("empty source")
	}
	c.original[0] = append(c.original[0], colName)
	for i := 1; i < len(c.original); i++ {
		c.original[i] = append(c.original[i], defaultValue)
	}
	return c.original, nil
}
