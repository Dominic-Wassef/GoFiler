package main

import (
	"fmt"
	"os"
)

// FileOpCreate creates a new file.
func FileOpCreate(name string) error {
	file, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	fmt.Printf("File %s created\n", name)
	return nil
}

// FileOpDelete deletes a file.
func FileOpDelete(name string) error {
	err := os.Remove(name)
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	fmt.Printf("File %s deleted\n", name)
	return nil
}

// FileOpRename renames a file.
func FileOpRename(oldName, newName string) error {
	err := os.Rename(oldName, newName)
	if err != nil {
		return fmt.Errorf("failed to rename file: %w", err)
	}

	fmt.Printf("File %s renamed to %s\n", oldName, newName)
	return nil
}

// FileOpMove moves a file from source to destination.
func FileOpMove(src, dest string) error {
	err := os.Rename(src, dest)
	if err != nil {
		return fmt.Errorf("failed to move file: %w", err)
	}

	fmt.Printf("File %s moved to %s\n", src, dest)
	return nil
}

// FileOpRead reads the content of a file and returns it as a string.
func FileOpRead(name string) (string, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return string(data), nil
}

// FileOpWrite writes a string to a file.
// If the file does not exist, FileOpWrite creates it.
// Otherwise, FileOpWrite truncates it before writing.
func FileOpWrite(name, content string) error {
	err := os.WriteFile(name, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// FileOpAppend appends a string to a file.
// If the file does not exist, FileOpAppend creates it.
func FileOpAppend(name, content string) error {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
