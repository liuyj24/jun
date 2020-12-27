package main

import (
	"fmt"
	"os"
)

func LsFiles(s bool) {
	if !s {
		fmt.Printf("ls-files only support --stage now!")
		return
	}
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		return
	}
	//read entry-list from index
	entryList := getEntryListFromIndex()

	for _, entry := range entryList.List {
		fmt.Printf("%s %s %d	%s\n", entry.Mode, entry.Sha1, entry.Num, entry.Path)
	}
}
