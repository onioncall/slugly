package main

import (
	"os"

	"github.com/onioncall/slugly"
)

func main() {
    // config := slugly.SlugConfig {
    //     Speed: slugly.FastPrintSpeed,
    //     Writer: os.Stdout,
    // }

	config := slugly.NewConsoleSlugConfig(slugly.MediumPrintSpeed, os.Stdout)
    config.SlugPrint("Hello! This is a test string... for testing.")
}
