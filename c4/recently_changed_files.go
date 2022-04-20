package c4

import (
	"container/list"
	"fmt"
)

type Recently struct {
}

func (this *Recently) insertSorted(fileList *list.List, fileNode FileNode) {
	if fileList.Len() == 0 { // If list is empty, just insert and return
		fileList.PushFront(fileNode)
		return
	}
	for element := fileList.Front(); element != nil; element = element.Next() {
		if fileNode.Info.ModTime().Before(element.Value.(FileNode).Info.ModTime()) {
			fileList.InsertBefore(fileNode, element)
			return
		}
	}
	fileList.PushBack(fileNode)
}

func FindRecentlyChangedFile() {
	fileList := list.New()
	getFilesInDirRecursivelyBySize(fileList, "/RyCode/security-go", new(Recently))

	for element := fileList.Front(); element != nil; element = element.Next() {
		fmt.Print(element.Value.(FileNode).Info.ModTime())
		fmt.Printf("%s\n", element.Value.(FileNode).FullPath)
	}
}
