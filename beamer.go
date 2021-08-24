package beamer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var (
	errBeamerDirNotFound           = errors.New("could not find .beamer directory. \n\tRun `beamer init` to create it")
	errRepoAbsolutePathNotSet      = errors.New("please set absolute path in `.beamer/.config/` to GCP Template repo")
	errTemplateDirNotFoundInConfig = errors.New("please set the template directory in .beamer/config. \n\ttemplateDir=/absolute/path/to/gcp/template ")
)

func Init() {
	if beamerDirIsExist() {
		fmt.Println("Found `.beamer` directory, skipping...")
		os.Exit(0)
	}

	// create .beamer directory
	fmt.Println("Could not find `.beamer` directory, attempting to create..")
	err := os.Mkdir(".beamer", os.ModePerm)
	if err != nil {
		panic(err)
	}
	tempConfig := "templateDir=<PATH/TO/REPO>/DataflowTemplates/src/main/java/com/google/cloud/teleport/templates/"
	err = ioutil.WriteFile(".beamer/config", []byte(tempConfig), 0644)
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

func Gen(templateName string) {
	if !beamerDirIsExist() {
		fmt.Println(errBeamerDirNotFound)
		os.Exit(64)
	}
	data, err := ioutil.ReadFile(".beamer/config")
	if err != nil {
		panic(err)
	}
	txt := string(data)
	if strings.Contains(txt, "<PATH/TO/REPO>") {
		fmt.Println(errRepoAbsolutePathNotSet)
		os.Exit(64)
	}
	if !strings.Contains(txt, "templateDir") {
		fmt.Println(errTemplateDirNotFoundInConfig)
		os.Exit(64)
	}
	config := strings.Split(txt, "=")
	filePath := fmt.Sprintf("%s%s.java", config[1], templateName)
	options := ExtractOptionsFromFile(filePath)
	options.WriteToFile(fmt.Sprintf("%s.json", templateName))
	fmt.Printf("Job config template generated for `%s` migration.\n", templateName)
}

func Run(templateName string) {
	data, err := ioutil.ReadFile(fmt.Sprintf(".beamer/%s.json", templateName))
	if err != nil {
		panic(err)
	}
	var config JobConfig
	json.Unmarshal(data, &config)
	config.Validate()
	gcloudExecPath, err := exec.LookPath("gcloud")
	if err != nil {
		panic(err)
	}
	cmdGCloud := &exec.Cmd{
		Path: gcloudExecPath,
		Args: []string{
			gcloudExecPath, "dataflow", "jobs", "run", config.JobName,
			fmt.Sprintf("--gcs-location=%v", config.GcsLocation),
			fmt.Sprintf("--region=%v", config.Region),
			fmt.Sprintf("--service-account-email=%v", config.ServiceAccount),
			fmt.Sprintf("--project=%v", config.Project),
			fmt.Sprintf("--parameters %v", config.ParamString()),
		},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	fmt.Println(cmdGCloud.String())
}

func beamerDirIsExist() bool {
	_, err := os.Stat(".beamer/")
	return err == nil
}
