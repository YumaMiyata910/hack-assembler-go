package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/YumaMiyata910/hack-assembler-go/parser"
)

func main() {

	args := os.Args
	if len(args) < 2 {
		fmt.Println("引数に対象ファイルを指定してください。")
		os.Exit(1)
	}

	path := args[1]
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("No such file: %s", path)
		os.Exit(1)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	parser := parser.NewParser(sc)

	for parser.HasMoreCommands() {
		if err = parser.ScannerError(); err != nil {
			fmt.Printf("ファイルの読み込みに失敗しました。path:【%s】", path)
			os.Exit(1)
		}

		parser.Advance()
		if parser.Text() == "" {
			continue
		}

		fmt.Printf("%s: %s\n", parser.Text(), parser.CommandType())
	}

}
