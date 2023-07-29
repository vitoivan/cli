package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/go-envparse"
)

func getFileContent(variables map[string]string) []byte {

	var fileContent []byte

	for k, v := range variables {
		v = strings.ReplaceAll(v, "\"", "\\\"")
		value := "\"" + v + "\""

		line := k + "=" + value + "\n"
		fileContent = append(fileContent, line...)
	}

	return fileContent
}

func main() {

	// get file content from path

	if len(os.Args) != 4 {
		fmt.Println("Cannot execute envparser without required args")
		fmt.Println("Usage: ./cmd <envpath> <key> <value>")
		os.Exit(1)
	}

	args := os.Args[1:]
	filepath := args[0]
	key := args[1]
	value := args[2]

	key = strings.ReplaceAll(key, " ", "_")

	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	variables, err := envparse.Parse(file)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	variables[key] = value

	fileContent := getFileContent(variables)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = os.WriteFile(filepath, fileContent, 0644)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
