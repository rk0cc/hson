package structre

import (
	"encoding/base64"
	"encoding/json"
	"reflect"

	"golang.org/x/crypto/sha3"
)

func generateHashing(ctx string) ([]byte, error) {
	s := sha3.New512()
	_, wsErr := s.Write([]byte(ctx))
	if wsErr != nil {
		return nil, wsErr
	}
	return s.Sum(nil), nil
}

func (hj *HashJSON) hashMatched() bool {
	provided, pErr := base64.StdEncoding.DecodeString(hj.Hashing)
	actual, aErr := generateHashing(hj.Context)
	if aErr != nil || pErr != nil {
		return false
	}
	return reflect.DeepEqual(provided, actual)
}

func (hj *HashJSON) ContextIsJSON() error {
	var dummyMap map[string]interface{}
	return json.Unmarshal([]byte(hj.Context), &dummyMap)
}
