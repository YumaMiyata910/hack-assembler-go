package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/YumaMiyata910/hack-assembler-go/code"
	"github.com/YumaMiyata910/hack-assembler-go/parser"
)

func main() {

	args := os.Args
	if len(args) < 2 {
		fmt.Println("引数に対象ファイルを指定してください。")
		os.Exit(1)
	}

	path := args[1]
	readfile, err := os.Open(path)
	if err != nil {
		log.Fatalf("No such file: %s", path)
	}
	defer readfile.Close()

	filename := filepath.Base(path[:len(path)-len(filepath.Ext(path))])
	writefile, err := os.Create(filename + ".hack")
	if err != nil {
		log.Fatalln("新規ファイルを作成できません。")
	}
	defer writefile.Close()

	sc := bufio.NewScanner(readfile)

	p := parser.NewParser(sc)

	for p.HasMoreCommands() {
		if err = p.ScannerError(); err != nil {
			log.Fatalf("ファイルの読み込みに失敗しました。path:【%s】", path)
		}

		p.Advance()
		if p.Text() == "" {
			continue
		}

		var bin string
		if p.CommandType() == parser.ACommand ||
			p.CommandType() == parser.LCommand {
			sym, err := strconv.Atoi(p.Symbol())
			if err != nil {
				log.Fatalln("symbolを数値変換できません。")
			}
			bin = fmt.Sprintf("%016b", sym)
		} else if p.CommandType() == parser.CCommand {
			dest, err := code.Dest(p.Dest())
			comp, err := code.Comp(p.Comp())
			jump, err := code.Jump(p.Jump())
			if err != nil {
				log.Fatalln(err)
			}
			bin = fmt.Sprintf("111%s%s%s", comp, dest, jump)
		}

		_, err = writefile.Write([]byte(bin + "\n"))
		if err != nil {
			log.Fatalln("ファイル書き込みに失敗しました。")
		}
	}
}
