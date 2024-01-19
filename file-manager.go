package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// Function to get the name of the current folder
func getCurrentFolderName() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Base(cwd), nil
}

// Function to replace old word with new words in a file, accepts a map for replacements
func ReplaceInFile(path string, replacements map[string]string, newPath string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	newFile, err := os.Create(newPath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(newFile)
	for scanner.Scan() {
		line := scanner.Text()
		for oldWord, newWord := range replacements {
			line = strings.ReplaceAll(line, oldWord, newWord)
		}
		if _, err = writer.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	if scanner.Err() != nil {
		return scanner.Err()
	}
	if err := writer.Flush(); err != nil {
		return err
	}
	return nil
}

func checkFileExist(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return err
	}
	return nil
}
