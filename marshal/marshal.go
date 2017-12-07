package marshal

import (
	"bytes"
	"encoding/json"
)

func PureMarshal(t interface{}) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(t)
	return buf.Bytes(), err
}

func PureMarshalIndent(t interface{}, prefix, indent string) ([]byte, error) {
	b, err := PureMarshal(t)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, b, prefix, indent)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
