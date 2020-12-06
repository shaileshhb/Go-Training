package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// var directory []folderStructure
	rootFolder := "E:\\Go\\src\\day4"
	directory, err := readCurrentDir(rootFolder)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(directory)
	for _, folders := range directory {
		fmt.Println(folders.folderName)
		for _, files := range folders.fileStructure.files {
			fmt.Println("\t", files)
		}
		if folders.folderName != "Root" {
			for _, subFolder := range *folders.folderPtr {
				fmt.Println("\t", subFolder.folderName)
				for _, subFiles := range subFolder.fileStructure.files {
					fmt.Println("\t\t", subFiles)
				}
			}
		}
	}

}

func readCurrentDir(folderPath string) ([]folderStructure, error) {

	var dir []folderStructure
	// var rootDirectoryFiles []string

	folders, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	for _, folder := range folders {
		files, err := os.Open(folderPath + "\\" + folder.Name())
		if err != nil {
			return nil, err
		}
		info, err := files.Stat()
		if err != nil {
			return nil, err
		}
		if info.IsDir() {
			// folder.Name() -> current directory folders
			folder, err := addFolder(folderPath, folder.Name())
			if err != nil {
				return nil, err
			}

			dir = append(dir, folder)
		}

	}

	// consists of root directory files
	folder, err := addFolder(folderPath, "Root")
	if err != nil {
		return nil, err
	}
	dir = append(dir, folder)

	return dir, nil

}

func getFolders(folderPath string) (*[]folderStructure, []string, error) {

	var foldersPtr []folderStructure
	// var currentFoldersPtr []string
	var currentFoldersFiles []string

	folders, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return nil, nil, err
	}

	for _, folder := range folders {
		files, err := os.Open(folderPath + "\\" + folder.Name())
		if err != nil {
			return nil, nil, err
		}
		info, err := files.Stat()
		if err != nil {
			return nil, nil, err
		}
		if info.IsDir() {
			// folder.Name() -> current directory folders

			subFolderPtr, subFolderFiles, err := getFolders(folderPath + "\\" + folder.Name())
			if err != nil {
				return nil, nil, err
			}
			// sets the folder ptr and its files
			currentFolderStructure, err := setSubFolder(folder.Name(), *subFolderPtr, subFolderFiles)
			if err != nil {
				return nil, nil, err
			}

			foldersPtr = append(foldersPtr, currentFolderStructure)

		} else {
			currentFoldersFiles = append(currentFoldersFiles, folder.Name())
		}
	}

	return &foldersPtr, currentFoldersFiles, nil

}

func setSubFolder(folderName string, folderPtr []folderStructure, currentFolderFiles []string) (folderStructure, error) {

	var folder = &folderStructure{
		folderName: folderName,
		folderPtr:  &folderPtr,
		fileStructure: fileStructure{
			currentFolderFiles,
		},
	}

	return *folder, nil
}

func addFolder(folderPath string, folderName string) (folderStructure, error) {

	var foldersPtr *[]folderStructure
	var filesName []string
	var err error

	if folderName != "Root" {
		foldersPtr, filesName, err = getFolders(folderPath + "\\" + folderName)
	} else {
		foldersPtr, filesName, err = getFolders(folderPath)
	}

	var folder = &folderStructure{
		folderName: folderName,
		folderPtr:  foldersPtr,
		fileStructure: fileStructure{
			filesName,
		},
	}

	if err != nil {
		return *folder, err
	}

	return *folder, nil
}

type fileStructure struct {
	files []string
}

type folderStructure struct {
	folderName string
	folderPtr  *[]folderStructure
	fileStructure
}
