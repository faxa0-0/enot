package cmd

import "fmt"

func helpCmd() {
	fmt.Println("🦝 is network monitoring helper")
	fmt.Println()
	fmt.Println("Usage: enot <command> [options]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  help, --help        display this help message.")
	fmt.Println("  version, -V         display the version of the CLI.")
}
