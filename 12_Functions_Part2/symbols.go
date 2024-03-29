package main

import "errors"

var GlobalSymbols = []SymTable{}

type SymTable struct {
	name    string
	symType VariableType
}

func findGlobalSymbol(name string) (int, error) {
	for i := 0; i < len(GlobalSymbols); i++ {
		if GlobalSymbols[i].name == name {
			return i, nil
		}
	}
	return -1, errors.New("could not find global symbol")
}

func findLastGlobalSymbol() int {
	return len(GlobalSymbols) - 1
}
func addGlobalSymbol(name string, symType VariableType) int {
	GlobalSymbols = append(GlobalSymbols, SymTable{name: name, symType: symType})
	return len(GlobalSymbols) - 1
}
