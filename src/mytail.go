package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	outputRowCount = flag.Int("n", 10, "help message for \"n\" option")
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("引数を1つ選択して下さい。")
		return
	}
	workflow(args[0])
}

func workflow(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	crlfCount := 0
	var i int64 = 0
	for crlfCount < *outputRowCount + 1 {
		s, _ := f.Seek(-i, os.SEEK_END)
		if s == 0 {
			break
		}

		b := []byte{0}
		_, err = f.Read(b)
		if b[0] == '\n' {
			crlfCount++
		}
		i++
	}

  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    fmt.Println(scanner.Text())
	}
}
