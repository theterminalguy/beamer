package main

import (
	"fmt"
	"os"

	"github.com/ls-simon-peter-damian/beamer"
)

func main() {
	if len(os.Args) < 2 {
		showHelp()
	}
	switch os.Args[1] {
	case "gen":
		validateTemplateName(os.Args[2])
		beamer.Gen(os.Args[2])
	case "run":
		validateTemplateName(os.Args[2])
		beamer.Run(os.Args[2])
	default:
		showHelp()
	}
}

func showHelp() {
	text := "usage: beamer <command> <args>\n\nThese are the only two commands available:\n\tgen\tgenerates a new job config\n\trun\truns the generated job config on GCP\n\nExamples:\n\t- Generates a job config for BigQueryToDatastore\n\t$ beamer gen BigQueryToDatastore\n\n\t- Run the job config for BigQueryToDatastore migration\n\t$ beamer run BigQueryToDatastore"
	fmt.Println(text)
	os.Exit(0)
}

func validateTemplateName(name string) {
	if name == "" {
		fmt.Println("Please provide the template name. See `beamer help` for more info.")
		os.Exit(64)
	}
}
