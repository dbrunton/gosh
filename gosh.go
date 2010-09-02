package main

import(
 "fmt"
)

func main() {
 fmt.Println("You've started gosh.")
 fmt.Printf("%s", Prompt())
}

func Prompt() (prompt string) {
 prompt = "gosh> "
 return
}
