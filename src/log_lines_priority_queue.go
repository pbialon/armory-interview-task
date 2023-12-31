package main

import (
	"container/heap"
	"github.com/delabania/armory-interview-task/src/log_files"
	"github.com/delabania/armory-interview-task/src/priority_queue"
	log "github.com/sirupsen/logrus"
)

type LogLine interface {
	Timestamp() (int64, error)
	Raw() string
	IsValid() bool
}

type PriorityQueueItem struct {
	fileName string
	line     LogLine
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
	for _, fileName := range filesPool.Files() {
		line := nextLineUntilIsValidOrEOF(fileName, filesPool)
		if line == nil {
			// EOF
			continue
		}
		item := PriorityQueueItem{fileName: fileName, line: line}
		heap.Push(pq, item)
	}
	heap.Init(pq)
}

func nextLineUntilIsValidOrEOF(fileName string, filesPool *log_files.LocalDiskFilePoolHandler) LogLine {
	for {
		line, err := filesPool.NextLine(fileName)
		if line == "" {
			if err == nil {
				// EOF
				return nil
			} else {
				// couldn't read line - try to read next one
				continue
			}
		}
		logLine := log_files.NewLogLineImpl(line)
		if !logLine.IsValid() {
			// try to read next line
			continue
		}
		return logLine
	}
}

func addNewLineAndFixHeap(pq *priority_queue.PriorityQueue, line LogLine, fileName string) {
	item := PriorityQueueItem{fileName: fileName, line: line}
	pq.Push(item)
	heap.Fix(pq, pq.Len()-1)
}
