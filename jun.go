package main

import "flag"

func main() {
	//解析指令
	flag.Parse()

	//指令判断
	var command = flag.Arg(0)
	//args := flag.Args()[1:]

	switch command {
	case "init":
		Init(flag.Arg(1))
	}
}
