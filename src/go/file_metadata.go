package main

import (
	"fmt"
	"os"
	"time"
)

// FileInfo is a structure that holds metadata about a file.
type FileInfo struct {
	Name         string
	Size         int64
	Permissions  os.FileMode
	ModTime      time.Time
	IsDir        bool
}

// GetFileInfo gets the metadata of a file or directory.
func GetFileInfo(name string) (*FileInfo, error) {
	info, err := os.Stat(name)
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}

	fileInfo := &FileInfo{
		Name:         info.Name(),
		Size:         info.Size(),
		Permissions:  info.Mode(),
		ModTime:      info.ModTime(),
		IsDir:        info.IsDir(),
	}

	return fileInfo, nil
}

// PrintFileInfo prints the metadata of a file or directory.
func PrintFileInfo(fileInfo *FileInfo) {
	fmt.Printf("Name: %s\n", fileInfo.Name)
	fmt.Printf("Size: %d\n", fileInfo.Size)
	fmt.Printf("Permissions: %s\n", fileInfo.Permissions)
	fmt.Printf("ModTime: %s\n", fileInfo.ModTime)
	fmt.Printf("IsDir: %t\n", fileInfo.IsDir)
}

// SetFilePermissions sets the permissions of a file or directory.
func SetFilePermissions(name string, mode os.FileMode) error {
	err := os.Chmod(name, mode)
	if err != nil {
		return fmt.Errorf("failed to set file permissions: %w", err)
	}

	return nil
}

// SetFileOwner sets the owner and group of a file or directory.
func SetFileOwner(name string, uid, gid int) error {
	err := os.Chown(name, uid, gid)
	if err != nil {
		return fmt.Errorf("failed to set file owner: %w", err)
	}

	return nil
}
