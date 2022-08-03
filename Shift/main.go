package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
    "Shift/Lexer"
    "Shift/Parser"
)

func readFile(path string) string {
    var data, err = ioutil.ReadFile(path);

    if err != nil {
        return "";
    }

    return string(data);
}

func main() {
    if len(os.Args) > 1 {
        var lines = strings.Split(readFile(os.Args[1]), "\n");
        for i := 0; i < len(lines); i++ {
            Parser.Parser(Lexer.Lexer(lines[i]));
        }
    } else {
        fmt.Println("version: 1.0");
    }
}