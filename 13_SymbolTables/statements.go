package main

import (
	"log"
	"os"
)

func singleStatement() *AstNode {
	var tree *AstNode = nil
	switch T.token {
	case T_PRINT:
		tree = printStatement()

	case T_VAR:
		tree = varDeclerationStatement()

	case T_FUNC:
		tree = functionDeclaration()

	case T_IDENT:
		tree = identifierStatement()

	case T_IF:
		tree = ifStatement()

	case T_WHILE:
		tree = whileStatement()

	case T_FOR:
		tree = forStatement()

	case T_NODE:
		tree = nodeStatement()

	default:
		log.Fatalf("Unknown token Line %d Column %d", Line, Column)
	}
	return tree
}

func compundStatement() *AstNode {
	var left *AstNode = nil
	var tree *AstNode

	// match as many newlines as you want before new statement
	skipNewLine()

	newScope()
	defer popScope()
	matchLBrace()

	for {
		// match as many newlines as you want before new statement
		skipNewLine()
		tree = singleStatement()

		if tree != nil {
			if left == nil {
				left = tree
			} else {
				left = makeAstNode(A_GLUETO, left, nil, tree, -1, "")
			}
		}

		skipNewLine()

		if T.token == T_RBRACE {
			matchRBrace()
			return left
		}

	}
}

func nodeStatement() *AstNode {
	var state, input *AstNode
	matchToken(T_NODE, "node")
	matchDot()
	matchIdent()
	id := LastScannedIdentifier
	addSymbolToCurrentScope(id, TYPE_NODE)
	matchLBrace()
	for {
		skipNewLine()
		state = stateStatement()
		skipNewLine()
		//input = functionDeclaration()
		skipNewLine()

		if T.token == T_RBRACE {
			matchRBrace()
			break
		}
	}
	return makeAstNode(A_NODE, state, input, nil, -1, id)
}

func stateStatement() *AstNode {
	var tree *AstNode
	matchToken(T_STATE, "state")
	matchLBrace()
	skipNewLine()
	tree = varDeclerationStatement()
	skipNewLine()
	matchRBrace()
	return tree
}

func forStatement() *AstNode {
	var conditionAst, bodyAst, preOpAst, postOpAst, tree *AstNode
	matchToken(T_FOR, "for")
	matchLParen()
	preOpAst = singleStatement()
	matchSemi()

	conditionAst = binExpr(0)
	if conditionAst.op < A_EQ || conditionAst.op > A_GE {
		log.Fatalf("Condition is not returning a boolean on Line %d", Line)
	}

	matchSemi()

	postOpAst = singleStatement()

	matchRParen()
	bodyAst = compundStatement()

	// Glue the compound statement and the postop tree
	tree = makeAstNode(A_GLUETO, bodyAst, nil, postOpAst, -1, "")

	// Make a WHILE loop with the condition and this new body
	tree = makeAstNode(A_WHILE, conditionAst, tree, nil, -1, "")

	// And glue the preop tree to the A_WHILE tree
	return (makeAstNode(A_GLUETO, preOpAst, nil, tree, -1, ""))

}

func whileStatement() *AstNode {
	var conditionAst, bodyAst *AstNode
	matchToken(T_WHILE, "while")
	matchLParen()
	conditionAst = binExpr(0)

	if conditionAst.op < A_EQ || conditionAst.op > A_GE {
		log.Fatalf("Condition is not returning a boolean on Line %d", Line)
	}
	matchRParen()
	bodyAst = compundStatement()
	return makeAstNode(A_WHILE, conditionAst, bodyAst, nil, -1, "")
}

func ifStatement() *AstNode {
	var conditionAst, trueAst, falseAst *AstNode
	matchToken(T_IF, "if")
	matchLParen()
	conditionAst = binExpr(0)

	if conditionAst.op < A_EQ || conditionAst.op > A_GE {
		log.Fatalf("Condition is not returning a boolean on Line %d", Line)
	}
	matchRParen()
	trueAst = compundStatement()
	if T.token == T_ELSE {
		scan(&T)
		falseAst = compundStatement()
	}
	return makeAstNode(A_IF, conditionAst, trueAst, falseAst, -1, "")
}

func printStatement() *AstNode {
	var tree *AstNode
	matchToken(T_PRINT, "print")
	tree = binExpr(0)
	return makeAstUnary(A_PRINT, tree, -1, "")
}

func assignStatement(symbolId string) *AstNode {
	var left, right, tree *AstNode
	left = makeLeaf(A_ASSIGNVAL, -1, symbolId)
	matchToken(T_ASSIGN, "=")
	right = binExpr(0)
	tree = makeAstNode(A_ASSIGN, left, nil, right, 0, "")

	return tree

}

func identifierStatement() *AstNode {
	matchIdent()
	id := LastScannedIdentifier
	symbol, err := findSymbol(id, currentScope)
	if err != nil {
		log.Fatalf("Couldn't find global Symbol %s", LastScannedIdentifier)
		os.Exit(7)
	}
	switch symbol.symType {
	case TYPE_FUNC:
		return functionCallStatement(id)
	case TYPE_INT:
		return assignStatement(id)
	default:
		log.Fatalf("Unkown type for identifier %v", symbol.name)
		os.Exit(9)
	}
	return nil
}

func functionCallStatement(symbolId string) *AstNode {
	matchLParen()
	matchRParen()
	var tree *AstNode
	makeLeaf(A_FUNC_CALL, -1, symbolId)
	return makeAstUnary(A_FUNC_CALL, tree, -1, symbolId)
}
