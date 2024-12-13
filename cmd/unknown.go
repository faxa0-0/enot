package cmd

import "fmt"

func unknownCmd(command string) {
	fmt.Println("Oops! 🦝 enot says: '" + command + "' is an unknown command.")
	fmt.Println("Hint: Run 'enot help' for usage. 🚀")
}
