package main

import (
	"dllshimmer/cli"
	"dllshimmer/def"
	"dllshimmer/dll"
	"dllshimmer/tmpl"
	"path/filepath"
)

func main() {
	flags := cli.ParseCli()

	outputDir := filepath.Clean(flags.Output)

	dll := dll.ParseDll(flags.Input)

	var params tmpl.TemplateParams
	params.Functions = dll.ExportedFunctions
	params.ProxyDll = flags.Proxy
	params.DllName = filepath.Base(flags.Input)
	params.Mutex = flags.Mutex

	tmpl.CreateCodeFile(outputDir, params)

	// 1. Create temp .def based on original file
	dll.CreateLibFile(filepath.Join(outputDir, "original.lib"))

	// 2. Create new .def based on generated code
	func() {
		var def def.DefFile
		def.DllName = params.DllName

		for _, function := range dll.ExportedFunctions {
			if function.Forwarder == "" {
				def.AddRenamedFunction(function.Name, function.Name+"Fwd")
			} else {
				def.AddForwardedFunction(function.Name, function.Forwarder)
			}
		}

		def.SaveFile(filepath.Join(outputDir, params.DllName+".def"))
	}()
}
