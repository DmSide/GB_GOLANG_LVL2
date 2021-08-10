package main

import (
	"flag"
	"fmt"

	"../task8"
)

func main() {
	// argsWithoutProg := os.Args[1:]

	helpPtr := flag.Bool("-h", false, "Help function")
	helpPtr2 := flag.Bool("--help", false, "Help function")
	deletePrt  := flag.Bool("-d", false, "-d")
	deletePrt2  := flag.Bool("--delete", false, "--delete")

	flag.Parse()

	_ = helpPtr
	_ = helpPtr2

	if *helpPtr || *helpPtr2 {
		fmt.Println("Help function!")
		return
	}

	deleteAfterFind := false

	if *deletePrt || *deletePrt2{
		deleteAfterFind = true
	}

	task8.FindDuplicates("./files", deleteAfterFind)
}
