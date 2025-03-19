package helper

import (
	"bytes"
	"html/template"
)

func Template(value string, data map[string]interface{}) (string, error) {

	t, err := template.New("").Parse(value)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	err = t.Execute(buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
