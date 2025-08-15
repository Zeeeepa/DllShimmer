package def

import (
	"log"
	"os"
)

type exportedFunction struct {
	OriginalName string
	Rename       string
	Forwarder    string
}

type DefFile struct {
	DllName           string
	Path              string
	exportedFunctions []exportedFunction
}

func (d *DefFile) AddExportedFunction(name string) {
	d.exportedFunctions = append(d.exportedFunctions, exportedFunction{
		OriginalName: name,
	})
}

func (d *DefFile) AddRenamedFunction(originalName string, rename string) {
	d.exportedFunctions = append(d.exportedFunctions, exportedFunction{
		OriginalName: originalName,
		Rename:       rename,
	})
}

func (d *DefFile) AddForwardedFunction(originalName string, forwarder string) {
	d.exportedFunctions = append(d.exportedFunctions, exportedFunction{
		OriginalName: originalName,
		Forwarder:    forwarder,
	})
}

func (d *DefFile) SaveFile() {
	var content string

	content += "LIBRARY \"" + d.DllName + "\"\n"
	content += "EXPORTS\n"

	for _, function := range d.exportedFunctions {
		if function.Forwarder != "" {
			// Forwarded functions
			content += "\t" + function.OriginalName + "=" + function.Forwarder + "\n"
			continue
		}

		if function.Rename != "" {
			// Exported-renamed functions
			content += "\t" + function.Rename + "=" + function.OriginalName + "\n"
			continue
		}

		content += "\t" + function.OriginalName + "\n"
	}

	content += "\n"

	err := os.WriteFile(d.Path, []byte(content), 0644)
	if err != nil {
		log.Fatalf("[!] Error while creating .def file: %v", err)
	}
}
