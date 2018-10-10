package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEmptyString(t *testing.T) {
	_, err := ParseLMessage("")

	assert.Equal(t, ErrInvalidMessage, err)
}
