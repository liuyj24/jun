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
	idStr := getSha1Str(path, t)
	fmt.Printf("%s\n", idStr)

	data := getData(path, t)
	writeObject(idStr, data)
}

func getSha1Str(path string, t string) string {
	data := getData(path, t)
	id := sha1.Sum(data)
	return fmt.Sprintf("%x", id)
}

//assemble according to the data format of the object
func getData(path string, t string) []byte {
	content := getContent(path)
	header := fmt.Sprintf("%s %d\u0000", t, len(content))
	data := append([]byte(header), content...)
	return data
}

func getContent(path string) []byte {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

//compress the object and write into the database
func writeObject(idStr string, data []byte) {
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
	compressedData := compress(data)
	file.Write(compressedData)
}

func compress(raw []byte) []byte {
	var b bytes.Buffer
	writer := zlib.NewWriter(&b)
	writer.Write(raw)
	writer.Close()
	return b.Bytes()
}
