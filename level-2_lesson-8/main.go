package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"sync"
	"time"
)

type File struct {
	name string
	size int64
	sha  string
	path string
	date time.Time
}

type SortedByDate []File

func (ff SortedByDate) Len() int {
	return len(ff)
}
func (ff SortedByDate) Less(i, j int) bool {
	return ff[i].date.Before(ff[j].date)
}
func (ff SortedByDate) Swap(i, j int) {
	ff[i], ff[j] = ff[j], ff[i]
}

type Duplicates struct {
	sync.Mutex
	mm map[string][]File
}

func NewDuplicates() *Duplicates {
	return &Duplicates{
		mm: map[string][]File{},
	}
}

func (d *Duplicates) Add(f File) {
	d.Lock()
	key := f.name + ":" + strconv.Itoa(int(f.size)) + ":" + f.sha
	dFiles, ok := d.mm[key]
	if !ok {
		dFiles = []File{}
	}
	d.mm[key] = append(dFiles, f)
	d.Unlock()
}

var (
	path  *string
	force *bool
	dir   string
	err   error
	wg    sync.WaitGroup
)

func init() {
	path = flag.String("path", ".", "Path where to search file duplicates")
	force = flag.Bool("force", false, "Delete file duplicates")
	flag.Parse()
}

func main() {
	if *path == "." {
		dir, err = os.Getwd()
		if err != nil {
			panic(err)
		}
	} else {
		dir = *path
	}
	dups := NewDuplicates()

	wg.Add(1)
	go exploreDir(dir, dups)
	wg.Wait()

	// Show and delete
	no_duplicates := true
	for _, v := range dups.mm {
		if len(v) > 1 {
			no_duplicates = false
			sort.Sort(SortedByDate(v))
			fmt.Printf("For file %s we found duplicates:\n", v[0].name)
			for _, f := range v {
				fmt.Printf("\t%s, last modified: %v\n", f.path, f.date)
			}
			if *force {
				for _, f := range v[1:] {
					err = os.Remove(f.path)
					if err != nil {
						panic(err)
					}
					fmt.Printf("\tFile %s was deleted\n", f.path)
				}
			}
		}
	}
	if no_duplicates {
		fmt.Println("No duplicates were found!")
	}
	fmt.Println("End of program!")
}

func exploreDir(path string, dups *Duplicates) {
	defer wg.Done()
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		fullPath := filepath.Join(path, f.Name())
		if f.Mode().IsDir() {
			wg.Add(1)
			go exploreDir(fullPath, dups)
		} else {
			hasher := sha256.New()
			s, err := ioutil.ReadFile(filepath.Join(path, f.Name()))
			if err != nil {
				panic(err)
			}
			hasher.Write(s)
			sha := hex.EncodeToString(hasher.Sum(nil))
			dups.Add(File{f.Name(), f.Size(), sha, fullPath, f.ModTime()})
		}
	}
}
