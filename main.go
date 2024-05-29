package main

import (
	"fmt"
	"os"
)

func main() {
    args := os.Args[1:]

    if len(args) == 0 {
        fmt.Println("not enough arguments")
        return
    }
    
    if args[0] == "init" {
        vcInit(args[1:])
        return
    }
    if args[0] == "hash-object" {
        vcHashObject(args[1:])
        return
    }
    if args[0] == "cat-file" {
        vcCatFile(args[1:])
        return
    }
}

func vcInit(args []string) {
    fmt.Println("Initialize empty vc repository")
    InitNewRepository()
}

func vcHashObject(args []string) {
    readFile, _ := os.ReadFile(args[0]) 
    result := HashObject(string(readFile))
    fmt.Println(result)
}

func vcCatFile(args []string) {
    fmt.Println(GetObject(args[0]))
}
