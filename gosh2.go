package main

import(
	"fmt"
	"bufio"
	"os"
	"io/ioutil"
	"regexp"
)

func main() {
	ch := make(chan string)

	for {
		fmt.Printf("gosh> ")

		// Tokenize it here
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		terms := split(input)

		// Dispatch it here
		switch input {
		        case "print\n": fmt.Printf("%s", <-ch)
			case "ls\n": ls(ch)
			default: nothing(ch, terms)
		}
	}
}

func nothing(ch chan string, terms []string) {
	go func(){
		ch <- fmt.Sprintf("not implemented %+v", terms)
	}()
}

func ls(ch chan string) {
	go func(){
		files, _ := ioutil.ReadDir(".")
		ret := ""
		for _, v := range files {
			ret += fmt.Sprintf("%s\n", v.Name)
		}
		ch <- ret
	}()
}

func split(input string) []string {
	rx := regexp.MustCompile("[^ ]+")
	return rx.FindAllString(input, -1)
}
