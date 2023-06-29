package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
)

const LogFileExtension = "log"

type LocalDiskFilePoolHandler struct {
	files map[string]*os.File
}

func NewLocalDiskFilePoolHandler(dir string) *LocalDiskFilePoolHandler {
	files := make(map[string]*os.File)

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
	}
	return &LocalDiskFilePoolHandler{files: files}
}

func (fp *LocalDiskFilePoolHandler) NextLine(fileName string) (string, error) {
	var line string
	file := fp.files[fileName]
	_, err := fmt.Fscanln(file, &line)
	return line, err
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
