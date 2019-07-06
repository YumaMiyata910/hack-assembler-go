package code

import "fmt"

// Dest returns dest bin code.
func Dest(mnemonic string) (string, error) {
	var bin string
	var err error
	switch mnemonic {
	case "M":
		bin = "001"
	case "D":
		bin = "010"
	case "MD":
		bin = "011"
	case "A":
		bin = "100"
	case "AM":
		bin = "101"
	case "AD":
		bin = "110"
	case "AMD":
		bin = "111"
	case "":
		bin = "000"
	default:
		err = fmt.Errorf("Failed to converting dest binary. [mnemonic = %s]", mnemonic)
	}

	return bin, err
}

// Comp returns comp bin code.
func Comp(mnemonic string) (string, error) {
	var bin string
	var err error
	switch mnemonic {
	case "0":
		bin = "0101010"
	case "1":
		bin = "0111111"
	case "-1":
		bin = "0111010"
	case "D":
		bin = "0001100"
	case "A":
		bin = "0110000"
	case "!D":
		bin = "0001101"
	case "!A":
		bin = "0110001"
	case "-D":
		bin = "0001111"
	case "-A":
		bin = "0110011"
	case "D+1":
		bin = "0011111"
	case "A+1":
		bin = "0110111"
	case "D-1":
		bin = "0001110"
	case "A-1":
		bin = "0110010"
	case "D+A":
		bin = "0000010"
	case "D-A":
		bin = "0010011"
	case "A-D":
		bin = "0000111"
	case "D&A":
		bin = "0000000"
	case "D|A":
		bin = "0010101"
	case "M":
		bin = "1110000"
	case "!M":
		bin = "1110001"
	case "-M":
		bin = "1110011"
	case "M+1":
		bin = "1110111"
	case "M-1":
		bin = "1110010"
	case "D+M":
		bin = "1000010"
	case "D-M":
		bin = "1010011"
	case "M-D":
		bin = "1000111"
	case "D&M":
		bin = "1000000"
	case "D|M":
		bin = "1010101"
	default:
		err = fmt.Errorf("Failed to converting comp binary. [mnemonic = %s]", mnemonic)
	}

	return bin, err
}

// Jump returns jump bin code.
func Jump(mnemonic string) (string, error) {
	var bin string
	var err error
	switch mnemonic {
	case "JGT":
		bin = "001"
	case "JEQ":
		bin = "010"
	case "JGE":
		bin = "011"
	case "JLT":
		bin = "100"
	case "JNE":
		bin = "101"
	case "JLE":
		bin = "110"
	case "JMP":
		bin = "111"
	case "":
		bin = "000"
	default:
		err = fmt.Errorf("Failed to converting jump binary. [mnemonic = %s]", mnemonic)
	}

	return bin, err
}
