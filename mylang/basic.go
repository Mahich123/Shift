package main

import (
    "fmt"
    "strings"
    "io/ioutil"
    "os"
)

const DIGITS = "0123456789";
const ALPABETS = "abcdefghijklmnopqrstuvwxyz";

func Lexer(code string) int {

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

func readFile(path string) string {
    var data, err = ioutil.ReadFile(path)

    if err != nil {
        return ""
    }

    return string(data)
}

func main() {
    if len(os.Args) > 1 {
        fmt.Println(Lexer(readFile(os.Args[1])))
    } else {
        fmt.Println("version: 1.0")
    }
}
