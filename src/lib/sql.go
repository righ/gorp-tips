package lib

import (
	"bytes"
	"text/template"
)

func GetSQL(filename string, req interface{}) string {
	var buf bytes.Buffer
	t := template.Must(template.ParseFiles("sql/" + filename))
	t.Execute(&buf, req)
	return buf.String()
}
