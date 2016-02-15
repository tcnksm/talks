package main

import (
	"os"
	"text/template"
)

const tmpl = `
{{ range . }}
{{ . }}
{{ end }}`

func main() {
	t := template.Must(template.New("").Parse(tmpl))
	t.Execute(os.Stdout, []string{"mars", "mercury", "pluto", "neptune"})
}
