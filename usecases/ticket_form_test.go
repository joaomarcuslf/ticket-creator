package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
}

func TestValidateForm01(t *testing.T) {
	validation := ValidateForm("", "")

	assert.Equal(t, validation.Ok, false)
	assert.Equal(t, validation.Title, REQUIRED_FIELD)
	assert.Equal(t, validation.Description, REQUIRED_FIELD)
}

func TestValidateForm02(t *testing.T) {
	validation := ValidateForm("                                                                                                                                                                                                             ", "")

	assert.Equal(t, validation.Ok, false)
	assert.Equal(t, validation.Title, FIELD_MAS_LENGTH)
	assert.Equal(t, validation.Description, REQUIRED_FIELD)
}

func TestValidateForm03(t *testing.T) {
	validation := ValidateForm("test", "test")

	assert.Equal(t, validation.Ok, true)
	assert.Equal(t, validation.Title, "")
	assert.Equal(t, validation.Description, "")
}
