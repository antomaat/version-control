package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
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

func ReadTree(treeOid string) DirItem {
    treeDir := updateTree(treeOid, "./")
    writeTree(treeDir)
    return treeDir
}

func writeTree(dirItem DirItem) {
    for i := 0; i < len(dirItem.items); i++ {
	createOriginFile(dirItem.path + dirItem.items[i].name, GetObject(dirItem.items[i].oid, "blob"))
    }
    for i := 0; i < len(dirItem.dirItems); i++ {
	if len(dirItem.dirItems[i].items) > 0 || len(dirItem.dirItems) > 0 {
	    os.MkdirAll(dirItem.dirItems[i].path, os.ModePerm)
	    writeTree(dirItem.dirItems[i])
	}
    }
}

func createOriginFile(name string, data string) {
    file, err := os.Create(name)
    check(err)
    file.WriteString(data)
    file.Close()
}

func updateTree(oid string, basePath string) DirItem {
    dirItem := DirItem{
	items: []TreeItem{},
	dirItems: []DirItem{},
	path: basePath,
    }
    parsedLayer := parseTree(oid)
    for i := 0; i < len(parsedLayer); i++ {
	parsedType := strings.Map(func(r rune) rune {
	    if unicode.IsGraphic(r) {
		return r
	    }
	    return -1
	}, parsedLayer[i].itemType)
	if parsedType == "blob" {
	    dirItem.items = append(dirItem.items, parsedLayer[i])
	}
	if parsedType == "tree" {
	    dirItem.dirItems = append(dirItem.dirItems, updateTree(parsedLayer[i].oid, basePath + parsedLayer[i].name + "/"))
	}
    }
    return dirItem
}

func parseTree(oid string) []TreeItem {
    items := []TreeItem{}
    tree := strings.Split(GetObject(oid, "tree"), "\n")
    for i := 0; i < len(tree); i++ {
	if (len(tree[i]) == 0) {
	    break
	}

	splitItem := strings.Split(tree[i], " ")
	treeItem := TreeItem{
	    itemType: splitItem[0],
	    oid: splitItem[1],
	    name: splitItem[2],
	}
	items = append(items, treeItem)
    }
    return items
}

type DirItem struct {
    items []TreeItem;
    dirItems []DirItem;
    path string;
}

type TreeItem struct {
    itemType string;
    oid string;
    name string;
}
