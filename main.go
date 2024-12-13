package main

import (
	"fmt"
	"os"

	"github.com/faxa0-0/enot/cmd"
)

func main() {
	err := cmd.ExecuteCLI()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
