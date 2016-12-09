package lexer

import (
  "fmt"
  "github.com/makenneth/aves/token"
)

type Lexer struct {
  input string
  position int
  readPosition int
  ch byte
}

func New(input string) *Lexer {
  l := &Lexer{input: input}

  l.readChar()
  return l
}

func (this *Lexer) readChar() {
  fmt.Println(len(this.input))
  fmt.Println(this.readPosition)
  if this.readPosition >= len(this.input) {
    //0 is null in ASCII
    this.ch = 0
  } else {
    this.ch = this.input[this.readPosition]
  }

  this.position = this.readPosition
  this.readPosition++
}

func isLetter(ch byte) bool {
  return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '|'
}
func isDigit(ch byte) bool {
  return '0' <= ch && ch <= '9'
}

func (this *Lexer) NextToken() token.Token {
  var tok token.Token

  this.consumeWhitespace()

  switch (this.ch) {
    case '=':
      if this.peekChar() == '=' {
        this.readChar()
        tok = token.Token{Type: token.EQ, Literal: "=="}
      } else {
        tok = token.Token{Type: token.ASSIGN, Literal: string(this.ch)}
      }
    case '>':
      tok = token.Token{Type: token.GT, Literal: string(this.ch)}
    case '<':
      tok = token.Token{Type: token.LT, Literal: string(this.ch)}
    case '!':
      if this.peekChar() == '=' {
        this.readChar()
        tok = token.Token{Type: token.NOT_EQ, Literal: "!="}
      } else {
        tok = token.Token{Type: token.BANG, Literal: string(this.ch)}
      }
    case '/':
      //read comment
      tok = token.Token{Type: token.BANG, Literal: string(this.ch)}
    case '-':
      tok = token.Token{Type: token.MINUS, Literal: string(this.ch)}
    case '+':
      tok = token.Token{Type: token.PLUS, Literal: string(this.ch)}
    case ',':
      tok = token.Token{Type: token.COMMA, Literal: string(this.ch)}
    case ';':
      tok = token.Token{Type: token.SEMICOLON, Literal: string(this.ch)}
    case '(':
      tok = token.Token{Type: token.LPAREN, Literal: string(this.ch)}
    case ')':
      tok = token.Token{Type: token.RPAREN, Literal: string(this.ch)}
    case '{':
      tok = token.Token{Type: token.LBRACE, Literal: string(this.ch)}
    case '}':
      tok = token.Token{Type: token.RBRACE, Literal: string(this.ch)}
    case 0:
      tok = token.Token{Type: token.EOF, Literal: ""}
    case '.':
      if this.peekChar() == '.' {
        err := this.readSpread()
        if !err {
          tok.Type = token.ILLEGAL
        } else {
          tok.Type = token.SPREAD
        }
      }
    default:
      if isLetter(this.ch) {
        tok.Literal = this.readIdentifier()
        tok.Type = token.LookupKeyword(tok.Literal)
        return tok
      } else if isDigit(this.ch) {
        tok.Literal = this.readNumber()
        tok.Type = token.INT
        return tok
      } else {
        tok = token.Token{Type: token.ILLEGAL, Literal: string(this.ch)}
      }
  }
  this.readChar()
  return tok
}

func (this *Lexer) readSpread() bool {
  for i := 0; i < 2; i++ {
    this.readChar()
    if this.ch != '.' {
      return true
    }
  }

  return false
}

func (this *Lexer) readIdentifier() string {
    position := this.position
    for isLetter(this.ch) {
        this.readChar()
    }
    return this.input[position:this.position]
}

func (this *Lexer) readNumber() string {
    position := this.position
    for isDigit(this.ch) {
        this.readChar()
    }
    return this.input[position:this.position]
}

func (this *Lexer) consumeWhitespace() {
  for this.ch == ' ' || this.ch == '\t' || this.ch == '\n' || this.ch == '\r' {
    this.readChar()
  }
}

func (this *Lexer) peekChar() byte {
  if this.readPosition >= len(this.input) {
    return 0
  } else {
    return this.input[this.readPosition]
  }
}
