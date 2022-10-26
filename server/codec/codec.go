package codec

import (
	"bytes"
	"encoding/gob"
)

func Encode(input string) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(input); err != nil {
		return nil, err
	}
	return &buf, nil

}

func Decode(input *[]byte) (*string, error){
	buf := bytes.NewBuffer(*input)
	dec := gob.NewDecoder(buf)

	s := ""
	if err := dec.Decode(&s); err != nil {
		return nil, err
	}

	return &s, nil
}