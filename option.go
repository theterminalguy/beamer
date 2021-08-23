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

type JobOptions []string

func (options JobOptions) WriteToFile(fileName string) {
	config := map[string]interface{}{
		"jobName":        blank,
		"gcsLocation":    blank,
		"region":         blank,
		"project":        blank,
		"serviceAccount": blank,
	}
	if len(options) < 1 {
		return
	}
	parameters := map[string]string{}
	for _, option := range options {
		parameters[option] = blank
	}
	config["parameters"] = parameters

	// write config to file
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
