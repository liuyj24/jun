package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func Log(args []string) {
	var curObjsha1 string
	if len(args) <= 1 {
		curObjsha1 = logWithoutArgs()
	} else {
		curObjsha1 = logWithArgs(args)
	}
	objStr := getCatFileStr(true, false, false, []string{curObjsha1})
	var commitObj CommitOjbect
	commitObj.Sha1 = curObjsha1
	commitObj.parseCommitObj([]byte(objStr))

	var buf bytes.Buffer
	printLog(&commitObj, &buf)
	fmt.Printf("%s", buf.Bytes())
}

func logWithoutArgs() string {
	//get objSha1 from HEAD
	bytes, err := ioutil.ReadFile(filepath.Join(".git", "HEAD"))
	if err != nil {
		log.Fatal(err)
	}
	s := string(bytes)
	i := strings.Index(s, " ")
	path := s[i+1:]

	sha1bytes, err := ioutil.ReadFile(filepath.Join(".git", path))
	if err != nil {
		log.Fatal(err)
	}
	return string(sha1bytes)
}

func logWithArgs(args []string) string {
	argStr := args[1]
	exist, curObjsha1 := isObjectExist(argStr)
	if !exist {
		exist, curObjsha1 = isRefExist(argStr)
		if !exist {
			log.Fatalf("Not a valid object name %s\n", argStr)
		}
	}
	return curObjsha1
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
