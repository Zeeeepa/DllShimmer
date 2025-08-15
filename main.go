package main

import (
	"dllshimmer/cli"
	"dllshimmer/def"
	"dllshimmer/dll"
	"dllshimmer/tmpl"
	"os/exec"
	"path/filepath"
)

func main() {
	flags := cli.ParseCli()

	dll := dll.ParseDll(flags.Input)

	var params tmpl.TemplateParams
	params.Functions = dll.ExportedFunctions
	params.ProxyDll = flags.Proxy
	params.DllName = filepath.Base(flags.Input)
	params.Mutex = flags.Mutex

	tmpl.CreateShimCode(params)

	// 1. Create temp .def based on original file
	func() {
		var def def.DefFile
		def.DllName = params.DllName
		def.Path = "tmp.def"

		for _, function := range dll.ExportedFunctions {
			if function.Forwarder == "" {
				def.AddExportedFunction(function.Name)
			} else {
				def.AddForwardedFunction(function.Name, function.Forwarder)
			}
		}

		def.SaveFile()
		cmd := exec.Command("x86_64-w64-mingw32-dlltool", "-d", "tmp.def", "-l", "tmp.lib")
	}()

	// 2. Create new .def based on generated code
	func() {
		var def def.DefFile
		def.DllName = params.DllName
		def.Path = "exported.def"

		for _, function := range dll.ExportedFunctions {
			if function.Forwarder == "" {
				def.AddRenamedFunction(function.Name, function.Name+"Fwd")
			} else {
				def.AddForwardedFunction(function.Name, function.Forwarder)
			}
		}

		def.SaveFile()

	}()

}

// DllShimmer -i version.dll -o "C:/Windows/System32/version.dll" -p ~/project -m
