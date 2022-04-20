package c4

import (
	"container/list"
	"fmt"
)

type LargestFile struct {
}

func (this *LargestFile) insertSorted(fileList *list.List, fileNode FileNode) {
	if fileList.Len() == 0 {
		fileList.PushFront(fileNode)
		return
	}
	for element := fileList.Front(); element != nil; element = element.Next() {
		if fileNode.Info.Size() < element.Value.(FileNode).Info.Size() {
			fileList.InsertAfter(fileNode, element)
			return
		}
	}
	fileList.PushBack(fileNode)
}

func RunFindLagestFile() {
	fileList := list.New()
	getFilesInDirRecursivelyBySize(fileList, "/home", new(LargestFile))
	for element := fileList.Front(); element != nil; element.Next() {
		fmt.Printf("%d ", element.Value.(FileNode).Info.Size())
		fmt.Printf("%s\n", element.Value.(FileNode).FullPath)
	}
}
