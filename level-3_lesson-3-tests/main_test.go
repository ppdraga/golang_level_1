package main

import (
	"bytes"
	"fmt"
	"github.com/chzyer/readline"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"os"
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
	//fmt.Println(db, db.List())

	// input
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Println(err)
	}
	_, err = w.Write([]byte("add item3\nadd item4\nlist\nunknown\n"))
	if err != nil {
		fmt.Println(err)
	}
	w.Close()

	todoer, err := NewTodoerServiceImpl(db, &readline.Config{
		Prompt:            "> ",
		HistoryFile:       "/tmp/todoer.tmp",
		InterruptPrompt:   "^C",
		EOFPrompt:         "exit",
		HistorySearchFold: true,
		Stdin:             r,
	})

	// catch output
	r1, w1, err := os.Pipe()
	if err != nil {
		fmt.Println(err)
	}

	todoer.Run(w1)
	w1.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r1)
	output := string(buf.Bytes())
	//fmt.Println("output")
	//fmt.Println(output)
	expected := "[item1 item2 item3]\n[item1 item2 item3 item4]\nitem1\nitem2\nitem3\nitem4\nunknown command\n"

	assert.Equal(t, expected, output)

}
