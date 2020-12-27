package main

import (
	"fmt"
	"log"
	"time"
)

func CommitTree(treeObjSha1 string, p string, m string, rest []string) *CommitOjbect {
	if treeObjSha1 == "" || len(treeObjSha1) < 4 {
		log.Fatalf("Not a valid object name %s\n", treeObjSha1)
	}
	var commitObj CommitOjbect

	// get the whole sha1 of the tree object
	exist, treeObjSha1 := isObjectExist(treeObjSha1)
	if !exist {
		log.Fatalf("Not a valid object name %s\n", treeObjSha1)
	}
	commitObj.treeObjSha1 = treeObjSha1

	if p != "" {
		exist, parentSha1 := isObjectExist(p)
		if !exist {
			log.Fatalf("The parent commit object is not exist!")
		}
		commitObj.parent = parentSha1
	}

	//in fact, we need to read info below from the config file ψ(._. )>
	commitObj.author = "liuyj24<liuyijun2017@email.szu.edu.cn>"
	commitObj.committer = "liuyj24<liuyijun2017@email.szu.edu.cn>"

	commitObj.date = fmt.Sprintf("%s", time.Now())
	commitObj.message = getMessage(m, rest)

	objSha1, data := getSha1AndRawData(&commitObj)
	commitObj.Sha1 = objSha1
	writeObject(objSha1, data)

	fmt.Printf("%s\n", objSha1)
	return &commitObj
}

func getMessage(m string, rest []string) string {
	for _, s := range rest {
		m += " " + s
	}
	return m
}
