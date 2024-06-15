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
        vcInit()
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
        vcWriteTree()
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

    if args[0] == "k" {
        vcVisualize()
    }
}

func vcInit() {
    fmt.Println("Initialize empty vc repository")
    InitNewRepository()
}

func vcHashObject(args []string) {
    if len(args) == 0 {
        fmt.Println("file name is needed")
    }
    readFile, _ := os.ReadFile(args[0]) 
    result := HashObject(string(readFile), "blob")
    fmt.Println(result)
}

func vcCatFile(args []string) {
    arguments := createArguments(args)
    oidValue := arguments["-oid"]
    expectedFileTypeValue := arguments["-type"]
    oid := "@"
    if oidValue != "" {
        oid = GetOid(args[0])
    }
    fmt.Println(GetObject(oid, expectedFileTypeValue))
}

func vcWriteTree() {
    fmt.Println(WriteTree("."))
}

func vcReadTree(args []string) {
    if len(args) == 0 {
        fmt.Println("need input argument oid/tag")
    }
    oid := "@"
    if len(args) > 0 {
        oid = GetOid(args[0])
    }
    fmt.Println(ReadTree(oid))
}

func vcCommit(args []string) {
    if args[0] == "-m" {
        fmt.Println(Commit(args[1]))
    }
}

func vcLog(args []string) {
    oid := "@"
    if (len(args) == 1) {
        oid = GetOid(args[0])
    } 

    for oid != "" {
        commit := GetCommit(oid)
        fmt.Printf("commit %s \n", oid)
        fmt.Println(commit.message)
        oid = commit.parent
    }
}

func vcCheckout(args []string) {
    oid := "@"
    if len(args) > 0 {
        oid = GetOid(args[0])
    }
    Checkout(oid)

}

func vcTag(args []string) {
    oid := "@"
    if len(args) == 0 {
    }
    arguments := createArguments(args)
    if arguments["-oid"] != "" {
        oid = arguments["-oid"]
    }
    name := arguments["-name"]
    oid = GetOid(oid)
    CreateTag(name, oid)
    fmt.Printf("oid: %s, name: %s\n", oid, name)
}

func vcVisualize() {
    refs := IterateRefs()
    for i := 0; i < len(refs); i++ {
        fmt.Println(refs[i])
    }
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
