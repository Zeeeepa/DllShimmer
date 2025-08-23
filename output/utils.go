package output

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func sanitizePathForInjection(path string) string {
	return strings.ReplaceAll(path, "\\", "\\\\")
}

func createFileFromTemplate[K interface{}](o *Output, template string, outputPath string, params K) {
	tmpl := o.GetTemplate(template)

	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("[!] Error while creating '%s' file: %v", outputPath, err)
	}
	defer f.Close()

	err = tmpl.Execute(f, params)
	if err != nil {
		log.Fatalf("[!] Error of template engine: %v", err)
	}

	fmt.Printf("[+] '%s' file created\n", outputPath)
}
