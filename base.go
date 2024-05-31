package main

import (
	"fmt"
	"os"
)

func WriteTree(directory string) string {
    items, _ := os.ReadDir(directory)
    // ignore the vc related directories
    dirItems := []string{}
    for i := 0; i < len(items); i++ {
	if items[i].IsDir() {
	    if items[i].Name() != ".vc" {
		oidc := WriteTree(directory + "/" + items[i].Name())
		dirItems = append(dirItems, fmt.Sprintf("tree %s %s", oidc, items[i].Name()))
	    }
	} else {
	    //fmt.Printf("%s/%s\n", directory, items[i].Name())
	    readFile, _ := os.ReadFile(fmt.Sprintf("%s/%s", directory, items[i].Name())) 
	    oidc := HashObject(string(readFile), "blob")
	    dirItems = append(dirItems, fmt.Sprintf("blob %s %s", oidc, items[i].Name()))
	}
    }

    treeResult := ""
    for i := 0; i < len(dirItems);  i++ {
	treeResult += fmt.Sprintf("%s \n", dirItems[i])
    }

    return HashObject(treeResult, "tree")
}
