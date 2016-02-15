package main

import (
	"os"
	"text/template"
)

func main() {
	master := `{{ block "T1" . }} T1 {{ end }}`
	masterTmpl, _ := template.New("master").Parse(master)

	overlay := `{{ define "T1" }} redefinition T1 {{ end }}`
	overlayTmpl, _ := template.Must(masterTmpl.Clone()).Parse(overlay)
	overlayTmpl.Execute(os.Stdout, nil)
}
