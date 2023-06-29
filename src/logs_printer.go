package main

import (
	"container/heap"
	"github.com/delabania/armory-interview-task/src/log_files"
	"github.com/delabania/armory-interview-task/src/priority_queue"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

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
		if err != nil {
			log.Fatal("Couldn't parse timestamp: ", err)
		}
		if otherErr != nil {
			log.Fatal("Couldn't parse timestamp: ", otherErr)
		}
	}
	return ts < otherTs
}

func initPriorityQueue(pq *priority_queue.PriorityQueue, filesPool *log_files.LocalDiskFilePoolHandler) {

	// init queue
	for _, fileName := range filesPool.Files() {
		line, err := filesPool.NextLine(fileName)
		if err != nil {
			log.Fatal("Couldn't read line", err)
		}
		logLine := log_files.NewLogLineImpl(line)
		pq.Push(PriorityQueueItem{fileName: fileName, line: logLine})
	}
	heap.Init(pq)
}

func main() {
	dir := parseInputArgs(os.Args[1:])

	filesPool := log_files.NewLocalDiskFilePoolHandler(dir)
	defer filesPool.CloseFiles()

	pq := priority_queue.PriorityQueue{}
	initPriorityQueue(&pq, filesPool)

	for pq.Len() > 0 {
		minLineItem := pq.Pop().(PriorityQueueItem)

		println(minLineItem.line.Raw())

		line, _ := filesPool.NextLine(minLineItem.fileName)
		if line == "" {
			continue
		}

		logLine := log_files.NewLogLineImpl(line)

		pq.Push(PriorityQueueItem{fileName: minLineItem.fileName, line: logLine})
		heap.Fix(&pq, pq.Len()-1)
	}
}

func parseInputArgs(args []string) string {
	if len(args) != 1 {
		log.Fatal("Invalid number of arguments")
	}
	return filepath.Clean(args[0])
}
