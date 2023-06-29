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

type PriorityQueueItem struct {
	fileName string
	line     log_files.LogLine
}

func (pq PriorityQueueItem) Lt(other interface{}) bool {
	otherItem := other.(PriorityQueueItem)
	ts, err := pq.line.Timestamp()
	otherTs, otherErr := otherItem.line.Timestamp()
	if err != nil || otherErr != nil {
		// todo: is it really fatal?
		log.Fatal("Couldn't parse timestamp")
	}
	return ts < otherTs
}

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
