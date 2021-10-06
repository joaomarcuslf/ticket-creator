package encoders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
}

func TestEncode01(t *testing.T) {
	key := "joaomarcuslf"

	result := Encode(key)

	assert.True(t, result == "am9hb21hcmN1c2xm")
}

func TestEncode02(t *testing.T) {
	key := "joaomarcuslf1"

	result := Encode(key)

	assert.True(t, result == "am9hb21hcmN1c2xmMQ==")
}

func TestDecode01(t *testing.T) {
	key := "am9hb21hcmN1c2xm"

	result, _ := Decode(key)

	assert.True(t, result == "joaomarcuslf")
}

func TestDecode02(t *testing.T) {
	key := "am9hb21hcmN1c2xmMQ=="

	result, _ := Decode(key)

	assert.True(t, result == "joaomarcuslf1")
}
