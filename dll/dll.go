package dll

import (
	"log"
	"path/filepath"

	peparser "github.com/saferwall/pe"
)

type ExportedFunction struct {
	Name      string
	Forwarder string
}

type Dll struct {
	Name              string
	ExportedFunctions []ExportedFunction
}

func ParseDll(dllPath string) *Dll {
	var dll Dll

	dll.Name = filepath.Base(dllPath)

	pe, err := peparser.New(dllPath, &peparser.Options{})
	if err != nil {
		log.Fatalf("[!] Error while opening file: %s, reason: %v", dllPath, err)
	}

	err = pe.Parse()
	if err != nil {
		log.Fatalf("[!] Error while parsing file: %s, reason: %v", dllPath, err)
	}

	for _, function := range pe.Export.Functions {
		dll.ExportedFunctions = append(dll.ExportedFunctions, ExportedFunction{
			Name:      function.Name,
			Forwarder: function.Forwarder,
		})
	}

	return &dll
}
