// Copyright 2025 The Hulo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package ast

import "github.com/hulo-io/fishparser/token"

type Node interface {
	Pos() token.Pos
	End() token.Pos
}

type Stmt interface {
	Node
	stmtNode()
}

type Decl interface {
	Node
	declNode()
}

type Expr interface {
	Node
	exprNode()
}

type CommentGroup struct {
	List []*Comment
}

func (g *CommentGroup) Pos() token.Pos { return g.List[0].Pos() }
func (g *CommentGroup) End() token.Pos { return g.List[len(g.List)-1].End() }

type Comment struct {
	Hash token.Pos // position of "#"
	Text string
}

func (c *Comment) Pos() token.Pos { return c.Hash }
func (c *Comment) End() token.Pos { return token.Pos(int(c.Hash) + len(c.Text)) }

// A FuncDecl node represents a function declaration.
type FuncDecl struct {
	Function token.Pos // position of "function"
	Name     *Ident
	Recv     []Expr
	Body     *BlockStmt
	EndPos   token.Pos // position of "end"
}

func (d *FuncDecl) Pos() token.Pos { return d.Function }

func (d *FuncDecl) End() token.Pos { return d.Body.Closing }

func (*FuncDecl) declNode() {}

type (
	// A AssignStmt node represents a assign statement.
	AssignStmt struct {
		Local  bool
		Lhs    Expr
		Assign token.Pos // position of "="
		Rhs    Expr
	}

	// A BlockStmt node represents a block statement.
	BlockStmt struct {
		Tok     token.Token // Token.NONE | Token.LBRACE
		Opening token.Pos
		List    []Stmt
		Closing token.Pos
	}

	// An ExprStmt node represents an expr statement.
	ExprStmt struct {
		X Expr
	}

	// A ReturnStmt node represents a return statement.
	ReturnStmt struct {
		Return token.Pos // position of "return"
		X      Expr
	}

	// A BreakStmt node represents a break statement.
	BreakStmt struct {
		Break token.Pos // position of "break"
	}

	// A ContinueStmt node represents a continue statement.
	ContinueStmt struct {
		Continue token.Pos // position of "continue"
	}

	// A WhileStmt node represents a while statement.
	WhileStmt struct {
		While  token.Pos // position of "while"
		Cond   Expr
		Body   *BlockStmt
		EndPos token.Pos // position of "end"
	}

	// A ForeachStmt node represents a foreach statement.
	ForeachStmt struct {
		For    token.Pos // position of "for"
		Elem   Expr
		In     token.Pos // position of "in"
		Group  []Expr
		Body   *BlockStmt
		EndPos token.Pos // position of "end"
	}

	// An IfStmt node represents an if statement.
	IfStmt struct {
		If     token.Pos // position of "if"
		Cond   Expr
		Body   *BlockStmt
		Elif   []*IfStmt
		Else   *BlockStmt
		EndPos token.Pos // position of "end"
	}

	// A SwitchStmt node represents a switch statement.
	SwitchStmt struct {
		Switch token.Pos // position of "switch"
		Lparen token.Pos // position of "("
		Var    Expr
		Rparen token.Pos // position of ")"
		Cases  []*CaseClause
		Else   *BlockStmt
		EndPos token.Pos // position of "end"
	}

	CaseClause struct {
		Case  token.Pos // position of "case"
		Conds []Expr
		Body  *BlockStmt
	}
)

func (s *AssignStmt) Pos() token.Pos { return s.Lhs.Pos() }
func (s *BlockStmt) Pos() token.Pos {
	if s.Opening.IsValid() {
		return s.Opening
	}
	if len(s.List) > 0 {
		return s.List[0].Pos()
	}
	return token.NoPos
}
func (s *ExprStmt) Pos() token.Pos     { return s.X.Pos() }
func (s *ReturnStmt) Pos() token.Pos   { return s.Return }
func (s *BreakStmt) Pos() token.Pos    { return s.Break }
func (s *ContinueStmt) Pos() token.Pos { return s.Continue }
func (s *WhileStmt) Pos() token.Pos    { return s.While }
func (s *ForeachStmt) Pos() token.Pos  { return s.For }
func (s *IfStmt) Pos() token.Pos       { return s.If }
func (s *SwitchStmt) Pos() token.Pos   { return s.Switch }

func (s *AssignStmt) End() token.Pos { return s.Rhs.End() }
func (s *BlockStmt) End() token.Pos {
	if s.Closing.IsValid() {
		return s.Closing
	}
	if len(s.List) > 0 {
		return s.List[len(s.List)-1].Pos()
	}
	return token.NoPos
}
func (s *ExprStmt) End() token.Pos     { return s.X.End() }
func (s *ReturnStmt) End() token.Pos   { return s.X.End() }
func (s *BreakStmt) End() token.Pos    { return s.Break }
func (s *ContinueStmt) End() token.Pos { return s.Continue }
func (s *WhileStmt) End() token.Pos    { return s.EndPos }
func (s *ForeachStmt) End() token.Pos  { return s.EndPos }
func (s *IfStmt) End() token.Pos       { return s.EndPos }
func (s *SwitchStmt) End() token.Pos   { return s.EndPos }

func (*AssignStmt) stmtNode()   {}
func (*BlockStmt) stmtNode()    {}
func (*ExprStmt) stmtNode()     {}
func (*ReturnStmt) stmtNode()   {}
func (*BreakStmt) stmtNode()    {}
func (*ContinueStmt) stmtNode() {}
func (*WhileStmt) stmtNode()    {}
func (*ForeachStmt) stmtNode()  {}
func (*IfStmt) stmtNode()       {}
func (*SwitchStmt) stmtNode()   {}

type (
	// Word string

	// A BinaryExpr node represents a binary expression.
	BinaryExpr struct {
		Compress bool
		X        Expr        // left operand
		OpPos    token.Pos   // position of Op
		Op       token.Token // operator
		Y        Expr        // right operand
	}

	// A CallExpr node represents a call expression.
	CallExpr struct {
		Func *Ident
		Recv []Expr
	}

	// A Ident node represents an identifier expression.
	Ident struct {
		NamePos token.Pos
		Name    string
	}

	// A BasicLit node represents a literal of basic type.
	BasicLit struct {
		Kind     token.Token // Token.Empty | Token.Null | Token.Boolean | Token.Byte | Token.Integer | Token.Currency | Token.Long | Token.Single | Token.Double | Token.Date | Token.String | Token.Object | Token.Error
		Value    string
		ValuePos token.Pos // literal position
	}

	// [ ]
	//
	// A BasicTestExpr node represents a basic test command expression.
	BasicTestExpr struct {
		Lbrack token.Pos // position of "["
		X      Expr
		Rbrack token.Pos // position of "]"
	}

	// [[ ]]
	//
	// An ExtendedTestExpr node represents an extended test command expression.
	ExtendedTestExpr struct {
		Lbrack token.Pos // position of "[["
		X      Expr
		Rbrack token.Pos // position of "]]"
	}

	// (( ))
	//
	// An ArithEvalExpr node represents an arithmetic evaluation expression.
	ArithEvalExpr struct {
		Lparen token.Pos // position of "(("
		X      Expr
		Rparen token.Pos // position of "))"
	}

	// Command Grouping: { }
	//
	// A CmdGroup node represents a command group expression.
	CmdGroup struct {
		Lbrace token.Pos // position of "{"
		List   []Stmt
		Rbrace token.Pos // position of "}"
	}

	// Command Substitution: $( ) or ` `
	//
	// A CmdSubst node represents a command substitution expression.
	CmdSubst struct {
		Dollar  token.Pos
		Tok     token.Token // Token.LPAREN or Token.BACK_QUOTE
		Opening token.Pos
		X       Expr
		Closing token.Pos
	}

	// Process Substitution: <( ) or >( )
	//
	// A ProcSubst node represents a process substitution expression.
	ProcSubst struct {
		Tok    token.Token // Token.GT or Token.LT
		TokPos token.Pos
		Lparen token.Pos // position of "("
		X      Expr
		Rparen token.Pos // position of ")"
	}

	// Arithmetic Expansion: $(())
	//
	// An ArithExp node represents an arithmetic expansion expression.
	ArithExp struct {
		Dollar token.Pos // position of "$"
		Lparen token.Pos // position of "(("
		X      Expr
		Rparen token.Pos // position of "))"
	}

	// Parameter Expandsion: ${}
	//
	// A ParamExp node represents an parameter expansion expression.
	ParamExp struct {
		Dollar               token.Pos // position of "$"
		Lbrace               token.Pos // position of "{"
		Var                  Expr
		*DefaultValExp                 // ${var:-val}
		*DefaultValAssignExp           // ${var:=val}
		*NonNullCheckExp               // ${var:?val}
		*NonNullExp                    // ${var:+val}
		*PrefixExp                     // ${!var*}
		*PrefixArrayExp                // ${!var@}
		*ArrayIndexExp                 // ${!var[*]} or ${!var[@]}
		*LengthExp                     // ${#var}
		*DelPrefix                     // ${var#val} or ${var##val}
		*DelSuffix                     // ${var%val} or ${var%%val}
		*SubstringExp                  // ${var:offset} or ${var:offset:length}
		*ReplaceExp                    // ${var/old/new}
		*ReplacePrefixExp              // ${var/#old/new}
		*ReplaceSuffixExp              // ${var/%old/new}
		*CaseConversionExp             // ${var,} or ${var^} or ${var,,} or ${var^^}
		*OperatorExp                   // ${var@op}
		Rbrace               token.Pos // position of "}"
	}

	// If parameter is unset or null, the expansion of word is substituted.
	// Otherwise, the value of parameter is substituted.
	DefaultValExp struct {
		Colon token.Pos
		Sub   token.Pos
		Val   Expr
	}

	// If parameter is unset or null, the expansion of word is assigned to parameter.
	// The value of parameter is then substituted.
	// Positional parameters and special parameters may not be assigned to in this way.
	DefaultValAssignExp struct {
		Colon  token.Pos
		Assign token.Pos
		Val    Expr
	}

	// If parameter is null or unset, the expansion of word (or a message to that effect if word is not present)
	// is written to the standard error and the shell, if it is not interactive, exits.
	// Otherwise, the value of parameter is substituted.
	NonNullCheckExp struct {
		Colon token.Pos
		Quest token.Pos
		Val   Expr
	}

	// If parameter is null or unset, nothing is substituted,
	// otherwise the expansion of word is substituted.
	NonNullExp struct {
		Colon token.Pos
		Add   token.Pos
		Val   Expr
	}

	PrefixExp struct {
		Bitnot token.Pos // position of '!'
		Mul    token.Pos // position of '*'
	}

	PrefixArrayExp struct {
		Bitnot token.Pos // position of '!'
		At     token.Pos // position of '@'
	}

	ArrayIndexExp struct {
		Bitnot   token.Pos   // position of '!'
		Lbracket token.Pos   // position of '['
		Tok      token.Token // Token.AT or Token.MUL
		TokPos   token.Pos
		Rbracket token.Pos // position of ']'
	}

	DelPrefix struct {
		Longest bool
		Hash    token.Pos // position of "#" or "##"
		Val     Expr
	}

	DelSuffix struct {
		Longest bool
		Mod     token.Pos // position of "%" or "%%"
		Val     Expr
	}

	SubstringExp struct {
		Colon1 token.Pos // position of ":"
		Offset int
		Colon2 token.Pos // position of ":"
		Length int
	}

	ReplaceExp struct {
		All  bool
		Div1 token.Pos // position of "/" or "//"
		Old  string
		Div2 token.Pos // position of "/"
		New  string
	}

	ReplacePrefixExp struct {
		Div1 token.Pos // position of "/"
		Hash token.Pos // position of "#"
		Old  string
		Div2 token.Pos // position of "/"
		New  string
	}

	ReplaceSuffixExp struct {
		Div1 token.Pos // position of "/"
		Mod  token.Pos // position of "%"
		Old  string
		Div2 token.Pos // position of "/"
		New  string
	}

	LengthExp struct {
		Hash token.Pos // position of "#"
	}

	CaseConversionExp struct {
		FirstChar bool
		ToUpper   bool
	}

	OperatorExp struct {
		At token.Pos // position of "@"
		Op ExpOperator
	}
)

// func (x Word) Pos() token.Pos              { return token.NoPos }
func (x *BinaryExpr) Pos() token.Pos       { return x.X.Pos() }
func (x *CallExpr) Pos() token.Pos         { return x.Func.NamePos }
func (x *Ident) Pos() token.Pos            { return x.NamePos }
func (x *BasicLit) Pos() token.Pos         { return x.ValuePos }
func (x *BasicTestExpr) Pos() token.Pos    { return x.Lbrack }
func (x *ExtendedTestExpr) Pos() token.Pos { return x.Lbrack }
func (x *ArithEvalExpr) Pos() token.Pos    { return x.Lparen }
func (x *CmdSubst) Pos() token.Pos         { return x.Dollar }
func (x *ProcSubst) Pos() token.Pos        { return x.TokPos }
func (x *ArithExp) Pos() token.Pos         { return x.Dollar }
func (x *ParamExp) Pos() token.Pos         { return x.Dollar }

// func (x Word) End() token.Pos        { return token.NoPos }
func (x *BinaryExpr) End() token.Pos { return x.Y.End() }
func (x *CallExpr) End() token.Pos {
	if len(x.Recv) > 0 {
		return x.Recv[len(x.Recv)-1].End()
	}
	return token.NoPos
}
func (x *Ident) End() token.Pos            { return token.Pos(int(x.NamePos) + len(x.Name)) }
func (x *BasicLit) End() token.Pos         { return token.Pos(int(x.ValuePos) + len(x.Value)) }
func (x *BasicTestExpr) End() token.Pos    { return x.Rbrack }
func (x *ExtendedTestExpr) End() token.Pos { return x.Rbrack }
func (x *ArithEvalExpr) End() token.Pos    { return x.Rparen }
func (x *CmdSubst) End() token.Pos         { return x.Closing }
func (x *ProcSubst) End() token.Pos        { return x.Rparen }
func (x *ArithExp) End() token.Pos         { return x.Rparen }
func (x *ParamExp) End() token.Pos         { return x.Rbrace }

// func (Word) exprNode()              {}
func (*BinaryExpr) exprNode()       {}
func (*CallExpr) exprNode()         {}
func (*Ident) exprNode()            {}
func (*BasicLit) exprNode()         {}
func (*BasicTestExpr) exprNode()    {}
func (*ExtendedTestExpr) exprNode() {}
func (*ArithEvalExpr) exprNode()    {}
func (*CmdSubst) exprNode()         {}
func (*ProcSubst) exprNode()        {}
func (*ArithExp) exprNode()         {}
func (*ParamExp) exprNode()         {}

type ExpOperator string

const (
	// The expansion is a string that is the value of parameter
	// with lowercase alphabetic characters converted to uppercase.
	ExpOperatorU = "U"

	// The expansion is a string that is the value of parameter
	// with the first character converted to uppercase,
	// if it is alphabetic.
	ExpOperatoru = "u"

	// The expansion is a string that is the value of parameter
	// with uppercase alphabetic characters converted to lowercase.
	ExpOperatorL = "L"

	// The expansion is a string that is the value of parameter quoted
	// in a format that can be reused as input.
	ExpOperatorQ = "Q"

	// The expansion is a string that is the value of parameter
	// with backslash escape sequences expanded as with the $'…' quoting mechanism.
	ExpOperatorE = "E"

	// The expansion is a string that is the result of expanding the value of parameter
	// as if it were a prompt string (see Controlling the Prompt).
	ExpOperatorP = "P"

	// The expansion is a string in the form of an assignment statement or declare command that,
	// if evaluated, will recreate parameter with its attributes and value.
	ExpOperatorA = "A"

	// Produces a possibly-quoted version of the value of parameter,
	// except that it prints the values of indexed and
	// associative arrays as a sequence of quoted key-value pairs (see Arrays).
	ExpOperatorK = "K"

	// The expansion is a string consisting of flag values representing parameter’s attributes.
	ExpOperatora = "a"

	// Like the ‘K’ transformation, but expands the keys and values of indexed and
	// associative arrays to separate words after word splitting.
	ExpOperatork = "k"
)

type File struct {
	Doc *CommentGroup

	Stmts []Stmt
	Decls []Decl
}

func (*File) Pos() token.Pos { return token.NoPos }
func (*File) End() token.Pos { return token.NoPos }
