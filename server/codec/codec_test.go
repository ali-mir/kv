package codec

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEncodeDecode(t *testing.T) {
	bytes, err := Encode("test"); if err != nil {
		panic("Failed to encode")
	}

	byteArray := bytes.Bytes()
	str, err := Decode(&byteArray); if err != nil {
		panic("Failed to decode")
	}
	assert.Equal(t, "test", *str, "decoded string is not 'test'")
}