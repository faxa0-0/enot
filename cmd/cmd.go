package cmd

import "flag"

var (
	command       string
	versionFlags  = flag.Bool("V", false, "Display the version of the CLI.")
	helpFlags     = flag.Bool("help", false, "Display this help message.")
)

func ExecuteCLI() error {
	flag.Parse()

	if *versionFlags {
		command = "version"
	}

	if *helpFlags {
		command = "help"
	}

	if len(flag.Args()) < 1 {
		command = ""
	} else {
		command = flag.Args()[0]
	}

	err := executeCommand(command)
	if err != nil {
		return err
	}

	return nil
}

func executeCommand(command string) error {
	switch command {
	case "":
		rootCmd()
		return nil
	case "version":
		versionCmd()
		return nil
	case "help":
		helpCmd()
		return nil
	default:
		unknownCmd(command)
		return nil
	}
}
