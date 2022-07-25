const {readFile} = require("fs")

const DIGITS = "0123456789"
const ALPABETS = "abcdefghijklmnopqrstuvwxyz"

function Lexer() {
    return
}

function Token() {
    return
}

function print(value) {
    console.log(value)
}

readFile("src/test.txt", {encoding : "utf-8"}, (err, f) => {
    if (f) print(f)
    if (err) print(err)
})