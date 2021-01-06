package version

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var showVersion = flag.Bool("version", false, "print version of this binary (only valid if compiled with make)")

var (
	version string
	date    string
)

// Init initializes the version flag to allow it to show version when the
// -version flag is passed to the binary.
func Init() {
	if !flag.Parsed() {
		flag.Parse()
	}
	if showVersion != nil && *showVersion {
		printVersion(os.Stdout, version, date)
		os.Exit(0)
	}
}

func printVersion(w io.Writer, version string, date string) {
	fmt.Fprintf(w, "Version: %s\n", version)
	fmt.Fprintf(w, "Binary: %s\n", os.Args[0])
	fmt.Fprintf(w, "Compile date: %s\n", date)
	fmt.Fprintf(w, "(version and date only valid if compiled with make)\n")
}
