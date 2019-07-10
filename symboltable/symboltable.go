package symboltable

import "fmt"

// SymbolTable has symbol and address mapping.
type SymbolTable map[string]int

// NewSymbolTable create new SymbolTable list.
func NewSymbolTable() SymbolTable {
	return SymbolTable{
		"SP":     0,
		"LCL":    1,
		"ARG":    2,
		"THIS":   3,
		"THAT":   4,
		"R0":     0,
		"R1":     1,
		"R2":     2,
		"R3":     3,
		"R4":     4,
		"R5":     5,
		"R6":     6,
		"R7":     7,
		"R8":     8,
		"R9":     9,
		"R10":    10,
		"R11":    11,
		"R12":    12,
		"R13":    13,
		"R14":    14,
		"R15":    15,
		"SCREEN": 16384,
		"KBD":    24576,
	}
}

// AddEntry is mapping symbol and address.
func (st SymbolTable) AddEntry(symbol string, address int) {
	st[symbol] = address
}

// Contains returns symbol including.
func (st SymbolTable) Contains(symbol string) bool {
	_, ok := st[symbol]
	return ok
}

// GetAddress returns address of symbol.
func (st SymbolTable) GetAddress(symbol string) (int, error) {
	var err error
	address, ok := st[symbol]
	if !ok {
		err = fmt.Errorf("This symbol doesn't exist. [symbol: %s]", symbol)
	}
	return address, err
}
