package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

//descripe an item in index and tree object
type EntryList struct {
	List []Entry
}

type Entry struct {
	Mode string
	Sha1 string
	Num  int
	Path string
	Type string
}

type Object interface {
	getContent() []byte
	getType() string
}

type BlobOjbect struct {
	Path string
	Sha1 string
	t    string
}

func (blob *BlobOjbect) getContent() []byte {
	file, err := ioutil.ReadFile(blob.Path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func (blob *BlobOjbect) getType() string {
	return "blob"
}

type TreeObject struct {
	List []Entry
	Sha1 string
	t    string
}

func (tree *TreeObject) getContent() []byte {
	var bytes bytes.Buffer
	for _, entry := range tree.List {
		bytes.WriteString(fmt.Sprintf("%s %s %s	%s\n", entry.Mode, entry.Type, entry.Sha1, entry.Path))
	}
	return bytes.Bytes()
}

func (tree *TreeObject) getType() string {
	return "tree"
}
