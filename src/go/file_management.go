package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// CreateFile creates a new file.
func CreateFile(name string) error {
	file, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	fmt.Printf("File %s created\n", name)
	return nil
}

// DeleteFile deletes a file.
func DeleteFile(name string) error {
	err := os.Remove(name)
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	fmt.Printf("File %s deleted\n", name)
	return nil
}

// RenameFile renames a file.
func RenameFile(oldName, newName string) error {
	err := os.Rename(oldName, newName)
	if err != nil {
		return fmt.Errorf("failed to rename file: %w", err)
	}

	fmt.Printf("File %s renamed to %s\n", oldName, newName)
	return nil
}

// MoveFile moves a file from source to destination.
func MoveFile(src, dest string) error {
	err := os.Rename(src, dest)
	if err != nil {
		return fmt.Errorf("failed to move file: %w", err)
	}

	fmt.Printf("File %s moved to %s\n", src, dest)
	return nil
}

// ListFiles lists all files in a directory.
func ListFiles(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	return nil
}

// GetPermissions returns the permissions of a file.
func GetPermissions(name string) error {
	info, err := os.Stat(name)
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	fmt.Printf("Permissions of file %s: %s\n", name, info.Mode())
	return nil
}

// SetPermissions sets the permissions of a file.
func SetPermissions(name string, mode os.FileMode) error {
	err := os.Chmod(name, mode)
	if err != nil {
		return fmt.Errorf("failed to set file permissions: %w", err)
	}

	fmt.Printf("Set permissions of file %s to %s\n", name, mode)
	return nil
}

// ReadFile reads the content of a file.
func ReadFile(name string) error {
	data, err := os.ReadFile(name)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	fmt.Printf("Content of file %s:\n%s\n", name, data)
	return nil
}

// WriteFile writes data to a file.
func WriteFile(name string, data []byte) error {
	err := os.WriteFile(name, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	fmt.Printf("Wrote data to file %s\n", name)
	return nil
}

// CreateDirectory creates a new directory.
func CreateDirectory(name string) error {
	err := os.MkdirAll(name, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	fmt.Printf("Directory %s created\n", name)
	return nil
}

// DeleteDirectory deletes a directory.
func DeleteDirectory(name string) error {
	err := os.RemoveAll(name)
	if err != nil {
		return fmt.Errorf("failed to delete directory: %w", err)
	}

	fmt.Printf("Directory %s deleted\n", name)
	return nil
}

// RenameDirectory renames a directory.
func RenameDirectory(oldName, newName string) error {
	err := os.Rename(oldName, newName)
	if err != nil {
		return fmt.Errorf("failed to rename directory: %w", err)
	}

	fmt.Printf("Directory %s renamed to %s\n", oldName, newName)
	return nil
}

// MoveDirectory moves a directory from source to destination.
func MoveDirectory(src, dest string) error {
	err := os.Rename(src, dest)
	if err != nil {
		return fmt.Errorf("failed to move directory: %w", err)
	}

	fmt.Printf("Directory %s moved to %s\n", src, dest)
	return nil
}

// GetDirectorySize returns the size of a directory.
func GetDirectorySize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})

	if err != nil {
		return 0, fmt.Errorf("failed to calculate directory size: %w", err)
	}

	return size, nil
}