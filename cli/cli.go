package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type CliFlags struct {
	Input  string
	Output string
	Proxy  string
	Mutex  bool
	Static bool
}

func IsValidWindowsDllName(filename string) bool {
	invalidChars := []rune{'<', '>', ':', '"', '/', '\\', '|', '?', '*'}

	// Check for invalid characters
	for _, char := range invalidChars {
		if strings.ContainsRune(filename, char) {
			return false
		}
	}

	if !strings.HasSuffix(filename, ".dll") {
		return false
	}

	return true
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

	flag.BoolVar(&flags.Static, "static", false, "")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: DllShimmer -i <path> -o <path> -p <path>\n")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println()
		fmt.Printf("  %-24s %s\n", "-i, --input <path>", "Input DLL file (required)")
		fmt.Printf("  %-24s %s\n", "-o, --output <path>", "Output directory (required)")
		fmt.Printf("  %-24s %s\n", "-p, --proxy <path>", "Original DLL on target (required)")
		fmt.Printf("  %-24s %s\n", "-m, --mutex", "Multiple execution prevention (default: false)")
		fmt.Printf("  %-24s %s\n", "--static", "Original DLL loaded via static IAT (default: false)")
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

	if flags.Static && !IsValidWindowsDllName(flags.Proxy) {
		log.Fatalln("[!] In case of static linking enabled the proxy file (-p, --proxy) must be valid Windows DLL file name with no path information. E.g. kernel32.dll, user32.dll")
	}

	return &flags
}
