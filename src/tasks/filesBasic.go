package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	dir, err := readDir("E:\\Go\\src\\day4")

	if err != nil {
		log.Fatal(err)
	}

	for folders, files := range dir {
		if folders != "CurrentFolder" {
			fmt.Println(folders)
		}
		for _, file := range files.files {
			if folders != "CurrentFolder" {
				fmt.Println("\t", file)
			} else {
				fmt.Println(file)
			}
		}
	}
}

// readDir will read the dir and return all files and folders as map[string]folder
func readDir(folderName string) (map[string]folders, error) {

	var dir = make(map[string]folders)
	files, err := ioutil.ReadDir(folderName)
	var currentFolders folders

	if err != nil {
		return nil, err
	}

	for _, file := range files {
		var folder folders
		f, err := os.Open(folderName + "\\" + file.Name())
		if err != nil {
			return nil, err
		}
		info, err := f.Stat()
		if err != nil {
			return nil, err
		}
		if info.IsDir() {
			currentFile, err := ioutil.ReadDir(folderName + "\\" + file.Name())
			if err != nil {
				return nil, err
			}
			for _, fileName := range currentFile {
				fmt.Println("files:", fileName.Name())
				folder.files = append(folder.files, fileName.Name())
			}
			dir[file.Name()] = folder
		} else {
			currentFolders.files = append(currentFolders.files, file.Name())
		}
	}

	dir["CurrentFolder"] = currentFolders

	return dir, nil
}

// it will contain file of particular folder
type folders struct {
	files []string
}
