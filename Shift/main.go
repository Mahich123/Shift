package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
    "Shift/Lexer"
    "Shift/Parser"
)

func main() {
    if len(os.Args) > 1 {
		var data, err = ioutil.ReadFile(os.Args[1]);

		if err != nil {
			fmt.Println("");
		}

		var lines = strings.Split(string(data), "\n");

		for i := 0; i < len(lines); i++ {
            Parser.Parser(Lexer.Lexer(lines[i]));
        }
    } else {
        fmt.Println("Shift version: 1.0a (https://github.com/sijey-praveen/Shift)");
    }
}
