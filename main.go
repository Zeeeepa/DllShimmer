package main

import (
	"dllshimmer/cli"
	"dllshimmer/dll"
	"dllshimmer/output"
	"path/filepath"
)

func main() {
	flags := cli.ParseCli()

	// var params tmpl.CodeFileParams
	// params.Functions = dll.ExportedFunctions
	// params.OriginalPath = flags.OriginalPath
	// params.DllName = filepath.Base(flags.Input)
	// params.Mutex = flags.Mutex

	out := output.Output{
		Dll:       dll.ParseDll(flags.Input),
		OutputDir: filepath.Clean(flags.Output),
	}

	out.CreateCodeFile(flags.Mutex, flags.Static)
	out.CreateDefFile()
	out.CreateCompileScript(flags.Static)

	if flags.Static {
		out.CreateLibFile()
	}
}
