package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if args == nil || len(args) < 2 {
		fmt.Println(`
args error.
please reset the args.
		`)
		return
	}

	fmt.Println(`
[1] aaa
[2] bbb
[3] ccc
	`)

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter")

		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)

		if line == "exit" || len(line) == 0 {
			break
		}

		if line == "a" {
			fmt.Println("Test Success")
		}
	}
}

