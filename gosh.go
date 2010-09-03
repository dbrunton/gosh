package main

import(
 "fmt"
 "bufio"
 "os"
)

func main() {
 fmt.Println("You've started gosh.")
 for {
  fmt.Printf("%s", Prompt())
  reader := bufio.NewReader(os.Stdin)
  input, _ := reader.ReadBytes('\n')
  fmt.Print(input, "\n")
 }
}

func Prompt() (prompt string) {
 prompt = "gosh> "
 return
}
