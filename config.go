package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func readConfig(file string) (routes []route, err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()
	return parseConfig(f)
}

func parseConfig(r io.Reader) (routes []route, err error) {
	var lineNum int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lineNum++
		txt := scanner.Text()
		line := strings.TrimSpace(txt)
		if line == "" || line[0] == '#' {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			err = fmt.Errorf("Invalid line(%d): '%s'", lineNum, txt)
			return
		}
		routes = append(routes, route{parts[0], parts[1]})
	}
	err = scanner.Err()
	return
}
