package entity

import (
	"encoding/json"
	"io"
)

type Note struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (n *Note) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(n)
}
