package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Commit(m string, rest []string) {
	tree := WriteTree()

	refPathBytes, err := ioutil.ReadFile(filepath.Join(".git", "HEAD"))
	if err != nil {
		log.Fatal(err)
	}
	refPathStr := string(refPathBytes)
	i := strings.LastIndex(refPathStr, "/")

	refName := refPathStr[i+1:]
	exist, parentSha1 := isRefExist(refName)
	if !exist {
		log.Fatal("missing initial commit...")
	}
	m = getMessage(m, rest)
	commit := CommitTree(tree.Sha1, parentSha1, m, []string{})

	i = strings.Index(refPathStr, " ")
	refPath := refPathStr[i+1:]
	err = ioutil.WriteFile(filepath.Join(".git", refPath), []byte(commit.Sha1), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
