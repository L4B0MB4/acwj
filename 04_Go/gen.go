package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var VariableCounter = 0

func openOutputFile(fileName string) {
	file, err := os.Create(fileName) // For read access.
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	OutputFilePtr = file
	OutputFile = bufio.NewWriter(file)
}

func generateVariable() string {
	variableName:=""
	numberChars := int(VariableCounter / 26)
	for i:=0;i<numberChars;i++{
		variableName+=
	}
}

func genAdd(left, right AstNode) {
	fmt.Fprintf(OutputFile, "(%v + %v) ", left.intval, right.intval)
}
func genSub(left, right AstNode) {
	fmt.Fprintf(OutputFile, "(%v - %v) ", left.intval, right.intval)
}
func genMul(left, right AstNode) {
	fmt.Fprintf(OutputFile, "(%v * %v) ", left.intval, right.intval)
}
func genDiv(left, right AstNode) {
	fmt.Fprintf(OutputFile, "(%v / %v) ", left.intval, right.intval)
}
