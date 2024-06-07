package main

import (
	"crypto/sha1"
	"encoding/base64"
	"os"
	"strings"
)

func InitNewRepository() {
    os.Mkdir(".vc", os.FileMode(0777))
    os.Mkdir(".vc/objects", os.FileMode(0777))
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

func SetHead(oid string) {
    file, err := os.Create("./.vc/HEAD")
    check(err)
    file.WriteString(oid)
    file.Close()
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
