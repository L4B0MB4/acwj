package main

import (
	"fmt"
	"log"
	"os"
)

func usage(prog string) {
	log.Printf("Usage: %s <infile> <outfile>\n", prog)
	os.Exit(1)
}

func initVariables() {
	Line = 1
	Putback = '\n'
}
func cleanup() {
	if InputFilePtr != nil {
		fmt.Println("Closing the inputfile")
		InputFilePtr.Close()
	}
	if OutputFilePtr != nil {
		fmt.Println("Closing the outputfile")
		OutputFilePtr.Close()
	}
}

func main() {
	defer cleanup()
	if len(os.Args) != 3 {
		usage(os.Args[0])
	}
	initVariables()
	openInputFile(os.Args[1])
	openOutputFile(os.Args[2])
	scan(&T)
	n := *binExpr(0)
	printAstDepth(n)
	interpretAST(&n)
}
