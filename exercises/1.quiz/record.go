package main

import (
	"encoding/csv"
	"os"
)

type Record struct {
	question string
	answer   string
}

// ReadFile Read a file. Must be a CSV.
// TODO check file type and return error if not CSV.
func ReadFile(file string) ([][]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	cr := csv.NewReader(f)
	r, err := cr.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return r, nil
}
