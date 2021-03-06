package file

import (
	"compress/gzip"
	"errors"
	"os"

	"github.com/rk0cc/hson/structre"
)

/*
Write data to HSON file

It uses GZip compressed file context with level 5 compression
*/
func WriteData(path string, hj *structre.HashJSON) error {
	if hj.Context == "" || hj.Hashing == "" {
		return errors.New("hashjson has empty field and abort incoming operation")
	} else if nonJsonErr := hj.ContextIsJSON(); nonJsonErr != nil {
		return nonJsonErr
	}
	if ivExtErr := isHsonExt(path); ivExtErr != nil {
		return ivExtErr
	}
	cF, cfErr := os.Create(path)
	if cfErr != nil {
		return cfErr
	}
	hjb, hjbErr := hj.ToByte()
	if hjbErr != nil {
		return hjbErr
	}
	w, nwErr := gzip.NewWriterLevel(cF, 5)
	if nwErr != nil {
		return nwErr
	}
	_, wErr := w.Write(hjb)
	if wErr != nil {
		return wErr
	}
	return w.Close()
}
