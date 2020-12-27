package main

import (
	"flag"
	"os"
)

func main() {
	//confirm command
	var command = os.Args[1]
	os.Args = os.Args[1:] //for using flag, I have to do this... no more better idea now

	switch command {
	case "init":
		flag.Parse()
		Init(flag.Arg(0))

	case "hash-object":
		w := flag.Bool("w", true, "write into object database")
		t := flag.String("t", "blob", "object type")
		flag.Parse()
		HashObject(*t, *w, flag.Args())

	case "cat-file":
		p := flag.Bool("p", false, "print object content")
		t := flag.Bool("t", false, "show object type")
		s := flag.Bool("s", false, "show object size")
		flag.Parse()
		CatFile(*p, *t, *s, flag.Args())

	case "update-index":
		a := flag.Bool("add", false, "add file content to the index")
		flag.Parse()
		UpdateIndex(*a, flag.Args())

	case "ls-files":
		s := flag.Bool("stage", false, "Show staged contents' mode bits, object name and stage number")
		flag.Parse()
		LsFiles(*s)

	case "write-tree":
		WriteTree()

	case "commit-tree":
		treeObjSha1 := os.Args[1]
		os.Args = os.Args[1:]
		p := flag.String("p", "", "indicates the id of a parent commit object")
		m := flag.String("m", "", "the commit log message")
		flag.Parse()
		CommitTree(treeObjSha1, *p, *m, flag.Args())

	case "log":
		Log(os.Args)

	case "update-ref":
		UpdateRef(os.Args)

	case "symbolic-ref":
		Symbolic(os.Args)

	case "commit":
		m := flag.String("m", "", "commit message")
		flag.Parse()
		Commit(*m, flag.Args())
	}
}
