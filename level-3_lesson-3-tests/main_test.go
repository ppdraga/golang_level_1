package main

import (
	"bytes"
	"fmt"
	"log"
	"testing"
)

type DBStub struct {
	Storage bytes.Buffer
	cache   []string
}

func (db *DBStub) AddItem(item string) {
	_, err := db.Storage.WriteString(item + "\n")
	if err != nil {
		log.Fatalf("cannot write string: %v", err)
	}
	db.cache = append(db.cache, item)
}

func (db *DBStub) List() []string {
	return db.cache
}

func TestTodoer(t *testing.T) {

	var b bytes.Buffer
	b.Write([]byte(""))

	db := &DBStub{b, []string{}}
	db.AddItem("item1")
	db.AddItem("item2")
	fmt.Println(db, db.List())

	//t.Error("Fail: expected duplicates, got none")
	//t.Error("Fail: expected file name duplicate file2, got")

}
