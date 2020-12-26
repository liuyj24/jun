package main

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func HashObject(t string, w bool, args []string) {
	path := args[len(args)-1]

	context := getContent(path)
	header := fmt.Sprintf("%s %d\u0000", t, len(context))
	data := append([]byte(header), context...)

	id := sha1.Sum(data)
	idStr := fmt.Sprintf("%x", id)
	fmt.Printf("%s\n", idStr)

	//write into object database
	prefix := idStr[:2]
	postfix := idStr[2:]

	objectDir := filepath.Join(".git", "objects", prefix)
	err := os.MkdirAll(objectDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	objectFile := filepath.Join(objectDir, postfix)
	file, err := os.Create(objectFile)
	if err != nil {
		log.Fatal(err)
	}

	//compress with zlib
	var b bytes.Buffer
	writer := zlib.NewWriter(&b)
	writer.Write(data)
	writer.Close()

	file.Write(b.Bytes())
}

func getContent(path string) []byte {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
