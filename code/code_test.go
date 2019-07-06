package code

import (
	"testing"
)

func TestDest(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"M", "001"},
		{"D", "010"},
		{"MD", "011"},
		{"A", "100"},
		{"AM", "101"},
		{"AD", "110"},
		{"AMD", "111"},
		{"", "000"},
	}

	for _, test := range tests {
		jump, err := Dest(test.in)

		if err != nil {
			t.Fatal(err)
		}

		if jump != test.out {
			t.Errorf("Dest(%s) = %s, want %s", test.in, jump, test.out)
		}
	}
}

func TestComp(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"0", "0101010"},
		{"1", "0111111"},
		{"-1", "0111010"},
		{"D", "0001100"},
		{"A", "0110000"},
		{"!D", "0001101"},
		{"!A", "0110001"},
		{"-D", "0001111"},
		{"-A", "0110011"},
		{"D+1", "0011111"},
		{"A+1", "0110111"},
		{"D-1", "0001110"},
		{"A-1", "0110010"},
		{"D+A", "0000010"},
		{"D-A", "0010011"},
		{"A-D", "0000111"},
		{"D&A", "0000000"},
		{"D|A", "0010101"},
		{"M", "1110000"},
		{"!M", "1110001"},
		{"-M", "1110011"},
		{"M+1", "1110111"},
		{"M-1", "1110010"},
		{"D+M", "1000010"},
		{"D-M", "1010011"},
		{"M-D", "1000111"},
		{"D&M", "1000000"},
		{"D|M", "1010101"},
	}

	for _, test := range tests {
		jump, err := Comp(test.in)

		if err != nil {
			t.Fatal(err)
		}

		if jump != test.out {
			t.Errorf("Comp(%s) = %s, want %s", test.in, jump, test.out)
		}
	}
}

func TestJump(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"JGT", "001"},
		{"JEQ", "010"},
		{"JGE", "011"},
		{"JLT", "100"},
		{"JNE", "101"},
		{"JLE", "110"},
		{"JMP", "111"},
		{"", "000"},
	}

	for _, test := range tests {
		jump, err := Jump(test.in)

		if err != nil {
			t.Fatal(err)
		}

		if jump != test.out {
			t.Errorf("Jump(%s) = %s, want %s", test.in, jump, test.out)
		}
	}
}
