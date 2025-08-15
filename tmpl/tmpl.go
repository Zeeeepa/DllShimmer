package tmpl

import (
	"dllshimmer/dll"
	"log"
	"os"
	"text/template"
)

type TemplateParams struct {
	Functions []dll.ExportedFunction
	ProxyDll  string
	DllName   string
	Mutex     bool
}

func CreateShimCode(params TemplateParams) {
	tmpl := template.Must(template.ParseFiles("templates/shim.c.template"))

	err := tmpl.Execute(os.Stdout, params)
	if err != nil {
		log.Fatalf("[!] Error of template engine: %v", err)
	}
}
