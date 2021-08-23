package beamer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"unicode"
)

const (
	optionSearchPattern = "options.get"
	blank               = ""
)

type JobConfig struct {
	JobName        string
	GcsLocation    string
	Region         string
	Project        string
	ServiceAccount string
	Parameters     map[string]string
}

func (jc *JobConfig) Validate() {
	errors := []string{}
	if jc.JobName == blank {
		errors = append(errors, "Error: Job Name is required")
	}
	if jc.GcsLocation == blank {
		errors = append(errors, "Error: GcsLocation is required")
	}
	if jc.Region == blank {
		errors = append(errors, "Error: Region is required")
	}
	if jc.Project == blank {
		errors = append(errors, "Error: Project is required")
	}
	if jc.ServiceAccount == blank {
		errors = append(errors, "Error: ServiceAccount is required")
	}
	for k, v := range jc.Parameters {
		if v == blank {
			errors = append(errors, fmt.Sprintf("Parameter Error: %v is required", k))
		}
	}
	if len(errors) > 0 {
		for _, err := range errors {
			fmt.Println(err)
		}
		os.Exit(64)
	}
}

type JobOptions []string

func (options JobOptions) WriteToFile(fileName string) {
	config := map[string]interface{}{
		"JobName":        blank,
		"GcsLocation":    blank,
		"Region":         blank,
		"Project":        blank,
		"ServiceAccount": blank,
	}
	if len(options) < 1 {
		return
	}
	parameters := map[string]string{}
	for _, option := range options {
		parameters[option] = blank
	}
	config["Parameters"] = parameters
	b, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fmt.Sprintf(".beamer/%s", fileName), b, 0644)
	if err != nil {
		panic(err)
	}
}

func ExtractOptionsFromFile(filePath string) JobOptions {
	var options JobOptions
	reader := open(filePath)
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		line = line[:len(line)-1]
		text := string(line)
		if strings.Contains(text, optionSearchPattern) {
			s := strings.Split(text, optionSearchPattern)
			option := re.ReplaceAllString(s[1], blank)
			options = append(options, lcFirst(option))
		}
	}
	return options
}

func open(path string) *bufio.Reader {
	file, err := os.Open(path)
	if err != nil {
		panic("cannot open file")
	}
	return bufio.NewReader(file)
}

func lcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return blank
}
