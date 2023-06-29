package main

import (
	"github.com/delabania/armory-interview-task/src/log_files"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

/*
type filePool interface {
	NextLine(fileName string) (string, error)
}*/

func main() {
	dir := parseInputArgs(os.Args[1:])
	filesPool := log_files.NewLocalDiskFilePoolHandler(dir)
	defer filesPool.CloseFiles()
}

func parseInputArgs(args []string) string {
	if len(args) != 1 {
		log.Fatal("Invalid number of arguments")
	}
	return filepath.Clean(args[0])
}
