package main

import (
	"errors"
	"log"
	"os"
)

type SymTable struct {
	entries []*SymTableEntry
	parent  *SymTable
}

type SymTableEntry struct {
	name    string
	symType VariableType
}

var GlobalSymbols = SymTable{
	entries: []*SymTableEntry{},
	parent:  nil,
}

var currentScope *SymTable

func newScope() {
	currentScope = &SymTable{
		entries: []*SymTableEntry{},
		parent:  currentScope,
	}
}
func popScope() {
	currentScope = currentScope.parent
}

func addSymbolToCurrentScope(name string, symType VariableType) {
	currentScope.entries = append(currentScope.entries, createSymTableEntry(name, symType))
}

func findSymbol(name string, scope *SymTable) (*SymTableEntry, error) {

	for i := 0; i < len(scope.entries); i++ {
		if scope.entries[i].name == name {
			return scope.entries[i], nil
		}
	}
	if scope.parent != nil {
		return findSymbol(name, scope.parent)
	}
	return nil, errors.New("could not find symbol")
}

func getAllSymbolsFromScope(scope *SymTable) []*SymTableEntry {
	return scope.entries
}

func createSymTableEntry(name string, symType VariableType) *SymTableEntry {
	return &SymTableEntry{name: name, symType: symType}
}

func getSymbolFromAst(n *AstNode) *SymTableEntry {
	id := n.v.id
	symbol, err := findSymbol(id, n.symTable)
	if err != nil {
		log.Fatalf("Unknown symbol %s. %v", id, err)
		os.Exit(12)
	}
	return symbol
}
