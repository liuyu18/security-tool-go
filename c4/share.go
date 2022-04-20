package c4

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type FileNode struct {
	FullPath string
	Info     os.FileInfo
}

type filterFile interface {
	insertSorted(fileList *list.List, fileNode FileNode)
}

func getFilesInDirRecursivelyBySize(fileList *list.List, path string, work filterFile) {
	fmt.Println("path :", path)
	dirFiles, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println("Error reading directory: " + err.Error())
	}
	for _, dirFile := range dirFiles {
		fullPath := filepath.Join(path, dirFile.Name())
		if dirFile.IsDir() {
			getFilesInDirRecursivelyBySize(fileList, filepath.Join(path, dirFile.Name()), work)
		} else if dirFile.Mode().IsRegular() {
			work.insertSorted(fileList, FileNode{fullPath, dirFile})
		}
	}
}
