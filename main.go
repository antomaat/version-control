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
    if args[0] == "status" {
        vcStatus(args[1:])
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

    if args[0] == "branch" {
        vcBranch(args[1:])
    }

    if args[0] == "k" {
        dot := vcVisualize()
        fmt.Println(dot)
    }
}

func vcInit() {
    fmt.Println("Initialize vc repository")
    InitSetup()
}

func vcStatus(args []string) {
    head := GetOid("@")
    branch := GetBranchName()
    if branch != "" {
        fmt.Printf("On branch %s\n", branch)
    } else {
        fmt.Printf("HEAD detached at %s\n", head[:10])
    }
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
    fmt.Println("start log")
    oid := GetOid("@")
    if (len(args) == 1) {
        oid = GetOid(args[0])
    }

    fmt.Printf("log start from oid %s\n", oid)

    oids := IterateCommitsAndParentsList([]string{oid})
    for i := 0; i < len(oids); i++ {
        fmt.Println()
        commit := GetCommit(oids[i])
        fmt.Printf("commit %s \n", oids[i])
        fmt.Println("   " + commit.message)
    }
}

func vcCheckout(args []string) {
    if len(args) == 0 {
        fmt.Println("commit name needed")
        return
    }
    Checkout(args[0])
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

func vcBranch(args []string) {
    arguments := createArguments(args)
    name := ""
    oid := "@"
    if arguments["-n"] != "" {
        name = arguments["-n"]
    }
    if arguments["-oid"] != "" {
        oid = arguments["-oid"]
    }
    if name != "" {
        oid = GetOid(oid)
        CreateBranch(name, oid)
        fmt.Printf("Branch %s created at %s \n", name, oid)
    } else {
        current := GetBranchName()
        branchNames := IterBranchNames()
        for i := 0; i < len(branchNames); i++ {
            prefix := ""
            if branchNames[i] == current {
                prefix = "* "
            }
            fmt.Printf("%s%s \n",prefix, branchNames[i])
        }
    }


}

func vcVisualize() string {
    dot := "digraph commits {\n"
    refs := IterateRefs("refs/tags/", false)
    oids := []string{}
    for i := 0; i < len(refs); i++ {
        dot += "\"" + refs[i].name + "\"" + "[shape=note]\n"
        dot += "\"" + refs[i].name + "\"" + "-> \"" + refs[i].commit.oid + "\"\n"
        fmt.Println(refs[i].name + "-> " + refs[i].commit.oid)
        oids = append(oids, refs[i].commit.oid)
    }
    iterOids := IterateCommitsAndParentsList(oids)
    for i := 0; i < len(iterOids); i++ {
        refAndComm := iterOids[i]
        commit := GetCommit(refAndComm)
        fmt.Println(commit.oid)
        dot += commit.oid + " [shape=box style=filled label = \"" + commit.oid[:10] + "\"]\n"
        if (commit.parent != "") {
            dot += "\"" + commit.oid + "\"" + "-> " + "\"" + commit.parent + "\"" +  "\n"
            fmt.Println(commit.parent)
        }
    }
    dot += "}"
    return dot
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
