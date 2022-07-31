package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const DIGITS = "0123456789";
const ALPABETS = "abcdefghijklmnopqrstuvwxyz";

func Lexer(code string) []string {

    var tok string = "";
    var chars = strings.Split(code, "");
    var types = [] string {};
    var __str string = ""
    var st int = 0

    for i := 0;i < len(chars);i++ {
        tok += chars[i];

        if chars[i] == " " {
            if st == 1 {
                __str += chars[i]
            } else {
                tok  = "";

            }
        } else if tok == "\n" {
            tok = "";
        } else if tok == "print" && st == 0 {
            types = append(types, "TT_PRINT");
            tok = "";
        } else if chars[i] == "(" && st == 0 {
            types = append(types, "TT_LPAREN");
            tok = "";
        } else if chars[i] == ")" && st == 0 {
            types = append(types, "TT_RPAREN");
            tok = "";
        } else if chars[i] == "\"" {
            if st == 0 {
                st = 1;
            } else if st == 1 {
                types = append(types, "TT_STR:", __str);
                st = 0;
                __str = "";
            }
            types = append(types, "TT_QUOTES");
            tok = "";
        } else if st == 1 {
            __str += chars[i];
            tok = "";
        }
    }

    return types;
}

func parser(token [] string) int {

    for i := 0; i < len(token); i++ {
        if token[i] == "TT_PRINT" {
            if token[i + 1] == "TT_LPAREN" && token[len(token) - 1] == "TT_RPAREN" {
                if token[i + 2] == "TT_QUOTES" && token[len(token) - 2] == "TT_QUOTES" {
                    fmt.Println(token[i + 4])
                }
            }
        }
        i += 1;
    }

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
        var lines = strings.Split(readFile(os.Args[1]), "\n");
        for i := 0; i < len(lines); i++ {
            parser(Lexer(lines[i]));
        }
    } else {
        fmt.Println("version: 1.0");
    }
}
