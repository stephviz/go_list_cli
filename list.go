package main

import (
	"fmt"
	"os"
	"flag"
)

func displayNumberedList(args []string) {
	for i := 0; i < len(args); i++ {
		fmt.Println(i + 1, args[i])
	}
}

func displayBulletedList(args []string) {
	for i := 0; i < len(args); i++ {
		fmt.Println("- ", args[i])
	}
}

func main() {
	args := os.Args[1:]
	flagPtr := flag.Bool("numbered", false, "ordered list flag")
	flag.Parse()

	if *flagPtr {
		items := args[1:len(args)]
		displayNumberedList(items)
	} else {
		displayBulletedList(args)
	}
}
