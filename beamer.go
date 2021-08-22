package beamer

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Init() {
	_, err := os.Stat(".beamer/")
	if os.IsNotExist(err) {
		fmt.Println(".beamer folder Not Found")
		fmt.Println("creating...")
		err := os.Mkdir(".beamer", os.ModePerm)
		if err != nil {
			fmt.Println("Error: Please create the `.beamer` directory manually")
			panic(err)
		}
		tempConfig := "templateDir=<PATH/TO/REPO>/DataflowTemplates/src/main/java/com/google/cloud/teleport/templates/"
		err = ioutil.WriteFile(".beamer/.config", []byte(tempConfig), 0644)
		if err != nil {
			panic(err)
		}
		fmt.Println("Done.")
		os.Exit(0)
	}

	fmt.Println("Found `.beamer` directory, skipping...")
}
