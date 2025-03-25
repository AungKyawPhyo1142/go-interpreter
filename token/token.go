package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// defining our tokens
const (
	ILLEGAL = "ILLEGAL" // for the tokens we dont know of
	EOF     = "EOF"     // to determine if its end of file to tell the parser to stop

	// IDENTIFIERS & LITERALS
	IDENT = "IDENT" // add, foo, bar,..
	INT   = "INT"   // 1234

	// OPERATORS
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	// DELIMITERS
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// KEYWORDS
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

// check if identifier is a keyword
func LookupIndent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
