package main

import (
  "fmt"
  "os"
  "github.com/makenneth/aves/repl"
)

func main() {
  fmt.Println("Kengo 0.0.1beta")
  fmt.Println("****************")
  fmt.Println("****************")
  repl.Start(os.Stdin, os.Stdout)
}