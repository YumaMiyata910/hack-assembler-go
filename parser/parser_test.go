package parser

import (
	"bufio"
	"os"
	"testing"
)

const testFile string = "test.asm"

func openFile(t *testing.T, path string) *os.File {
	readfile, err := os.Open(path)
	if err != nil {
		t.Fatal("Failed to open test file.")
	}
	return readfile
}

func closeFile(f *os.File) {
	f.Close()
}

func TestParser(t *testing.T) {
	f := openFile(t, testFile)
	defer closeFile(f)

	sc := bufio.NewScanner(f)
	p := NewParser(sc)

	t.Run("test dest", func(t *testing.T) {
		testDest(t, p, "D=M", "D")
		testDest(t, p, "M=A+1", "M")
		testDest(t, p, "AM=M-1;JEQ", "AM")
		testDest(t, p, "AMD=D|A;JGE", "AMD")
		testDest(t, p, "D;JMP", "")
	})

	t.Run("test comp", func(t *testing.T) {
		testComp(t, p, "D=M+1", "M+1")
		testComp(t, p, "AM=D&M", "D&M")
		testComp(t, p, "M=0", "0")
		testComp(t, p, "AMD=!A", "!A")
		testComp(t, p, "-M;JMP", "-M")
		testComp(t, p, "D+A;JNE", "D+A")
		testComp(t, p, "D;JGE", "D")
		testComp(t, p, "A=M;JGT", "M")
		testComp(t, p, "MD=M-D;JLE", "M-D")
		testComp(t, p, "D=1;JEQ", "1")
	})

	t.Run("test jump", func(t *testing.T) {
		testJump(t, p, "D;JGT", "JGT")
		testJump(t, p, "M;JEQ", "JEQ")
		testJump(t, p, "D=M-1;JEQ", "JEQ")
		testJump(t, p, "D=M", "")
	})
}

func testDest(t *testing.T, p *Parser, in string, expected string) {
	p.text = in
	dest := p.Dest()
	if dest != expected {
		t.Errorf("Dest() = %s, wanted %s", dest, expected)
	}
}

func testComp(t *testing.T, p *Parser, in string, expected string) {
	p.text = in
	comp := p.Comp()
	if comp != expected {
		t.Errorf("Comp() = %s, wanted %s", comp, expected)
	}
}

func testJump(t *testing.T, p *Parser, in string, expected string) {
	p.text = in
	jump := p.Jump()
	if jump != expected {
		t.Errorf("Jump() = %s, wanted %s", jump, expected)
	}
}
