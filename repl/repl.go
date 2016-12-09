package repl

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "github.com/makenneth/aves/lexer"
  "github.com/makenneth/aves/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
  scanner := bufio.NewScanner(in)

  for {
    fmt.Printf(PROMPT)
    scanned := scanner.Scan()
    if !scanned {
      return
    }

    line := scanner.Text()
    l := lexer.New(line)

    for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
      if tok.Type == token.EXIT {
        os.Exit(0)
      }
      fmt.Printf("  %+v\n", tok)
    }
  }
}