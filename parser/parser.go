package parser

import (
	"bufio"
	"strings"
	"unicode"
)

const (
	// ACommand is A_COMMAND.
	ACommand string = "A_COMMAND"
	// CCommand is C_COMMAND.
	CCommand string = "C_COMMAND"
	// LCommand is L_COMMAND.
	LCommand string = "L_COMMAND"
	// CommentChars is comment symbol.
	CommentChars string = "//"
)

// Parser is Assembly parser.
type Parser struct {
	scanner *bufio.Scanner
	text    string
	command string
	symbol  string
}

func stripComment(source string) string {
	if cut := strings.IndexAny(source, CommentChars); cut >= 0 {
		return strings.TrimFunc(source[:cut], unicode.IsSpace)
	}
	return source
}

// NewParser create new struct.
func NewParser(scanner *bufio.Scanner) *Parser {
	return &Parser{scanner, "", "", ""}
}

// HasMoreCommands scanning next row.
func (p *Parser) HasMoreCommands() bool {
	return p.scanner.Scan()
}

// ScannerError returns scannings error.
func (p *Parser) ScannerError() error {
	return p.scanner.Err()
}

// Text return row text.
func (p *Parser) Text() string {
	return p.text
}

// Advance read current row text.
func (p *Parser) Advance() {
	text := stripComment(p.scanner.Text())
	p.text = strings.ReplaceAll(text, " ", "")
}

// CommandType returns command types.
func (p *Parser) CommandType() string {
	if strings.HasPrefix(p.text, "@") {
		p.command = ACommand
	} else if strings.HasPrefix(p.text, "(") {
		p.command = LCommand
	} else {
		p.command = CCommand
	}
	return p.command
}

// Symbol returns symbol from current text.
func (p *Parser) Symbol() string {
	switch p.command {
	case ACommand:
		p.symbol = strings.TrimLeft(p.text, "@")
	case LCommand:
		sym := strings.TrimLeft(p.text, "(")
		p.symbol = strings.TrimRight(sym, ")")
	}
	return p.symbol
}

// Dest returns dest mnemonic.
func (p *Parser) Dest() string {
	if cut := strings.IndexAny(p.text, "="); cut >= 0 {
		return p.text[:cut]
	}
	return ""
}

// Comp returns comp mnemonic.
func (p *Parser) Comp() string {
	dest := 0
	jump := len(p.text)
	if cut := strings.IndexAny(p.text, "="); cut >= 0 {
		dest = cut + 1
	}
	if cut := strings.IndexAny(p.text, ";"); cut >= 0 {
		jump = cut
	}
	return p.text[dest:jump]
}

// Jump returns jump mnemonic.
func (p *Parser) Jump() string {
	if cut := strings.IndexAny(p.text, ";"); cut >= 0 {
		return p.text[cut+1:]
	}
	return ""
}
