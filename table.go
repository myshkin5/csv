package csv

import (
	"errors"
	"io"
)

type Table struct {
	values      [][]string
	columnIndex map[string]int

	current int
}

func New(columns []string, values [][]string) (*Table, error) {
	columnIndex := make(map[string]int)
	for n, column := range values[0] {
		columnIndex[column] = n
	}
	for _, column := range columns {
		_, ok := columnIndex[column]
		if !ok {
			return nil, errors.New("Column " + column + " not found in header record")
		}
	}

	return &Table{
		values:      values,
		columnIndex: columnIndex,
		current:     0,
	}, nil
}

func (t *Table) Value(column string) (string, bool) {
	columnIndex, ok := t.columnIndex[column]
	if !ok {
		return "", false
	}
	return t.values[t.current][columnIndex], true
}

func (t *Table) Next() error {
	t.current += 1
	if t.current >= len(t.values) {
		return io.EOF
	}
	return nil
}
