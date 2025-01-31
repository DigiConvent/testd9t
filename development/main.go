package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	actionsFlagSet := flag.NewFlagSet("Options", flag.ExitOnError)

	check := actionsFlagSet.Bool("check", false, "Checks if everything is in order for a new release")
	helpFlag := actionsFlagSet.Bool("help", false, "Show help")
	actionsFlagSet.Parse(os.Args[1:])

	if *helpFlag || actionsFlagSet.NFlag() == 0 {
		actionsFlagSet.PrintDefaults()
		return
	}

	if *check {
		fmt.Println("Check")
	}
}
