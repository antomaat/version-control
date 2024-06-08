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
    oid := GetHead()
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
