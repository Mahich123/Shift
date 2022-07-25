const { readFile } = require("fs")

const DIGITS = "0123456789"
const ALPABETS = "abcdefghijklmnopqrstuvwxyz"

function Lexer(code) {
    let tok = ""
    let chars = code.split("")

    chars.forEach(char => {
        tok += char
        if (tok == " ") {
            tok  = ""
        }
        else if (tok == "print") {
            console.log("TT_PRINT was found!")
            tok = ""
        }
    })
}

function Token(code) {
    const tt = {
        True : "TT_BOOLEAN",
        False : "TT_BOOLEAN",
        None : "TT_None",
        def : "TT_DEF",
        return : "TT_RETURN",
        print : "TT_PRINT",
    }

    return tt
}

readFile("test.txt", {encoding : "utf-8"}, (e, f) => {
    if (f) Lexer(f)
    if (e) console.error(e)
})
