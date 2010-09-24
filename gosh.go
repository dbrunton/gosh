package main

import(
 "fmt"
 "bufio"
 "os"
 "io/ioutil"
 "regexp"
)

func main() {
 fmt.Println("You've started gosh. Try typing 'ls'.")
 for {
  Prompt()
//  reader := bufio.NewReader(os.Stdin)
//  input, _ := reader.ReadBytes('\n')
  input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
  input = chomp(input)
  arg := "/home/dbrunton"
  switch {
   case (input == "cd"):
    os.Chdir(arg)
   case (input == "ls"):
    ls()
   case true:
    fmt.Printf("You typed '%s'.\nType something else.\n", input)
  }
 }
}

func Prompt() {
 fmt.Print("gosh> ")
 return
}

func ls() {
 files, _ := ioutil.ReadDir(".")
 for _, v := range files {
  fmt.Printf("%s\n", v.Name)
 }
}

func chomp(input string) string {
  rx, _ := regexp.Compile("\n");
  chompedInput := rx.ReplaceAllString(string(input), "")
  return chompedInput
}
