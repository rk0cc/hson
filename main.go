package main

import "C"
import (
	"github.com/rk0cc/hson/file"
	"github.com/rk0cc/hson/structre"
)

func main() {
	// Placeholder for CGO
}

//export readHSON
func readHSON(path *C.char) *C.char {
	hsonCtx, rErr := file.ReadData(C.GoString(path))
	if rErr != nil {
		return nil
	}
	return C.CString(hsonCtx.Context)
}

//export writeHSON
func writeHSON(context, path *C.char) bool {
	hj := &structre.HashJSON{}
	ucErr := hj.UpdateContext(C.GoString(context))
	if ucErr != nil {
		return false
	}
	if wdErr := file.WriteData(C.GoString(path), hj); wdErr != nil {
		return false
	}
	return true
}
