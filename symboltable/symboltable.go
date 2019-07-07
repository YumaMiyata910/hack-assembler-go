package symboltable

import "fmt"

// SymbolTable has symbol and address mapping.
type SymbolTable map[string]int

// NewSymbolTable create new SymbolTable list.
func NewSymbolTable() SymbolTable {
	return make(SymbolTable, 0)
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
