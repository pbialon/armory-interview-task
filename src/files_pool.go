package main

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
)

const LogFileExtension = "log"

type LocalDiskFilePoolHandler struct {
	files    map[string]*os.File
	scanners map[string]*bufio.Scanner
}

func NewLocalDiskFilePoolHandler(dir string) *LocalDiskFilePoolHandler {
	files := make(map[string]*os.File)
	scanners := make(map[string]*bufio.Scanner)

	globPattern := filepath.Join(dir, fmt.Sprintf("*.%s", LogFileExtension))
	filePaths, err := filepath.Glob(globPattern)
	if err != nil {
		log.Fatal("Couldn't find log files", err)
	}
	log.Infof("Found log %v files", len(filePaths))

	for _, filePath := range filePaths {
		fileHandle, err := os.Open(filePath)
		if err != nil {
			log.Fatal("Couldn't open file", err)
		}
		fileName := path.Base(filePath)
		files[fileName] = fileHandle
		scanners[fileName] = bufio.NewScanner(fileHandle)
	}
	return &LocalDiskFilePoolHandler{files: files, scanners: scanners}
}

func (fp *LocalDiskFilePoolHandler) NextLine(fileName string) (string, error) {
	var line string
	scanner := fp.scanners[fileName]
	success := scanner.Scan()
	if !success {
		err := scanner.Err()
		if err != nil {
			// todo: is it really fatal?
			log.Fatal("Couldn't read line", err)
		}
		return "", err
	}
	line = scanner.Text()
	return line, nil
}

func (fp *LocalDiskFilePoolHandler) CloseFiles() {
	for _, file := range fp.files {
		err := file.Close()
		if err != nil {
			log.Fatal("Couldn't close file", err)
		}
	}
	log.Infof("Closed %v files", len(fp.files))

	fp.files = nil
}
