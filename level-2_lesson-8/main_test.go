package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExploreDir(t *testing.T) {
	dir, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	// Create files for test
	f1, err := os.Create(filepath.Join(dir, "file1"))
	if err != nil {
		panic(err)
	}
	_, err = f1.WriteString("file1 string 1")
	if err != nil {
		panic(err)
	}
	f1.Close()

	f2, err := os.Create(filepath.Join(dir, "file2"))
	if err != nil {
		panic(err)
	}
	_, err = f2.WriteString("file2 string 2")
	if err != nil {
		panic(err)
	}
	f2.Close()

	err = os.Mkdir(filepath.Join(dir, "temp"), os.ModePerm)
	if err != nil {
		panic(err)
	}
	// duplicate
	f3, err := os.Create(filepath.Join(dir, "temp", "file2"))
	if err != nil {
		panic(err)
	}
	_, err = f3.WriteString("file2 string 2")
	if err != nil {
		panic(err)
	}
	f3.Close()

	// not duplicate
	f4, err := os.Create(filepath.Join(dir, "temp", "file4"))
	if err != nil {
		panic(err)
	}
	_, err = f4.WriteString("file4 string 4")
	if err != nil {
		panic(err)
	}
	f4.Close()

	dups := NewDuplicates()

	wg.Add(1)
	go exploreDir(dir, dups)
	wg.Wait()

	// remove all
	err = os.Remove(filepath.Join(dir, "file1"))
	if err != nil {
		panic(err)
	}
	err = os.Remove(filepath.Join(dir, "file2"))
	if err != nil {
		panic(err)
	}
	err = os.Remove(filepath.Join(dir, "temp", "file2"))
	if err != nil {
		panic(err)
	}
	err = os.Remove(filepath.Join(dir, "temp", "file4"))
	if err != nil {
		panic(err)
	}
	err = os.Remove(filepath.Join(dir, "temp"))
	if err != nil {
		panic(err)
	}

	//Check results
	no_duplicates := true
	dupFileName := ""
	for _, v := range dups.mm {
		if len(v) > 1 {
			no_duplicates = false
			dupFileName = v[0].name
		}
	}
	if no_duplicates {
		t.Error("Fail: expected duplicates, got none")
	}
	if dupFileName != "file2" {
		t.Error("Fail: expected file name duplicate file2, got", dupFileName)
	}

}
