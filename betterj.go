package betterj

import (
	"bytes"
	"encoding/json"
)

func MinifyJ(j string) (string, error) {

	//I used the bytes.buffer because its dynamic and works good with compact and indent
	var buffer bytes.Buffer

	//json.Compact will remove the whitespace and the new lines
	err := json.Compact(&buffer, []byte(j))
	if err != nil {
		return "", err
	}
	return buffer.String(), nil //converted the buffer content back to string
}

func BeautifyJ(j string, indentWith string) (string, error) {
	var buffer bytes.Buffer
	//json.Indent will format the JSON with custom indentation
	if err := json.Indent(&buffer, []byte(j), "", indentWith); err != nil {
		return "", err
	}
	return buffer.String(), nil
}
