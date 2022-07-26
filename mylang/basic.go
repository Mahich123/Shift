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

    var types = [] string {};

    for i := 0;i < len(chars);i++ {
        tok += chars[i];

        // fmt.Println(tok)
        // fmt.Println(chars[i])

        if tok == " " {
            tok  = "";
        } else if tok == "print" {
            types = append(types, "TT_PRINT");
            tok = "";
        } else if chars[i] == "(" {
            types = append(types, "TT_LPAREN");
            tok = "";
        } else if chars[i] == ")" {
            types = append(types, "TT_RPAREN");
            tok = "";
        } else if chars[i] == "\"" {
            types = append(types, "TT_QUOTES");
            tok = "";
        } else if strings.ContainsAny(DIGITS, tok) == true {
            if chars[1 -1] == "\"" {
                types = append(types, "TT_STR");
                tok = "";
            } else {
                types = append(types, "TT_INT");
                tok = "";
            }
        } else if strings.ContainsAny(ALPABETS, tok) == true {
            types = append(types, "TT_STR");
            tok = "";
        }
    }

    fmt.Println(types)

    return 0;
}

func readFile(path string) string {
    var data, err = ioutil.ReadFile(path);

    if err != nil {
        return "";
    }

    return string(data);
}

func main() {
    if len(os.Args) > 1 {
        fmt.Println(Lexer(readFile(os.Args[1])));
    } else {
        fmt.Println("version: 1.0");
    }
}
