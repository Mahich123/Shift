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
    var types = []string {};
    var __str string = "";
    var sst int = 0;

    for i := 0;i < len(chars);i++ {
        tok += chars[i];

        switch tok {
		case "print":
			if sst == 0 {
				types = append(types, "TT_PRINT");
				tok = "";    
			}
		case "True":
			if sst == 0 {
				types = append(types, "TT_BOOLEAN", "True");
				tok = "";
			}
		case "False":
			if sst == 0 {
				types = append(types, "TT_BOOLEAN", "False");
				tok = "";
			}
        }

		switch chars[i] {
		case "#":
			if sst == 0 {
				return types;
			}
		case "\n":
			if sst == 0 {
				tok = "";
			}
		case "(":
			if sst == 0 {
				types = append(types, "TT_LPAREN");
				tok = "";
			}
		case ")":
			if sst == 0 {
				types = append(types, "TT_RPAREN");
				tok = "";
			}
		}

        if chars[i] == " " {
            if sst == 1 {
                __str += chars[i];
            } else {
                tok  = "";
            }
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
        }
    }

    // fmt.Println(types)

    return types;
}
