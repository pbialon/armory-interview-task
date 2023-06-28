package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const LogFileExtension = "log"

func main() {
	dir := parseInputArgs(os.Args[1:])
	files, err := getLogFiles(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(files)
}

func getLogFiles(dir string) ([]string, error) {
	return filepath.Glob(filepath.Join(dir, fmt.Sprintf("*.%s", LogFileExtension)))
}

func parseInputArgs(args []string) string {
	if len(args) != 1 {
		fmt.Println("Invalid number of arguments")
		os.Exit(1)
	}
	return filepath.Clean(args[0])
}
