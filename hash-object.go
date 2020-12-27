package main

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func HashObject(t string, w bool, args []string) {
	path := args[len(args)-1]

	var blob BlobOjbect
	blob.Path = path

	objSha1, data := getSha1AndRawData(&blob)
	blob.Sha1 = objSha1
	fmt.Printf("%s\n", blob.Sha1)

	//write into database
	writeObject(objSha1, data)
}

func getSha1AndRawData(o Object) (string, []byte) {
	content := o.getContent()
	header := fmt.Sprintf("%s %d\u0000", o.getType(), len(content))
	data := append([]byte(header), content...)
	s := fmt.Sprintf("%x", sha1.Sum(data))
	return s, data
}

// compress the object and write into the database
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
