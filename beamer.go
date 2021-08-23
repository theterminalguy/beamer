package beamer

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Init() {
	if foundBeamerDir() {
		fmt.Println("Found `.beamer` directory, skipping...")
		os.Exit(0)
	}

	fmt.Println("Could not find `.beamer` directory, attempting to create..")
	err := os.Mkdir(".beamer", os.ModePerm)
	if err != nil {
		panic(err)
	}
	tempConfig := "templateDir=<PATH/TO/REPO>/DataflowTemplates/src/main/java/com/google/cloud/teleport/templates/"
	err = ioutil.WriteFile(".beamer/.config", []byte(tempConfig), 0644)
	if err != nil {
		panic(err)
	}

	// ignore .beamer directory
	f, err := os.OpenFile(".gitignore", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err := f.WriteString("\n.beamer/"); err != nil {
		panic(err)
	}

	fmt.Println("Created, done!")
}

func Gen() {
	// TODO: this calls options write to file
	// Fails if `.beamer` directory does not exist
}

func Run() {
	// Executes the job, fails if no option is set
}

func foundBeamerDir() bool {
	_, err := os.Stat(".beamer/")
	return !os.IsNotExist(err)
}
