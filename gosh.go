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
  command, arg := parse(input)
  switch {
   case (command == "cd"):
    os.Chdir(arg)
   case (command == "ls"):
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

func parse(input string) (string, string) {
	rx := regexp.MustCompile("[^ ]*")
	output := rx.FindAllString(input, 2)
  	if(len(output) == 1) {
  		return output[0], ""
	}
	return output[0], output[1]
}
