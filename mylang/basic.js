const {readFile} = require("fs")

const DIGITS = "0123456789"
const ALPABETS = "abcdefghijklmnopqrstuvwxyz"
const DECIMAL = "0.00"

function Lexer(code) {
    return
}

function Token(code) {
    const tt = {
        True : "TT_BOOLEAN",
        False : "TT_BOOLEAN",
        None : "TT_None",
        def : "TT_DEF",
        return : "TT_RETURN",
    }
    
    return tt
}

readFile("src/test.txt", {encoding : "utf-8"}, (err, f) => {
    if (f) console.log(f)
    if (err) console.log(err)
})
