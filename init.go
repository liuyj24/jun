package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var FILE_MODE = os.ModePerm

func Init(path string) {
	if path == "" {
		return
	}

	//confirm the Path doesn't exist or the dir is empty
	if stat, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, FILE_MODE)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		if stat.IsDir() {
			dir, err := ioutil.ReadDir(path)
			if err != nil {
				log.Fatal(err)
			} else {
				if len(dir) != 0 {
					log.Fatalf("dir %v is not empty", path)
				}
			}
		} else {
			log.Fatalf("%v is a file", path)
		}
	}

	//create dir .git
	gitPath := filepath.Join(path, ".git")
	err := os.Mkdir(gitPath, FILE_MODE)
	if err != nil {
		log.Fatal(err)
	}

	//create file: config, description, HEAD
	os.Create(filepath.Join(gitPath, "config"))
	os.Create(filepath.Join(gitPath, "description"))
	head, _ := os.Create(filepath.Join(gitPath, "HEAD"))
	head.Write([]byte("ref: refs/heads/master"))

	//create dir: hooks, info, object, refs
	os.Mkdir(filepath.Join(gitPath, "hooks"), FILE_MODE)
	os.Mkdir(filepath.Join(gitPath, "info"), FILE_MODE)
	os.Mkdir(filepath.Join(gitPath, "objects"), FILE_MODE)
	os.Mkdir(filepath.Join(gitPath, "refs"), FILE_MODE)

	//create dir tags and heads in refs
	refsPath := filepath.Join(gitPath, "refs")
	os.Mkdir(filepath.Join(refsPath, "tags"), FILE_MODE)
	os.Mkdir(filepath.Join(refsPath, "heads"), FILE_MODE)

	dir, _ := os.Getwd()
	fmt.Printf("Initialized empty Git repository in :%s\n", filepath.Join(dir, gitPath))
}
