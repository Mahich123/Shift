# Shift (WIP)

## Get Started

An interpreted programming language designed and developed for general-purpose programming, with syntax similar to the English language. This language is not for production-level use. This language is implemented with [Go](https://go.dev/) and distributed under the [Creative Commons license](https://creativecommons.org/).

### Write your first program

- create a new `hello-world.shift` file and enter the following code and save the file
```js
print("Hello, World!")
```

### Executing the program

- open your terminal and type the following command
```
Shift hello-world.shift
```
- If everything is fine, you’ll see the following message on the screen
```
Hello, World!
```

## Process

- [ ] Currently, I'm rewriting the Lexer.

## Build

| OS/Architecture | Build |
| :---: | :---: |
| darwin/amd64 | ![passing](https://img.shields.io/badge/build-passing-brightgreen) |
| darwin/arm64 | ![passing](https://img.shields.io/badge/build-passing-brightgreen) |
| linux/amd64 | ![passing](https://img.shields.io/badge/build-passing-brightgreen) |
| linux/arm64 | ![passing](https://img.shields.io/badge/build-passing-brightgreen) |
| windows/amd64 | ![passing](https://img.shields.io/badge/build-passing-brightgreen) |
| windows/arm64 | ![passing](https://img.shields.io/badge/build-passing-brightgreen) |

## Building instructions

Install **Go** from the official [go.dev](https://go.dev/) website for your operating system.

- Clone the branch.
```
git clone https://github.com/sijey-praveen/Shift.git
```

- CD into the project
```
cd Shift
```

- Run the following command to build it
```
go build
```

## Project structure

```
.
├── Shift
│   ├── Lexer
|   |   ├── Lexer.go
│   ├── Parser
|   |   ├── Parser.go
├── go.mod
└── main.go
```
