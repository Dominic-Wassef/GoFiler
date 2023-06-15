package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func handleError(err error) {
	if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
	}
}

func validateInput(input []string, expectedLen int, message string) {
	if len(input) != expectedLen {
		fmt.Println(message)
		os.Exit(1)
	}
}

func help() {
	fmt.Println(`
Usage: 
  -create='filename'                  : Create a new file with specified name
  -read='filename'                    : Read a file with specified name
  -write='filename' -data='data'      : Write to a file
  -append='filename' -data='data'     : Append to a file
  -delete='filename'                  : Delete a file with specified name
  -rename='oldname,newname'           : Rename a file
  -move='src,dest'                    : Move a file
  -backup='filename'                  : Backup a file with specified name
  -restore='filename'                 : Restore a file from backup with specified name
  -listBackups='path'                 : List backups for a file with specified path
  -checkintegrity='filename,backupfilename' : Check file integrity
  -user='username,role'               : Specify a user
  -edit='filename' -data='data' -user='username,role' : Edit a file
  -save='filename'                    : Save a file with specified name
  -loadVersion='filename,version'     : Load a specific version of a file
  -printChanges='filename'            : Print changes of a file with specified name
  -assignRole='filename' -user='username,role' : Assign a role to a user for a file
  -compressEncrypt='filename'         : Compress and encrypt a file
  -decompressDecrypt='filename'       : Decompress and decrypt a file
  -listFiles='directory'              : List all files in a directory
  -getPermissions='filename'          : Get permissions of a file
  -setPermissions='filename,mode'     : Set permissions of a file
  -info='filename'                    : Get metadata of a file or directory`)
}


func main() {
	createPtr := flag.String("create", "", "Create a new file with specified name")
	readPtr := flag.String("read", "", "Read a file with specified name")
	writePtr := flag.String("write", "", "Write to a file. Use in the format -write='filename' -data='data to write'")
	appendPtr := flag.String("append", "", "Append to a file. Use in the format -append='filename' -data='data to append'")
	deletePtr := flag.String("delete", "", "Delete a file with specified name")
	renamePtr := flag.String("rename", "", "Rename a file. Use in the format -rename='oldname,newname'")
	movePtr := flag.String("move", "", "Move a file. Use in the format -move='src,dest'")
	dataPtr := flag.String("data", "", "Data to write or append to file")
	backupPtr := flag.String("backup", "", "Backup a file with specified name")
	restorePtr := flag.String("restore", "", "Restore a file from backup with specified name")
	listBackupsPtr := flag.String("listBackups", "", "List backups for a file with specified path")
	checkIntegrityPtr := flag.String("checkintegrity", "", "Check file integrity. Use in the format -checkintegrity='filename,backupfilename'")
	userPtr := flag.String("user", "", "Specify a user. Use in the format -user='username,role'")
	editPtr := flag.String("edit", "", "Edit a file. Use in the format -edit='filename' -data='data to append' -user='username,role'")
	savePtr := flag.String("save", "", "Save a file with specified name")
	loadVersionPtr := flag.String("loadVersion", "", "Load a specific version of a file. Use in the format -loadVersion='filename,version'")
	printChangesPtr := flag.String("printChanges", "", "Print changes of a file with specified name")
	assignRolePtr := flag.String("assignRole", "", "Assign a role to a user for a file. Use in the format -assignRole='filename' -user='username,role'")
	compressEncryptPtr := flag.String("compressEncrypt", "", "Compress and encrypt a file. Use in the format -compressEncrypt='filename'")
	decompressDecryptPtr := flag.String("decompressDecrypt", "", "Decompress and decrypt a file. Use in the format -decompressDecrypt='filename'")
	createFilePtr := flag.String("createFile", "", "Create a file. Use in the format -createFile='filename'")
	deleteFilePtr := flag.String("deleteFile", "", "Delete a file. Use in the format -deleteFile='filename'")
	renameFilePtr := flag.String("renameFile", "", "Rename a file. Use in the format -renameFile='oldname,newname'")
	moveFilePtr := flag.String("moveFile", "", "Move a file. Use in the format -moveFile='src,dest'")
	listFilesPtr := flag.String("listFiles", "", "List all files in a directory. Use in the format -listFiles='directory'")
	getPermissionsPtr := flag.String("getPermissions", "", "Get permissions of a file. Use in the format -getPermissions='filename'")
	setPermissionsPtr := flag.String("setPermissions", "", "Set permissions of a file. Use in the format -setPermissions='filename,mode'")
	infoPtr := flag.String("info", "", "Get metadata of a file or directory")
	
	// Set the custom usage function before parsing the flags
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Unknown command or flag.\n")
		help()
	}

	// Parse the flags
	flag.Parse()

	if *createPtr != "" {
    handleError(FileOpCreate(*createPtr))
	}

	if *readPtr != "" {
		data, err := FileOpRead(*readPtr)
		handleError(err)
		fmt.Println(data)
	}

	if *writePtr != "" {
		err := FileOpWrite(*writePtr, *dataPtr)
		handleError(err)
	}
	
	if *appendPtr != "" {
		err := FileOpAppend(*appendPtr, *dataPtr)
		handleError(err)
	}
	
	if *deletePtr != "" {
		err := FileOpDelete(*deletePtr)
		handleError(err)
	}
	
	if *renamePtr != "" {
		names := strings.Split(*renamePtr, ",")
		if len(names) != 2 {
			fmt.Println("Invalid rename format. Use -rename='oldname,newname'")
			os.Exit(1)
		}
		err := FileOpRename(names[0], names[1])
		handleError(err)
	}
	
	if *movePtr != "" {
		paths := strings.Split(*movePtr, ",")
		if len(paths) != 2 {
			fmt.Println("Invalid move format. Use -move='src,dest'")
			os.Exit(1)
		}
		err := FileOpMove(paths[0], paths[1])
		handleError(err)
	}	

	if *backupPtr != "" {
		err := BackupFile(*backupPtr)
		handleError(err)
	}
	
	if *listBackupsPtr != "" {
		err := ListBackups(*listBackupsPtr)
		handleError(err)
	}
	
	if *checkIntegrityPtr != "" {
		checksum, err := CalculateChecksum(*checkIntegrityPtr)
		handleError(err)
		fmt.Println("Checksum:", checksum)
	}
	
	if *restorePtr != "" {
		err := RestoreBackup(*restorePtr)
		handleError(err)
	}	

	if *checkIntegrityPtr != "" {
		paths := strings.Split(*checkIntegrityPtr, ",")
		validateInput(paths, 2, "Invalid integrity check format. Use -checkintegrity='filename,backupfilename'")
		err := CheckFileIntegrity(paths[0], paths[1])
		handleError(err)
	}
	
	var user *User
	if *userPtr != "" {
		userDetails := strings.Split(*userPtr, ",")
		validateInput(userDetails, 2, "Invalid user format. Use -user='username,role'")
		user = &User{Name: userDetails[0], Role: userDetails[1]}
	}
	
	if *editPtr != "" {
		file := NewFile(*editPtr)
		err := file.Edit(*user, *dataPtr)
		handleError(err)
	}
	
	if *savePtr != "" {
		file := NewFile(*savePtr)
		err := file.Save()
		handleError(err)
	}
	
	if *loadVersionPtr != "" {
		fileVersionDetails := strings.Split(*loadVersionPtr, ",")
		validateInput(fileVersionDetails, 2, "Invalid loadVersion format. Use -loadVersion='filename,version'")
		file := NewFile(fileVersionDetails[0])
		version, err := strconv.Atoi(fileVersionDetails[1])
		handleError(err)
		err = file.LoadVersion(version)
		handleError(err)
	}
	

	var file *File
	if *printChangesPtr != "" {
		file = NewFile(*printChangesPtr)
		file.PrintChanges()
	}

	if *assignRolePtr != "" && user != nil {
		file = NewFile(*assignRolePtr)
		file.AssignRole(*user, user.Role)
	}

	if *compressEncryptPtr != "" {
		err := CompressAndEncryptFile(*compressEncryptPtr)
		handleError(err)
	}

	if *decompressDecryptPtr != "" {
		err := DecryptAndDecompressFile(*decompressDecryptPtr)
		handleError(err)
	}

	if *createFilePtr != "" {
		err := CreateFile(*createFilePtr)
		handleError(err)
	}

	if *deleteFilePtr != "" {
		err := DeleteFile(*deleteFilePtr)
		handleError(err)
	}

	if *renameFilePtr != "" {
		files := strings.Split(*renameFilePtr, ",")
		validateInput(files, 2, "Invalid format. Use -renameFile='oldname,newname'")
		err := RenameFile(files[0], files[1])
		handleError(err)
	}

	if *moveFilePtr != "" {
		files := strings.Split(*moveFilePtr, ",")
		validateInput(files, 2, "Invalid format. Use -moveFile='src,dest'")
		err := MoveFile(files[0], files[1])
		handleError(err)
	}

	if *listFilesPtr != "" {
		err := ListFiles(*listFilesPtr)
		handleError(err)
	}

	if *getPermissionsPtr != "" {
		err := GetPermissions(*getPermissionsPtr)
		handleError(err)
	}

	if *setPermissionsPtr != "" {
		args := strings.Split(*setPermissionsPtr, ",")
		validateInput(args, 2, "Invalid format. Use -setPermissions='filename,mode'")
		mode, err := strconv.ParseUint(args[1], 8, 32)
		handleError(err)
		err = SetPermissions(args[0], os.FileMode(mode))
		handleError(err)
	}

	if *infoPtr != "" {
		fileInfo, err := GetFileInfo(*infoPtr)
		handleError(err)
		PrintFileInfo(fileInfo)
	}
}