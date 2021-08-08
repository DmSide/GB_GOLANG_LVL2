package main

import (
	"flag"

	"../task8"
)

func main() {
	// argsWithoutProg := os.Args[1:]

	_ = flag.String("-h", "Help function", "Help function")
	_ = flag.String("--help", "Help function", "Help function")

	flag.Parse()

	task8.FindDuplicates("./files")
}
