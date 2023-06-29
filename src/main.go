package main

import (
	"container/heap"
	"github.com/delabania/armory-interview-task/src/log_files"
	"github.com/delabania/armory-interview-task/src/priority_queue"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func main() {
	dir := parseInputArgs(os.Args[1:])

	filesPool := log_files.NewLocalDiskFilePoolHandler(dir)
	defer filesPool.CloseFiles()

	pq := priority_queue.PriorityQueue{}
	initPriorityQueue(&pq, filesPool)

	for pq.Len() > 0 {
		minLineItem := pq.Pop().(PriorityQueueItem)

		println(minLineItem.line.Raw())

		line, _ := nextLineUntilValidOrEndOfFile(minLineItem.fileName, filesPool)
		if line == nil {
			// EOF
			continue
		}

		pq.Push(PriorityQueueItem{fileName: minLineItem.fileName, line: line})
		heap.Fix(&pq, pq.Len()-1)
	}
}

func parseInputArgs(args []string) string {
	if len(args) != 1 {
		log.Fatal("Invalid number of arguments")
	}
	return filepath.Clean(args[0])
}
