package file

import (
	"compress/gzip"
	"io/ioutil"
	"os"

	"github.com/rk0cc/hson/structre"
)

/*
Read the hash JSON context

It returns empty HashJSON if error is not nil
*/
func ReadData(path string) (structre.HashJSON, error) {
	if _, ipErr := os.Stat(path); ipErr != nil {
		return structre.HashJSON{}, ipErr
	} else if ivExtErr := isHsonExt(path); ivExtErr != nil {
		return structre.HashJSON{}, ivExtErr
	}
	f, ofErr := os.Open(path)
	if ofErr != nil {
		return structre.HashJSON{}, ofErr
	}
	defer f.Close()
	r, rErr := gzip.NewReader(f)
	if rErr != nil {
		return structre.HashJSON{}, rErr
	}
	defer r.Close()
	sr, srErr := ioutil.ReadAll(r)
	if srErr != nil {
		return structre.HashJSON{}, srErr
	}
	return structre.ParseFromByte(sr)
}
