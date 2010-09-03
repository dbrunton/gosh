package main

import(
 "fmt"
 "bufio"
 "os"
)

func main() {
 fmt.Println("You've started gosh.")
 fmt.Printf("%s", Prompt())
 reader := bufio.NewReader(os.Stdin)
 input, _ := reader.ReadBytes('\n')
 fmt.Print(input)
}

func Prompt() (prompt string) {
 prompt = "gosh> "
 return
}
