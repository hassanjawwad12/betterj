package betterj

import (
	"bytes"
	"encoding/json"
)

// I used the bytes.buffer because its dynamic and works good with compact and indent

func MinifyJ(j string) (string, error) {
	var buffer bytes.Buffer

	// json string converted to bytes for the compact function
	bytesData := []byte(j)

	// json.Compact will remove the whitespace and the new lines
	err := json.Compact(&buffer, bytesData)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil //converted the buffer content back to string
}

func BeautifyJ(j string, indentWith string) (string, error) {
	var buffer bytes.Buffer

	bytesData := []byte(j)

	// json.Indent will format the JSON with custom indentation
	err := json.Indent(&buffer, bytesData, "", indentWith)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
