package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	path  *string
	force *bool
	dir   string
	err   error
)

func init() {
	path = flag.String("path", ".", "Path where to search file duplicates")
	force = flag.Bool("force", false, "Delete file duplicates")
	flag.Parse()
}

func main() {
	//var path = flag.String("path", ".", "Path where to search file duplicates")
	fmt.Println(*path)
	fmt.Println(*force)
	if *path == "." {
		dir, err = os.Getwd()
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(dir)
}
