package main

import (
	"bytes"
	"fmt"
)

func Log(objSha1 string) {
	if objSha1 == "" || len(objSha1) < 1 {
		fmt.Printf("Not a valid object name %s\n", objSha1)
		return
	}
	exist, curObjsha1 := isObjectExist(objSha1)
	if !exist {
		fmt.Printf("Not a valid object name %s\n", objSha1)
		return
	}
	objStr := getCatFileStr(true, false, false, []string{curObjsha1})
	var commitObj CommitOjbect
	commitObj.Sha1 = curObjsha1
	commitObj.parseCommitObj([]byte(objStr))

	var buf bytes.Buffer
	printLog(&commitObj, &buf)
	fmt.Printf("%s", buf.Bytes())
}

func printLog(commit *CommitOjbect, buf *bytes.Buffer) {
	buf.WriteString(fmt.Sprintf("commit %s\n", commit.Sha1))
	buf.WriteString(fmt.Sprintf("Author: %s\n", commit.author))
	buf.WriteString(fmt.Sprintf("Date:	%s\n", commit.date))
	buf.WriteString(fmt.Sprintf("\n"))
	buf.WriteString(fmt.Sprintf("%s\n", commit.message))
	buf.WriteString(fmt.Sprintf("\n"))

	if commit.parent != "" {
		exist, parentSha1 := isObjectExist(commit.parent)
		if exist {
			var parent CommitOjbect
			parent.Sha1 = parentSha1
			objStr := getCatFileStr(true, false, false, []string{parentSha1})
			parent.parseCommitObj([]byte(objStr))
			printLog(&parent, buf)
		}
	}
}
