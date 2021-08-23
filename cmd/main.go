package main

import (
	"fmt"
	"os"
)

func main() {
	//beamer.Gen("BigQueryToDatastore")
	if len(os.Args) < 2 {
		showHelp()
	}
	switch os.Args[1] {
	case "gen":
		fmt.Println("Running `gen` command", os.Args[2])
	case "run":
		fmt.Println("Running `run` command", os.Args[2])
	default:
		showHelp()
	}
}

func showHelp() {
	text := "usage: beamer <command> <args>\n\nThese are the only two commands available:\n\tgen\tgenerates a new job config\n\trun\truns the generated job config on GCP\n\nExamples:\n\t- Generates a job config for BigQueryToDatastore\n\t$ beamer gen BigQueryToDatastore\n\n\t- Run the job config for BigQueryToDatastore migration\n\t$ beamer run BigQueryToDatastore"
	fmt.Println(text)
	os.Exit(0)
}
