package main

import (
	"fmt"
	"os"
	"strings"
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
    if args[0] == "write-tree" {
        vcWriteTree(args[1:])
        return
    }
    if args[0] == "read-tree" {
        vcReadTree(args[1:])
        return
    }
    if args[0] == "commit" {
        vcCommit(args[1:])
        return
    }
    if args[0] == "log" {
        vcLog(args[1:])
    }
    if args[0] == "checkout" {
        vcCheckout(args[1:])
    }
    if args[0] == "tag" {
        vcTag(args[1:])
    }
}

func vcInit(args []string) {
    fmt.Println("Initialize empty vc repository")
    InitNewRepository()
}

func vcHashObject(args []string) {
    readFile, _ := os.ReadFile(args[0]) 
    result := HashObject(string(readFile), "blob")
    fmt.Println(result)
}

func vcCatFile(args []string) {
    fmt.Println(GetObject(args[0], args[1]))
}

func vcWriteTree(args []string) {
    fmt.Println(WriteTree("."))
}

func vcReadTree(args []string) {
    fmt.Println(ReadTree(args[0]))
}

func vcCommit(args []string) {
    if args[0] == "-m" {
        fmt.Println(Commit(args[1]))
    }
}

func vcLog(args []string) {
    oid := GetRef("HEAD")
    if (len(args) == 1) {
        oid = args[0]
    }

    for oid != "" {
        commit := GetCommit(oid)
        fmt.Printf("commit %s \n", oid)
        fmt.Println(commit.message)
        oid = commit.parent
    }
}

func vcCheckout(args []string) {
    if len(args) == 0 {
        fmt.Println("oid needed for checkout")
        return
    }
    oid := args[0]
    Checkout(oid)

}

func vcTag(args []string) {
    if len(args) == 0 {
        fmt.Println("at least the name should be provided")
        return
    }
    arguments := createArguments(args)
    oid := arguments["-oid"]
    name := arguments["-name"]
    if oid == "" {
        oid = GetRef("HEAD")
    }
    CreateTag(name, oid)
    fmt.Printf("oid: %s, name: %s\n", oid, name)
}

func createArguments(args [] string) map[string]string {
    arguments := make(map[string]string)
    for i := 0; i < len(args); i++ {
        if strings.HasPrefix(args[i], "-") {
            arguments[args[i]] = args[i+1]
        }
    }
    return arguments

}
