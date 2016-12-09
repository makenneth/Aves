package token

type TokenType string
type Token struct {
  Type TokenType
  Literal string
}

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  IDENT = "IDENT"
  INT = "INT"

  //operators
  ASSIGN = "="
  EQ = "=="
  NOT_EQ = "!="
  LT = "<"
  GT = ">"

  PLUS = "+"
  MINUS = "-"
  BANG = "!"
  ASTERISK = "*"
  SLASH = "/"
  SPREAD = "..."

  //syntax
  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  //keywords
  FUNCTION = "FUNCTION"
  LET = "LET"
  TRUE = "TRUE"
  FALSE = "FALSE"
  IF = "IF"
  ELSEIF = "ELSEIF"
  ELSE = "ELSE"
  RETURN = "RETURN"
  EXIT = "EXIT"
)

var keywords = map[string]TokenType{
  "func": FUNCTION,
  "let": LET,
  "true": TRUE,
  "false": FALSE,
  "if": IF,
  "else": ELSE,
  "elseif": ELSEIF,
  "return": RETURN,
  "exit": EXIT,
}

func LookupKeyword(keyword string) TokenType {
  if tok, ok := keywords[keyword]; ok {
    return tok
  }

  return IDENT
}