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
	"github.com/YumaMiyata910/hack-assembler-go/symboltable"
)

func main() {

	args := os.Args
	if len(args) < 2 {
		fmt.Println("引数に対象ファイルを指定してください。")
		os.Exit(1)
	}

	path := args[1]

	// read symboltable
	st, err := makeSymbolTable(path)
	if err != nil {
		log.Fatal(err)
	}

	// write .hack
	err = makeHack(path, st)
	if err != nil {
		log.Fatal(err)
	}

}

func makeSymbolTable(path string) (symboltable.SymbolTable, error) {
	readfile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("No such file: %s", path)
	}
	defer readfile.Close()

	sc := bufio.NewScanner(readfile)
	p := parser.NewParser(sc)

	st := symboltable.NewSymbolTable()
	var address int

	for p.HasMoreCommands() {
		if err = p.ScannerError(); err != nil {
			return nil, fmt.Errorf("ファイルの読み込みに失敗しました。path:【%s】", path)
		}

		p.Advance()
		if p.Text() == "" {
			continue
		}

		if p.CommandType() == parser.ACommand ||
			p.CommandType() == parser.CCommand {
			address += 1
		} else if p.CommandType() == parser.LCommand {
			sym := p.Symbol()
			st.AddEntry(sym, address)
		}
	}

	return st, nil
}

func makeHack(path string, st symboltable.SymbolTable) error {
	readfile, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("No such file: %s", path)
	}
	defer readfile.Close()

	sc := bufio.NewScanner(readfile)
	p := parser.NewParser(sc)

	filename := filepath.Base(path[:len(path)-len(filepath.Ext(path))])
	writefile, err := os.Create(filename + ".hack")
	if err != nil {
		return fmt.Errorf("新規ファイルを作成できません。 %v", err)
	}
	defer writefile.Close()

	var ram int = 16
	for p.HasMoreCommands() {
		if err = p.ScannerError(); err != nil {
			return fmt.Errorf("ファイルの読み込みに失敗しました。path:【%s】", path)
		}

		p.Advance()
		if p.Text() == "" {
			continue
		}

		var bin string
		if p.CommandType() == parser.ACommand {
			symbol := p.Symbol()
			val, err := strconv.Atoi(symbol)
			// 「err != nil」 is symbol.
			if err != nil {
				// 「!exists」 is new variable symbol.
				if st.Contains(symbol) {
					val, _ = st.GetAddress(symbol)
				} else {
					st.AddEntry(symbol, ram)
					val = ram
					ram += 1
				}
			}
			bin = fmt.Sprintf("%016b", val)
		} else if p.CommandType() == parser.CCommand {
			dest, err := code.Dest(p.Dest())
			comp, err := code.Comp(p.Comp())
			jump, err := code.Jump(p.Jump())
			if err != nil {
				log.Fatalln(err)
			}
			bin = fmt.Sprintf("111%s%s%s", comp, dest, jump)
		} else {
			continue
		}

		_, err = writefile.Write([]byte(bin + "\n"))
		if err != nil {
			return fmt.Errorf("ファイル書き込みに失敗しました。 %v", err)
		}
	}

	return nil
}
