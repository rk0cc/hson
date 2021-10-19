package file

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidExt(t *testing.T) {
	errorCount := 0
	testExt := []string{
		"hson",
		"HSON",
		"Hson",
		"foo",
		"bar",
		"hashJson",
		"hashjson",
		"hj",
		"hjson",
	}

	for _, extN := range testExt {
		extErr := isHsonExt("testhson." + extN)
		if extErr != nil {
			assert.EqualError(t, extErr, fmt.Sprintf("the extension must be either .hson or .hashjson insteaded of .%s", extN))
			errorCount++
		}
	}

	var expectErr int
	if runtime.GOOS == "windows" {
		expectErr = 4
	} else {
		expectErr = 7
	}

	assert.EqualValues(t, expectErr, errorCount)
}
