package lib

import (
	"bytes"
	"io/ioutil"
	"os"
	"text/template"

	_ "gorp-with-template/lib/statik"

	"github.com/rakyll/statik/fs"
)

//go:generate statik -f -src=../sql -m

var files, _ = fs.New()

func GetSQL2(filename string, req interface{}) string {
	var buf bytes.Buffer
	f, _ := files.Open(string(os.PathSeparator) + filename)
	b, _ := ioutil.ReadAll(f)
	t := template.Must(template.New(filename).Parse(string(b)))
	t.Execute(&buf, req)
	return buf.String()
}
