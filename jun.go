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
	}
}
