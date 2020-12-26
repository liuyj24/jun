package main

import (
	"fmt"
	"strings"
)

func WriteTree() {
	entryList := getEntryListFromIndex()
	treeObj := getTreeObject(entryList.List)
	fmt.Printf("%s\n", treeObj.Sha1)
}

func getTreeObject(list []Entry) *TreeObject {

	var treeObject TreeObject

	m := make(map[string][]Entry)
	var blobList []Entry

	// sort out blobs by their Path
	for _, entry := range list {

		//os.PathSeparator may be better here, but I am on windows and test in shell, so "/" make sense here
		index := strings.Index(entry.Path, "/")
		if index != -1 {
			dirPath := entry.Path[:index]
			entry.Path = entry.Path[index+1:]
			m[dirPath] = append(m[dirPath], entry)
		} else {
			blobList = append(blobList, entry)
		}
	}

	for _, entry := range blobList {
		var newEntry Entry
		newEntry.Path = entry.Path
		newEntry.Sha1 = entry.Sha1
		newEntry.Type = entry.Type
		newEntry.Mode = entry.Mode
		newEntry.Num = entry.Num
		treeObject.List = append(treeObject.List, newEntry)
	}

	for k, v := range m {
		var newEntry Entry
		childTreeObj := getTreeObject(v)

		newEntry.Path = k
		newEntry.Mode = "040000"
		newEntry.Type = "tree"
		newEntry.Sha1 = childTreeObj.Sha1
		treeObject.List = append(treeObject.List, newEntry)
	}

	//write tree object to object database
	objSha1, data := getSha1AndRawData(&treeObject)
	treeObject.Sha1 = objSha1

	writeObject(objSha1, data)
	return &treeObject
}
