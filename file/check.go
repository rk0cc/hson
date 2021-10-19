package file

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

/*
Check is a valid HSON file extension

It must be either .hson or .hashjson
*/
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
	return fmt.Errorf("the extension must be either .hson or .hashjson insteaded of %s", fext)
}
