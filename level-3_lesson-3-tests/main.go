package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/chzyer/readline"
	"io"
	"log"
	"os"
	"strings"
)

type DB interface {
	AddItem(i string)
	List() []string
}

type DBImpl struct {
	Storage *os.File
	cache   []string
}

func NewDBImpl() (*DBImpl, error) {
	db, err := os.OpenFile("todoer.db", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("cannot open db: %v", err)
		return nil, err
	}
	var items []string
	scanner := bufio.NewScanner(db)
	for scanner.Scan() {
		items = append(items, scanner.Text())
	}

	return &DBImpl{db, items}, nil
}

func (db *DBImpl) AddItem(item string) {
	_, err := db.Storage.WriteString(item + "\n")
	if err != nil {
		log.Fatalf("cannot write string: %v", err)
	}
	db.cache = append(db.cache, item)
}

func (db *DBImpl) List() []string {
	return db.cache
}

func (db *DBImpl) Close() {
	err := db.Storage.Close()
	if err != nil {
		log.Fatalf("cannot close db %v", err)
	}
}

type TodoerService interface {
	Run()
}

type TodoerServiceImpl struct {
	db      DB
	lineRdr *readline.Instance
}

func NewTodoerServiceImpl(db DB, settings *readline.Config) (*TodoerServiceImpl, error) {
	lineRdr, err := readline.NewEx(settings)
	if err != nil {
		log.Fatalf("todoer: create line reader")
		return nil, err
	}
	return &TodoerServiceImpl{db, lineRdr}, nil
}

type myBuffer struct {
	bytes.Buffer
}

func (b myBuffer) Close() error {
	return nil
}

type myBufio struct {
	bufio.Reader
}

func (b myBufio) Close() error {
	return nil
}
func NewReader(rd io.Reader) *bufio.Reader {
	return NewReaderSize(rd, bufio.defaultBufSize)
}

func (todoer *TodoerServiceImpl) Run() {
	for {
		str, err := todoer.lineRdr.Readline()
		if err != nil {
			if err != readline.ErrInterrupt && err != io.EOF {
				log.Fatalf("read line: %v", err)
			}
			break
		}

		if str == "" {
			continue
		}
		tokens := strings.Split(str, " ")
		switch tokens[0] {
		case "add":
			if len(tokens) == 1 {
				continue
			}
			item := strings.Join(tokens[1:], " ")
			todoer.db.AddItem(item)
			fmt.Println(todoer.db.List())

		case "list":
			fmt.Println(strings.Join(todoer.db.List(), "\n"))

		default:
			fmt.Println("unknown command")
		}
	}
}

func main() {
	db, err := NewDBImpl()
	if err != nil {
		log.Fatalf("cannot open db: %v", err)
	}
	defer db.Close()

	//var b myBuffer
	var b bytes.Buffer
	b.Write([]byte("add asd\nlist\n"))

	todoer, err := NewTodoerServiceImpl(db, &readline.Config{
		Prompt:            "> ",
		HistoryFile:       "/tmp/todoer.tmp",
		InterruptPrompt:   "^C",
		EOFPrompt:         "exit",
		HistorySearchFold: true,
		Stdin:             bufio.NewReader(&b),
	})

	//fmt.Println(todoer)
	todoer.Run()

}
