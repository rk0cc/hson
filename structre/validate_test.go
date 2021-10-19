package structre

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFalseWhenMismatchedHashing(t *testing.T) {
	thj1 := &HashJSON{
		Hashing: "aaaaaaaaaaaaaaaaa",
		Context: `{"foo":"bar}`,
	}
	assert.EqualValues(t, false, thj1.hashMatched())
}

func TestErrorWhenNonJSONApplied(t *testing.T) {
	thj2 := &HashJSON{}
	njErr := thj2.UpdateContext("I am not JSON")
	if assert.Error(t, njErr) {
		assert.EqualError(t, njErr, "invalid character 'I' looking for beginning of value")
	} else {
		t.Error("The context is 'I am not JSON' which is not JSON")
	}
}
