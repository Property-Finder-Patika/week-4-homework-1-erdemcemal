package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	listSourceCodeFiles(".")
}

// listSourceCodeFiles lists all the source code files in the given directory.
func listSourceCodeFiles(root string) {
	list := make([]string, 0) // create a list of strings to store the file names
	// list all the files in the given directory that end with .go
	walkDirFunc := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if d.Name()[len(d.Name())-3:] == ".go" {
			list = append(list, path)
		}
		return nil
	}
	// walk the given directory
	result := filepath.WalkDir(root, walkDirFunc)
	if result != nil {
		panic(result)
	}
	// print the list of files
	for _, path := range list {
		fmt.Println(path)
	}
}
