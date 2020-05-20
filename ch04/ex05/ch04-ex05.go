package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Parse()

	if len(flag.Args()) > 0 {
		fmt.Printf("%s\n", removeSameChar(flag.Arg(0)))
	} else {
		cmd := os.Args[0][strings.LastIndex(os.Args[0], "/")+1:]
		fmt.Fprintf(os.Stderr, "Usage %s: %s string\n", cmd, cmd)
		flag.PrintDefaults()
	}
}

func removeSameChar(s string) string {
	b := []byte(s)
	for i := 1; i < len(b); i++ {
		if b[i-1] == b[i] {
			copy(b[i-1:], b[i:])
			b = b[:len(b)-1]
			i--
		}
	}
	return string(b)
}
