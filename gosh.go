package main

import(
 "fmt"
 "bufio"
 "os"
 "io/ioutil"
)

func main() {
	ch := make(chan string)
	for {
		fmt.Printf("gosh> ")

		// Tokenize it here
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		// Dispatch it here
		switch input {
		        case "print\n": fmt.Printf("%s", <-ch)
			case "ls\n": ls(ch)
			default: nothing(ch, input)
		}
	}
}

func nothing(ch chan string, input string) {
	go func(){ ch <- fmt.Sprintf("not implemented!\n"); }()
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
