package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
		OutputFile.Flush()
		fmt.Println("Closing the outputfile")
		OutputFilePtr.Close()
	}
}

func compileFile(path string) {
	log.Printf(path)
	cmd := exec.Command("go", "build", "-o", "build.exe", path)
	log.Printf("Compiling...")
	err := cmd.Run()
	if err != nil {
		log.Printf("Compiling finished with error: %v", err)
	}
	log.Printf("Successfully compiled")
	cmd = exec.Command(filepath.Join(filepath.Base(path), "build.exe"))
	cmd.Run()
}

func main() {
	generateVariable()
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
	genMainFuncStart()
	interpretAST(&n)
	genMainFuncEnd()
	compileFile(os.Args[2])
}
