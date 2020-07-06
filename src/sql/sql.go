package sql

import (
	"bytes"
	"text/template"
)

func GetSQL(filename string, req interface{}) string {
	var buf bytes.Buffer
	t := template.Must(template.ParseFiles("files/" + filename))
	t.Execute(&buf, req)
	return buf.String()
}
