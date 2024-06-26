package main

import (
	"flag"
	"fmt"
	"os"
)

var usage string = `
gid is a command line utility for working with UUIDs

usage	: gid [options]
options	:`

func main() {

	// Usage
	flag.Usage = func() {
		fmt.Fprintln(os.Stdout, usage)
		flag.PrintDefaults()
	}

	// Flags
	var gen int
	flag.IntVar(&gen, "gen", 1, "number of random UUIDs (Version 4) to be generated")
	flag.Parse()

	generateUuids(gen, true)
}

func handleAndExit(msg ...any) {
	fmt.Fprintln(os.Stderr, "error:", msg)
	os.Exit(1)
}

func checkAndExit(doExit bool) {
	if doExit {
		os.Exit(0)
	}
}
