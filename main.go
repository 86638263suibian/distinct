package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func main() {
	uniqs := map[string]struct{}{}
	h := md5.New()
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		h.Reset()
		line := s.Text()
		if _, err := io.WriteString(h, line); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		hash := fmt.Sprintf("%x", h.Sum(nil))
		if _, ok := uniqs[hash]; ok {
			continue
		}
		uniqs[hash] = struct{}{}
		fmt.Println(line)
	}
	if err := s.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
