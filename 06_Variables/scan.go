package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

func openInputFile(fileName string) {
	file, err := os.Open(fileName) // For read access.
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	InputFilePtr = file
	InputFile = bufio.NewReader(file)
}

func next() (rune, error) {
	var c rune
	if Putback > 0 {
		c = Putback
		Putback = 0
		return c, nil
	}

	c, _, err := InputFile.ReadRune()
	Column++
	if err != nil {
		if err == io.EOF {
			return 0, err
		} else {
			log.Fatal(err)
			panic(err)
		}
	}
	if c == '\n' {
		Line++
		Column = 0
	}
	return c, nil
}

func putback(c rune) {
	Putback = c
}

func skip() (rune, error) {
	c, err := next()
	if err != nil {
		return 0, err
	}
	for c == ' ' || c == '\t' || c == '\r' || c == '\f' {
		c, err = next()
		if err != nil {
			return 0, err
		}
	}
	return c, nil
}

func isCurrentTokenNewLine() bool {
	if T.token == T_NEWLINE {
		return true
	}
	return false
}

func matchToken(t int, expected string) {
	if T.token == t {
		scan(&T)
	} else {
		fmt.Printf("Expected %s on line %d column %d\n", expected, Line, Column)
	}
}

func matchIdent() {
	matchToken(T_IDENT, "identifier")
}

func matchNewLine() {
	matchToken(T_NEWLINE, "\\n")
}

func scanIdent(c rune) string {
	var identBuilder strings.Builder
	var err error
	for unicode.IsLetter(c) || unicode.IsDigit(c) || c == '_' {
		identBuilder.WriteRune(c)
		c, err = next()
		if err != nil {
			if err != io.EOF {
				log.Fatalf("Could not read identifier. Line %d Col %d Char: %q", Line, Column, c)

			}
			break
		}
	}
	putback(c)
	LastScannedIdentifier = identBuilder.String()
	return LastScannedIdentifier
}

func scanint(c rune) int {
	var err error
	var val = 0
	for unicode.IsDigit(c) {
		val = val*10 + int(c-'0')
		c, err = next()
		if err != nil {
			if err != io.EOF {
				log.Fatalf("Could not read integer. Line %d Col %d Char: %q", Line, Column, c)

			}
			break
		}
	}
	putback(c)
	return val
}

func getKeyword(ident string) int {
	switch ident[0] {
	case 'p':
		if ident == "print" {
			return T_PRINT
		}
		break
	case 'v':
		if ident == "v" {
			return T_VAR
		}
		break
	}
	return 0
}

func scan(t *Token) bool {
	c, err := skip()
	if err != nil {
		t.token = T_EOF
		return false
	}
	switch c {
	case '\n':
		t.token = T_NEWLINE
		break
	case '+':
		t.token = T_PLUS
		break
	case '-':
		t.token = T_MINUS
		break
	case '*':
		t.token = T_STAR
		break
	case '/':
		t.token = T_SLASH
		break
	case '=':
		t.token = T_EQ
		break
	default:
		if unicode.IsDigit(c) {
			t.intvalue = scanint(c)
			t.token = T_INTLIT
		} else if unicode.IsLetter(c) || '_' == c {
			// Read in a keyword or identifier
			identifier := scanIdent(c)
			tokentype := getKeyword(identifier)
			// If it's a recognised keyword, return that token
			if tokentype > 0 {
				t.token = tokentype
				break
			} else {
				//if its not a keyword, it must be an identifier
				t.token = T_IDENT
				break
			}
			// Not a recognised keyword, so an error for now
			fmt.Printf("Unrecognised symbol %s on line %d column %d \n", identifier, Line, Column)
			os.Exit(1)
		} else {
			fmt.Printf("Unrecognised character %q on Line %d Col %d \n", c, Line, Column)
			os.Exit(2)
		}
	}
	return true
}
