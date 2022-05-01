package data

import (
	"encoding/json"
	"io"
)

// ToJSON serialize the given interface into a string based JSON format
func ToJSON(i interface{}, w io.writer) error {
	e := json.NewEncoder(w)

	return e.Encode(I)
}

// FromJSON deserialize the object from JSON string
// in an io.Reader to the given interface
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}
