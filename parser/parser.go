package parser

import (
	"bufio"
	"strings"
	"unicode"
)

const (
	A_COMMAND     string = "A_COMMAND"
	D_COMMAND     string = "D_COMMAND"
	L_COMMAND     string = "L_COMMAND"
	COMMENT_CHARS string = "//"
)

func stripComment(source string) string {
	if cut := strings.IndexAny(source, COMMENT_CHARS); cut >= 0 {
		return strings.TrimRightFunc(source[:cut], unicode.IsSpace)
	}
	return source
}

type Parser struct {
	scanner *bufio.Scanner
	text    string
}

func NewParser(scanner *bufio.Scanner) *Parser {
	return &Parser{scanner, ""}
}

func (p *Parser) HasMoreCommands() bool {
	return p.scanner.Scan()
}

func (p *Parser) ScannerError() error {
	return p.scanner.Err()
}

func (p *Parser) Text() string {
	return p.text
}

func (p *Parser) Advance() {
	p.text = stripComment(p.scanner.Text())
}

func (p *Parser) CommandType() string {
	if strings.HasPrefix(p.text, "@") {
		return A_COMMAND
	} else if strings.HasPrefix(p.text, "(") {
		return L_COMMAND
	} else {
		return D_COMMAND
	}
}

// func (p *Parser) Symbol() string {

// }

// func (p *Parser) Dest() string {

// }

// func (p *Parser) Comp() string {

// }

// func (p *Parser) Jump() string {

// }
