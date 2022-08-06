package Lexer

import (
	// "fmt"
	"strings"
)

const DIGITS = "0123456789";
const ALPABETS = "abcdefghijklmnopqrstuvwxyz";
const SYMBOLS = "";

func Lexer(code string) []string {

    var tok string = "";
    var chars = strings.Split(code, "");
    var types = [] string {};
    var __str string = "";
    var __int string = "";
    var sst int = 0;
    var ist int = 0;

    for i := 0;i < len(chars);i++ {
        tok += chars[i];

        switch tok {
            case "print":
                if sst == 0 {
                    types = append(types, "TT_PRINT");
                    tok = "";    
                }
        }

        if chars[i] == " " {
            if sst == 1 {
                __str += chars[i]
            } else {
                tok  = "";
            }
        } else if chars[i] == "\n" && sst == 0 {
            tok = "";
        } else if chars[i] == "#" && sst == 0 {
            return types
        } else if chars[i] == "(" && sst == 0 {
            types = append(types, "TT_LPAREN");
            tok = "";
        } else if chars[i] == ")" && sst == 0 {
            if ist == 1 {
                types = append(types, "TT_INT", __int);
                ist = 0
            }
            types = append(types, "TT_RPAREN");
            tok = "";
        } else if chars[i] == "\"" {
            if sst == 0 {
                sst = 1;
            } else if sst == 1 {
                types = append(types, "TT_STR:", __str);
                sst = 0;
                __str = "";
            }
            types = append(types, "TT_QUOTES");
            tok = "";
        } else if sst == 1 {
            __str += chars[i];
            tok = "";
        } else if sst == 0 && strings.ContainsAny(DIGITS, tok) {
            if strings.ContainsAny(ALPABETS, tok) {
                return types
            }
            ist = 1
            __int += chars[i];
        } else if sst == 0 && tok == "True" {
            types = append(types, "TT_BOOLEAN", "True")
            tok = "";
        } else if sst == 0 && tok == "False" {
            types = append(types, "TT_BOOLEAN", "False")
            tok = "";
        }
    }

    // fmt.Println(types)

    return types;
}
