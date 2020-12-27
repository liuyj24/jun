package main

import (
	"fmt"
	"time"
)

func CommitTree(treeObjSha1 string, p string, m string) {
	if treeObjSha1 == "" || len(treeObjSha1) < 4 {
		fmt.Printf("Not a valid object name %s\n", treeObjSha1)
		return
	}
	var commitObj CommitOjbect

	// get the whole sha1 of the tree object
	exist, treeObjSha1 := isObjectExist(treeObjSha1)
	if !exist {
		fmt.Printf("The tree object is not exist!")
		return
	}
	commitObj.treeObjSha1 = treeObjSha1

	if p != "" {
		exist, parentSha1 := isObjectExist(p)
		if !exist {
			fmt.Printf("The parent commit object is not exist!")
			return
		}
		commitObj.parent = parentSha1
	}

	//in fact, we need to read info below from the config file ψ(._. )>
	commitObj.author = "liuyj24<liuyijun2017@email.szu.edu.cn>"
	commitObj.committer = "liuyj24<liuyijun2017@email.szu.edu.cn>"

	commitObj.date = fmt.Sprintf("%s", time.Now())
	commitObj.message = m

	objSha1, data := getSha1AndRawData(&commitObj)
	commitObj.Sha1 = objSha1
	writeObject(objSha1, data)

	fmt.Printf("%s\n", objSha1)
}
