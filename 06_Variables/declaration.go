package main

func varDeclerationStatement() {
	matchToken(T_VAR, "v")
	matchIdent()
	id := addGlobalSymbol(LastScannedIdentifier)
	genGlobalSymbol() //??

	if T.token == T_EQ {
		matchToken(T_EQ, "=")
		var left, right, tree *AstNode
		right = makeLeaf(A_ASSIGNVAL, id)
		left = binExpr(0)
		tree = makeAstNode(A_ASSIGN, left, right, 0)
		interpretAST(tree)
	}
}
