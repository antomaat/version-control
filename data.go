package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"os"
)

func InitNewRepository() {
    os.Mkdir(".vc", os.FileMode(0777))
    os.Mkdir(".vc/objects", os.FileMode(0777))
}

func HashObject(data string) string {
    hash := sha1.New()
    hash.Write([]byte(data))
    // I think the encoding should not be with URLEncoding
    // Will fix at later date if this does not make sense in the long run
    objId := base64.URLEncoding.EncodeToString(hash.Sum(nil))
    fmt.Println("object created")
    fmt.Println(string(objId))
    createFile(string(objId), data)
    return string(objId)
}

func createFile(name string, data string) {
    file, err := os.Create("./.vc/objects/" + name)
    check(err)
    file.WriteString(data)
    file.Close()
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
