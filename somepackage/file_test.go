package somepackage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddInts(t *testing.T) {
	result := AddInts(3, 4)
	assert.Equal(t, 7, result, "they should be equal")
}