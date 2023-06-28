package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

const LogFileExtension = "log"

func main() {
	dir := parseInputArgs(os.Args[1:])
	files, err := getLogFiles(dir)
	log.Infof("Found log %v files", len(files))
	if err != nil {
		log.Fatal("Couldn't get log files", err)
	}

	sortLogLines(files)
}

func sortLogLines(files []string) {
	fmt.Println(files)
}

func getLogFiles(dir string) ([]string, error) {
	return filepath.Glob(filepath.Join(dir, fmt.Sprintf("*.%s", LogFileExtension)))
}

func parseInputArgs(args []string) string {
	if len(args) != 1 {
		log.Fatal("Invalid number of arguments")
	}
	return filepath.Clean(args[0])
}
