package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

const LogFileExtension = "log"

type localDiskFilePoolHandler struct {
	files map[string]*os.File
}

func newLocalDiskFilePoolHandler(dir string) *localDiskFilePoolHandler {
	files := make(map[string]*os.File)
	fileNames, err := filepath.Glob(filepath.Join(dir, fmt.Sprintf("*.%s", LogFileExtension)))
	if err != nil {
		log.Fatal("Couldn't find log files", err)
	}
	log.Infof("Found log %v files", len(fileNames))

	for _, fileName := range fileNames {
		fileHandle, err := os.Open(fileName)
		if err != nil {
			log.Fatal("Couldn't open file", err)
		}
		files[fileName] = fileHandle
	}
	return &localDiskFilePoolHandler{files: files}
}

func (fp *localDiskFilePoolHandler) NextLine(fileName string) (string, error) {
	var line string
	file := fp.files[fileName]
	_, err := fmt.Fscanln(file, &line)
	return line, err
}

func (fp *localDiskFilePoolHandler) closeFiles() {
	for _, file := range fp.files {
		err := file.Close()
		if err != nil {
			log.Fatal("Couldn't close file", err)
		}
	}
	log.Infof("Closed %v files", len(fp.files))

	fp.files = nil
}