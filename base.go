package main

import (
	"fmt"
	"os"
)

func WriteTree(directory string) {
    items, _ := os.ReadDir(directory)
    // ignore the vc related directories
    if directory == "./.vc" {
	return
    }
    for i := 0; i < len(items); i++ {
	if items[i].IsDir() {
	    WriteTree(directory + "/" + items[i].Name())
	} else {
	    //fmt.Printf("%s/%s\n", directory, items[i].Name())
	    readFile, _ := os.ReadFile(fmt.Sprintf("%s/%s", directory, items[i].Name())) 
	    HashObject(string(readFile), "blob")
	}
    }
}
