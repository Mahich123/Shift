# py .\src\main.py .\src\test.txt

from sys import argv

def Lexer(src___):
    print(src___.split("\n"))


if __name__ == "__main__":    
    if argv:
        Lexer(open(argv[1], "r", encoding="utf-8").read())