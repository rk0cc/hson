package file

import (
	"errors"
	"path/filepath"
	"runtime"
	"strings"
)

func isHsonExt(path string) error {
	validExt := []string{"hson", "hashjson"}
	fext := filepath.Ext(path)
	for _, vextName := range validExt {
		if runtime.GOOS == "windows" {
			fext = strings.ToLower(fext)
		}
		if "."+vextName == fext {
			return nil
		}
	}
	return errors.New("the extension must be either .hson or .hashjson insteaded of " + fext)
}
