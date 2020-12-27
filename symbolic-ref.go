package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func Symbolic(args []string) {
	if len(args) < 3 || args[1] != "HEAD" {
		log.Fatal("error args...only support update HEAD now!")
		return
	}
	path := args[2]
	content := fmt.Sprintf("ref: %s", path)
	err := ioutil.WriteFile(filepath.Join(".git", "HEAD"), []byte(content), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
