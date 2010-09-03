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
  reader := bufio.NewReader(os.Stdin)
  input, _ := reader.ReadBytes('\n')
  rx, _ := regexp.Compile("\n");
  parsedInput := rx.ReplaceAllString(string(input), "")
  switch {
   case parsedInput == "ls":
    fmt.Println("Yeah, baby, we know how to ls."
    ls()
   case true:
    fmt.Printf("You typed '%s'.\nType something else.\n", parsedInput)
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
