package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// BackupDir is the directory where backups will be stored.
const BackupDir = "backups"

// BackupSuffix is the suffix that will be added to the file names to create backups.
const BackupSuffix = "_backup"

// backupFile creates a backup copy of the file with the given path.
func BackupFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open the file for backup: %w", err)
	}
	defer file.Close()

	// Ensure the backup directory exists.
	err = os.MkdirAll(BackupDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create backup directory: %w", err)
	}

	// Append a timestamp to the backup filename to avoid overwriting previous backups.
	timestamp := time.Now().Format("20060102150405")
	backupPath := filepath.Join(BackupDir, filepath.Base(path)+BackupSuffix+"_"+timestamp)
	backupFile, err := os.Create(backupPath)
	if err != nil {
		return fmt.Errorf("failed to create backup file: %w", err)
	}
	defer backupFile.Close()

	_, err = io.Copy(backupFile, file)
	if err != nil {
		return fmt.Errorf("failed to write to backup file: %w", err)
	}

	fmt.Printf("Backup of %s created at %s\n", path, backupPath)
	return nil
}

// restoreBackup restores the latest backup of the file with the given path, if it exists.
func RestoreBackup(path string) error {
	// Find the latest backup file.
	files, err := os.ReadDir(BackupDir)
	if err != nil {
		return fmt.Errorf("failed to read backup directory: %w", err)
	}

	var backupPath string
	for _, file := range files {
		if strings.HasPrefix(file.Name(), filepath.Base(path)+BackupSuffix) {
			backupPath = filepath.Join(BackupDir, file.Name())
		}
	}

	if backupPath == "" {
		fmt.Println("No backup found for the given file.")
		return nil
	}

	backupFile, err := os.Open(backupPath)
	if err != nil {
		return fmt.Errorf("failed to open the backup file: %w", err)
	}
	defer backupFile.Close()

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create the original file from backup: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, backupFile)
	if err != nil {
		return fmt.Errorf("failed to restore the file from backup: %w", err)
	}

	fmt.Printf("File %s restored from backup at %s\n", path, backupPath)
	return nil
}

// listBackups prints a list of all existing backups for the file with the given path.
func ListBackups(path string) error {
	files, err := os.ReadDir(BackupDir)
	if err != nil {
		return fmt.Errorf("failed to read backup directory: %w", err)
	}

	fmt.Printf("Existing backups for %s:\n", path)
	for _, file := range files {
		if strings.HasPrefix(file.Name(), filepath.Base(path)+BackupSuffix) {
			fmt.Println(file.Name())
		}
	}

	return nil
}

// checkFileIntegrity checks the integrity of the backup by comparing
// the checksum of the original file and the backup.
func CheckFileIntegrity(originalPath, backupPath string) error {
	originalChecksum, err := CalculateChecksum(originalPath)
	if err != nil {
			return fmt.Errorf("failed to calculate checksum for the original file: %w", err)
	}

	backupChecksum, err := CalculateChecksum(backupPath)
	if err != nil {
			return fmt.Errorf("failed to calculate checksum for the backup file: %w", err)
	}

	if originalChecksum != backupChecksum {
			return fmt.Errorf("checksums don't match; the backup might be corrupted")
	}

	return nil
}

// calculateChecksum calculates a SHA-256 checksum for the file.
func CalculateChecksum(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
			return "", fmt.Errorf("failed to open the file: %w", err)
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
			return "", fmt.Errorf("failed to calculate checksum: %w", err)
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
