package main

import(
 "fmt"
 "bufio"
 "os"
 "io/ioutil"
)

func main() {
 fmt.Println("You've started gosh.")
 for {
  Prompt()
  reader := bufio.NewReader(os.Stdin)
  input, _ := reader.ReadBytes('\n')
  if string(input) == "ls\n" { Ls(); } else { fmt.Printf("%s", input); }
 }
}

func Prompt() {
 fmt.Print("gosh> ")
 return
}

func Ls() {
 files, _ := ioutil.ReadDir(".")
 for _, v := range files {
  fmt.Printf("%s\n", v.Name)
 }
}
