package main

import (
	"fmt"
	"os"
	"sync"
)

// User struct represents a user in the system.
type User struct {
	Name string
	Role string
}

// File represents a file in the system.
type File struct {
	Name       string
	Content    string
	Versions   map[int]string // versions of the file
	mutex      sync.Mutex
	changes    map[string]string // maps usernames to changes
	Permission map[string]string // maps usernames to their role
}

// NewFile creates a new File.
func NewFile(name string) *File {
	return &File{
		Name:       name,
		Content:    "",
		Versions:   make(map[int]string),
		changes:    make(map[string]string),
		Permission: make(map[string]string),
	}
}

// Edit simulates editing a file by appending a string to its content.
func (f *File) Edit(user User, change string) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	// Check user permission
	if f.Permission[user.Name] != "editor" && f.Permission[user.Name] != "admin" {
		return fmt.Errorf("user does not have permission to edit the file")
	}

	version := len(f.Versions) + 1
	f.Content += change
	f.changes[user.Name] = change
	f.Versions[version] = f.Content

	fmt.Printf("%s added: %q\n", user.Name, change)
	return nil
}

// Save writes the file's content to disk.
func (f *File) Save() error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	err := os.WriteFile(f.Name, []byte(f.Content), 0644)
	if err != nil {
		return fmt.Errorf("failed to save the file: %w", err)
	}

	fmt.Printf("File %s saved\n", f.Name)
	return nil
}

// LoadVersion loads a specific version of the file's content.
func (f *File) LoadVersion(version int) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if _, ok := f.Versions[version]; !ok {
		return fmt.Errorf("version %d does not exist", version)
	}

	f.Content = f.Versions[version]

	fmt.Printf("Loaded version %d of file %s\n", version, f.Name)
	return nil
}

// PrintChanges prints the changes made to the file.
func (f *File) PrintChanges() {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	fmt.Printf("Changes to %s:\n", f.Name)
	for user, change := range f.changes {
		fmt.Printf("%s: %q\n", user, change)
	}
}

// AssignRole assigns a role to a user for this file.
func (f *File) AssignRole(user User, role string) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	f.Permission[user.Name] = role

	fmt.Printf("Role %s assigned to user %s for file %s\n", role, user.Name, f.Name)
}
