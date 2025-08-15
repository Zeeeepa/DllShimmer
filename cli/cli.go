package cli

import (
	"flag"
	"fmt"
	"os"
)

type CliFlags struct {
	Input  string
	Proxy  string
	Mutex  bool
	Output string
}

func ParseCli() *CliFlags {
	var flags CliFlags

	flag.StringVar(&flags.Input, "i", "", "")
	flag.StringVar(&flags.Input, "input", "", "")

	flag.StringVar(&flags.Output, "o", "", "")
	flag.StringVar(&flags.Output, "output", "", "")

	flag.StringVar(&flags.Proxy, "p", "", "")
	flag.StringVar(&flags.Proxy, "proxy", "", "")

	flag.BoolVar(&flags.Mutex, "m", false, "")
	flag.BoolVar(&flags.Mutex, "mutex", false, "")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: DllShimmer -i <path> -o <path> -p <path>\n")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println()
		fmt.Printf("  %-24s %s\n", "-i, --input <path>", "Input DLL file (required)")
		fmt.Printf("  %-24s %s\n", "-o, --output <path>", "Output directory (required)")
		fmt.Printf("  %-24s %s\n", "-p, --proxy <path>", "Path to original DLL on target (required)")
		fmt.Printf("  %-24s %s\n", "-m, --mutex", "Multiple execution prevention (default: false)")
		fmt.Printf("  %-24s %s\n", "-h, --help", "Show this help")
		fmt.Println()
		fmt.Println("Example:")
		fmt.Println()
		fmt.Println("  DllShimmer -i version.dll -o ./project -p 'C:\\Windows\\System32\\version.dll' -m")
		fmt.Println()
		fmt.Println("Created by Print3M (print3m.github.io)")
		fmt.Println()
	}

	flag.Parse()

	if flags.Input == "" || flags.Output == "" || flags.Proxy == "" {
		flag.Usage()
		os.Exit(1)
	}

	return &flags
}
