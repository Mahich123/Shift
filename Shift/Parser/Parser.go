package Parser

import (
	"fmt"
)

func Parser(token [] string) int {

    for i := 0; i < len(token); i++ {
        if token[i] == "TT_PRINT" {
            if token[i + 1] == "TT_LPAREN" && token[len(token) - 1] == "TT_RPAREN" {
                if token[i + 2] == "TT_QUOTES" && token[len(token) - 2] == "TT_QUOTES" {
                    fmt.Println(token[i + 4])
                } else if token[i + 2] == "TT_INT" {
                    fmt.Println(token[i + 3])
                }
            }
        }
        i += 1;
    }

    return 0;
}