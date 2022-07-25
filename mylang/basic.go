package main

import (
    "fmt"
    "strings"
)

const DIGITS = "0123456789";
const ALPABETS = "abcdefghijklmnopqrstuvwxyz";

var code = "print(\"Hello, World!\")";

func Lexer() int {

    var tok string = "";
    var chars = strings.Split(code, "");

    for i := 0;i < len(chars);i++ {
        tok += chars[i];
        
        // fmt.Println(tok)

        if tok == " " {
            tok  = "";
        } else if tok == "print" {
            fmt.Println("TT_PRINT was found!");
            tok = "";
        }

    }

    return 0;
}

func main() {
	fmt.Println(Lexer())
}
