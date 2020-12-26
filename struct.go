package main

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

type TreeObject struct {
	List []Entry
	sha1 string
}
