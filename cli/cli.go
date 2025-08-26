package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const VERSION = "1.1.1"

type CliFlags struct {
	Input       string
	Output      string
	Original    string
	Mutex       bool
	Static      bool
	DebugFile   string
	ShowVersion bool
}

func IsValidDllName(filename string) bool {
	invalidChars := []rune{'<', '>', ':', '"', '/', '\\', '|', '?', '*'}

	// Check for invalid characters
	for _, char := range invalidChars {
		if strings.ContainsRune(filename, char) {
			return false
		}
	}

	if !strings.HasSuffix(strings.ToLower(filename), ".dll") {
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

	flag.StringVar(&flags.Original, "x", "", "")
	flag.StringVar(&flags.Original, "original", "", "")

	flag.StringVar(&flags.DebugFile, "debug-file", "", "")

	flag.BoolVar(&flags.Mutex, "m", false, "")
	flag.BoolVar(&flags.Mutex, "mutex", false, "")

	flag.BoolVar(&flags.Static, "static", false, "")

	flag.BoolVar(&flags.ShowVersion, "v", false, "")
	flag.BoolVar(&flags.ShowVersion, "version", false, "")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: DllShimmer -i <path> -o <path> -p <path>\n")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println()
		fmt.Printf("  %-26s %s\n", "-i, --input <path>", "Input DLL file (required)")
		fmt.Printf("  %-26s %s\n", "-o, --output <path>", "Output directory (required)")
		fmt.Printf("  %-26s %s\n", "-x, --original <path>", "Path to original DLL on target (required)")
		fmt.Printf("  %-26s %s\n", "-m, --mutex", "Multiple execution prevention (default: false)")
		fmt.Printf("  %-26s %s\n", "    --static", "Static linking to original DLL via IAT (default: false)")
		fmt.Printf("  %-26s %s\n", "    --debug-file <path>", "Save debug logs to a file (default: stdout)")
		fmt.Printf("  %-26s %s\n", "-v, --version", "Show version of DllShimmer")
		fmt.Printf("  %-26s %s\n", "-h, --help", "Show this help")
		fmt.Println()
		fmt.Println("Example:")
		fmt.Println()
		fmt.Println("  DllShimmer -i version.dll -o ./project -x 'C:\\Windows\\System32\\version.dll' -m")
		fmt.Println()
		fmt.Println("Created by Print3M (print3m.github.io)")
		fmt.Println()
	}

	flag.Parse()

	if flags.ShowVersion {
		fmt.Printf("DllShimmer %s\n", VERSION)
		os.Exit(0)
	}

	if flags.Input == "" || flags.Output == "" || flags.Original == "" {
		flag.Usage()
		os.Exit(1)
	}

	if flags.Static && !IsValidDllName(flags.Original) {
		fmt.Fprintf(os.Stderr, "[!] Invalid '-x' parameter value:\n")
		fmt.Fprintf(os.Stderr, "In case of static linking enabled '-x' parameter must be valid Windows DLL file name with no path information. E.g. kernel32.dll, user32.dll.")
		os.Exit(1)
	}

	return &flags
}

func PrintBanner() {
	banner := fmt.Sprintf(`
▓█████▄  ██▓     ██▓                           
▒██▀ ██▌▓██▒    ▓██▒                 By @Print3M
░██   █▌▒██░    ▒██░            (print3m.github.io)
░▓█▄   ▌▒██░    ▒██░                           
░▒████▓ ░██████▒░██████▒            Documentation:
 ▒▒▓  ▒ ░ ▒░▓  ░░ ▒░▓  ░     github.com/Print3M/DllShimmer
 ░ ▒  ▒ ░ ░ ▒  ░░ ░ ▒  ░                                   
 ░ ░  ░   ░ ░     ░ ░                   %s
   ░        ░  ░    ░  ░                                   
 ░                                                         
  ██████  ██░ ██  ██▓ ███▄ ▄███▓ ███▄ ▄███▓▓█████  ██▀███  
▒██    ▒ ▓██░ ██▒▓██▒▓██▒▀█▀ ██▒▓██▒▀█▀ ██▒▓█   ▀ ▓██ ▒ ██▒
░ ▓██▄   ▒██▀▀██░▒██▒▓██    ▓██░▓██    ▓██░▒███   ▓██ ░▄█ ▒
  ▒   ██▒░▓█ ░██ ░██░▒██    ▒██ ▒██    ▒██ ▒▓█  ▄ ▒██▀▀█▄  
▒██████▒▒░▓█▒░██▓░██░▒██▒   ░██▒▒██▒   ░██▒░▒████▒░██▓ ▒██▒
▒ ▒▓▒ ▒ ░ ▒ ░░▒░▒░▓  ░ ▒░   ░  ░░ ▒░   ░  ░░░ ▒░ ░░ ▒▓ ░▒▓░
░ ░▒  ░ ░ ▒ ░▒░ ░ ▒ ░░  ░      ░░  ░      ░ ░ ░  ░  ░▒ ░ ▒░
░  ░  ░   ░  ░░ ░ ▒ ░░      ░   ░      ░      ░     ░░   ░ 
      ░   ░  ░  ░ ░         ░          ░      ░  ░   ░     
`, VERSION)

	fmt.Print(banner)
	fmt.Println()
}
