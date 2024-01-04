package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/no-src/nsgo/browser"
	"github.com/no-src/toolkit/config"
)

func main() {
	run()
}

func run() {
	if len(os.Args) < 2 {
		return
	}
	name := strings.ToLower(os.Args[1])
	num := 1
	if len(os.Args) > 2 {
		if parseNum, err := strconv.Atoi(os.Args[2]); err == nil && parseNum > 0 {
			num = parseNum
		}
	}
	for _, t := range config.Tools {
		ok := slices.ContainsFunc(t.Names, func(s string) bool {
			return strings.ToLower(s) == name
		})
		if ok {
			if len(t.Urls) < num {
				num = 1
			}
			url := t.Urls[num-1]
			browser.OpenBrowser(url)
			fmt.Printf("found %s: %s\n", name, url)
			return
		}
	}
	fmt.Printf("%s not found\n", name)
}
