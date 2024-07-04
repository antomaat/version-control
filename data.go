package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func InitNewRepository() {
    os.Mkdir(".vc", os.FileMode(0777))
    os.Mkdir(".vc/objects", os.FileMode(0777))
    os.Mkdir(".vc/refs", os.FileMode(0777))
    os.Mkdir(".vc/refs/tags", os.FileMode(0777))
    os.Mkdir(".vc/refs/heads", os.FileMode(0777))
}

func HashObject(data string, metaType string) string {
    hash := sha1.New()
    hash.Write([]byte(data))
    // I think the encoding should not be with URLEncoding
    // Will fix at later date if this does not make sense in the long run
    objId := base64.URLEncoding.EncodeToString(hash.Sum(nil))
    dataWithMeta := metaType + "\x00" + data
    createFile(string(objId), dataWithMeta)
    return string(objId)
}

func GetObject(hashString string, expectedType string) string {
    file, err := os.ReadFile("./.vc/objects/" + hashString)
    check(err)
    metaFields, fileData := separateMetaFields(string(file))
    if expectedType != "debug" && metaFields != expectedType {
	panic("meta value not what expected")
    }
    return fileData 
    
}

func UpdateRef(refName string, refValue RefValue, isDeref bool) {
    if refValue.value == "" {
	fmt.Printf("value of the ref is null\n")
    }
    ref, _ := GetRefInternal(refName, isDeref)
    fmt.Printf("ref internal %s\n", ref)
    file, err := os.Create("./.vc/" + ref)
    check(err)
    value := refValue.value
    if refValue.symbolic {
	value = "ref: " + value
    }
    file.WriteString(value)
    file.Close()
}

func UpdateRefInLocation(location string, refName string, refValue RefValue, isDeref bool) {
    ref, _ := GetRefInternal(refName, isDeref)
    os.MkdirAll("./.vc/" + location, os.ModePerm)
    file, err := os.Create("./.vc/"+ location + ref)
    check(err)
    value := refValue.value
    if refValue.symbolic {
	value = "ref: " + value
    }
    file.WriteString(value)
    file.Close()
}

func IterateRefs(startPoint string, isDeref bool) []RefItem {
    refs := []string{"HEAD"}
    items, _ := os.ReadDir("./.vc/" + startPoint)
    for i := 0; i < len(items); i++ {
	refs = append(refs, startPoint + items[i].Name())
    }
    refResults := []RefItem{}
    for i := 0; i < len(refs); i++ {
	if strings.HasPrefix(refs[i], startPoint) {
	    refResults = append(refResults, RefItem{name: refs[i], commit: GetCommit(GetRef(refs[i], isDeref).value)} )
	}
    }
    return refResults

}

func GetRef(name string, isDeref bool) RefValue {
    _, refValue := GetRefInternal(name, isDeref)
    return refValue
}

func GetRefInternal(name string, isDeref bool) (string, RefValue) {
    //fmt.Printf("ref internal name %s\n", name)
    file, err := os.ReadFile("./.vc/" + name)
    if err != nil {
	return name, RefValue{symbolic: false, value: ""}
    }
    value := string(file)
    isSymbolic := value != "" && strings.HasPrefix(value, "ref:") 
    if isSymbolic {
	value = strings.Trim(strings.Split(value, ":")[1], " ")
	if isDeref {
	    //fmt.Println("call recursively")
	    return GetRefInternal(value, true)
	}
    }
    return name, RefValue{
	symbolic: isSymbolic,
	value: value,
    } 
}

func createFile(name string, data string) {
    file, err := os.Create("./.vc/objects/" + name)
    check(err)
    file.WriteString(data)
    file.Close()
}

func separateMetaFields(file string) (string, string) {
    nullIndex := strings.Index(file, "\x00")
    return file[:nullIndex], file[nullIndex:]
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type RefValue struct {
    symbolic bool;
    value string;
}
