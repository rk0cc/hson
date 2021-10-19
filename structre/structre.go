package structre

import (
	"encoding/base64"
	"encoding/json"
	"errors"
)

/*
Structre for storing HashJSON
*/
type HashJSON struct {
	// Base64 encoded hashing
	Hashing string `json:"hashing"`
	// Context string
	Context string `json:"context"`
}

func (hj *HashJSON) UpdateContext(ctx string) error {
	hj.Context = ctx
	hash, ghErr := generateHashing(ctx)
	if ghErr != nil {
		return ghErr
	}
	hj.Hashing = base64.StdEncoding.EncodeToString(hash)
	if !hj.hashMatched() {
		return errors.New("update hashing failed")
	}
	if nonJsonErr := hj.ContextIsJSON(); nonJsonErr != nil {
		return nonJsonErr
	}
	return nil
}

func ParseFromByte(data []byte) (HashJSON, error) {
	var phj HashJSON
	jsonUMErr := json.Unmarshal(data, &phj)
	if jsonUMErr != nil {
		return HashJSON{}, jsonUMErr
	} else if nonJsonErr := phj.ContextIsJSON(); nonJsonErr != nil {
		return HashJSON{}, nonJsonErr
	} else if !phj.hashMatched() {
		return HashJSON{}, errors.New("hashing mismatched")
	}
	return phj, nil
}

func (hj *HashJSON) ToByte() ([]byte, error) {
	return json.Marshal(hj)
}
