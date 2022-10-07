# lexer

def Lexer(code: str):
    integer = "1234567890"
    chars: list = [*code]
    token: list = []
    c: str = ""
    __: str = ""
    sst: int = 0

    for char in chars:
        c += char

        if c == "print" and sst == 0:
            token.append("TT_PRINT")
            c = ""
        elif char == " ":
            if sst == 1:
                pass
            else:
                c = ""
        elif char == '"':
            if sst == 1:
                sst = 0
                token.append("TT_STR")
                token.append(c[0:-1])
                token.append("TT_QUOTES")
                c = ""
            else:
                sst = 1
                token.append("TT_QUOTES")
                c = ""
        elif char == "#" and sst == 0:
            return token
        elif char == "\n" and sst == 0:
            c = ""
        elif char == "(" and sst == 0:
            token.append("TT_LPAREN")
            c = ""
        elif char == ")" and sst == 0:
            token.append("TT_RPAREN")
            c = ""
        elif char == "+" and sst == 0:
            token.append("TT_PLUS")
            c = ""
        elif char == "-" and sst == 0:
            token.append("TT_MINUS")
            c = ""
        elif char == "*" and sst == 0:
            token.append("TT_MUL")
            c = ""
        elif char == "/" and sst == 0:
            token.append("TT_DIV")
            c = ""
        elif char in integer and sst == 0:
            token.append("TT_INT")
            token.append(char)
            c = ""
        elif char == ">" and sst == 0:
            token.append("TT_GT")
            c = ""
        elif char == "<" and sst == 0:
            token.append("TT_LT")
            c = ""
        elif char == "[" and sst == 0:
            token.append("TT_LSQUARE")
            c = ""
        elif char == "]" and sst == 0:
            token.append("TT_RSQUARE")
            c = ""
        elif char == "," and sst == 0:
            token.append("TT_COMMA")
            c = ""
        elif c == "==" and sst == 0:
            token.append("TT_EQ")
            c = ""
        elif c == "!=" and sst == 0:
            token.append("TT_NE")
            c = ""
        elif c == ">=" and sst == 0:
            token.append("TT_GTE")
            c = ""
        elif c == "<=" and sst == 0:
            token.append("TT_LTE")
            c = ""

    return token