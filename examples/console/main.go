package main

import (
	"os"

	"github.com/onioncall/slugly"
)

func main() {
	config := slugly.NewConsoleSlugConfig(slugly.FastPrintSpeed, os.Stdout)
    config.SlugPrint("Hello! This is a test string... for testing.")
}
