package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var indexPath = filepath.Join(".git", "index")

func UpdateIndex(a bool, args []string) {
	if !a {
		log.Fatal("sorry, only support add file to the index now...")
	}
	path := args[len(args)-1]

	//create an object for the file content if the object is not exist
	var blob BlobOjbect
	blob.Path = path
	sha1, data := getSha1AndRawData(&blob)
	if exist := isObjectExist(sha1); !exist {
		writeObject(sha1, data)
	}

	//create file index
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		_, err := os.Create(indexPath)
		if err != nil {
			log.Fatal(err)
		}
	}

	//read entry-list from index
	entryList := getEntryListFromIndex()

	//return if the entry was existed
	for _, e := range entryList.List {
		if e.Sha1 == sha1 {
			return
		}
	}

	entry := Entry{"100644", sha1, 0, path, "blob"}
	entryList.List = append(entryList.List, entry)

	//write entry-list into index
	writeEntryListToIndex(entryList)
}

func isObjectExist(sha1 string) bool {
	dir, err := ioutil.ReadDir(filepath.Join(".git", "objects"))
	if err != nil {
		log.Fatal(err)
	}
	prefix := sha1[:2]
	//todo: binary search will be faster
	for _, v := range dir {
		if prefix == v.Name() {
			return true
		}
	}
	return false
}

func getEntryListFromIndex() *EntryList {
	bytes, err := ioutil.ReadFile(indexPath)
	if err != nil {
		log.Fatal(err)
	}
	var entryList EntryList
	if len(bytes) > 0 {
		bytes = unCompressData(bytes)
		err = json.Unmarshal(bytes, &entryList)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &entryList
}

func writeEntryListToIndex(entryList *EntryList) {
	bytes, err := json.Marshal(entryList)
	if err != nil {
		log.Fatal(err)
	}
	bytes = compress(bytes)
	err = ioutil.WriteFile(indexPath, bytes, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
