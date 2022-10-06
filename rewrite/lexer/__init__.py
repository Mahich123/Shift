# lexer

def Lexer(code:str):
    integer = "1234567890"
    chars : list = [*code]
    token : list = []
    c : str = ""
    __ : str = ""
    sst : int = 0

    for char in chars:
        c  += char

        if c == "print" and sst == 0:
            token.append("TT_PRINT")
            c = ""
        elif char == " ":
            if sst == 1:
                __ += char
                c = ""
            else:
                pass
        elif char == "\"":
            if sst == 1:
                sst = 0
                token.append("STR")
                token.append(__)
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
            token.append("TT_MULTIPLY")
            c = ""
        elif char == "/" and sst == 0:
            token.append("TT_DIVIDE")
            c = ""
        elif char in integer and sst == 0:
            token.append("TT_INT")
            token.append(char)
            c = ""
        elif char == ">" and sst == 0:
            token.append("TT_GREATER")
            c = ""
        elif char == "<" and sst == 0:
            token.append("TT_LESSER")
            c = ""
        elif c == "==" and sst == 0:
            token.append("TT_EQUALTO")
            c = ""     
        elif c == ">=" and sst == 0:
            token.append("TT_HIGHEQUAL")
            c = ""     
        elif c == "<=" and sst == 0:
            token.append("TT_LESSEQUAL")
            c = ""     
        elif c == "True" and sst == 0:
            token.append("TT_BOOLEAN")
            token.append("True")
            c = ""
        elif c == "False" and sst == 0:
            token.append("TT_BOOLEAN")
            token.append("False")
            c = ""
        else:
            if sst == 1:
                __ += char
                c = ""
        
    return token
