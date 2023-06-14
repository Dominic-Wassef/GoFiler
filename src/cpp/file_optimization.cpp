#include <iostream>
#include <cstdlib>
#include "file_optimization.h"

// Check file system integrity
void FileOptimizer::CheckFileSystemIntegrity(const std::string& fs) {
    std::string command = "diskutil verifyVolume " + fs;
    int result = std::system(command.c_str());
    if (result != 0) {
        std::cerr << "FileSystem integrity check failed with error: " << result << '\n';
        RepairFileSystem(fs);
    }
}

// Repair the file system
void FileOptimizer::RepairFileSystem(const std::string& fs) {
    std::string command = "diskutil repairVolume " + fs;
    int result = std::system(command.c_str());
    if (result != 0) {
        std::cerr << "Repairing file system failed with error: " << result << '\n';
    }
}

// Manage permissions
void FileOptimizer::ManagePermissions(const std::string& path) {
    std::string command = "chmod -R 777 " + path;
    int result = std::system(command.c_str());
    if (result != 0) {
        std::cerr << "Changing permissions failed with error: " << result << '\n';
    }
}

// Check disk space
void FileOptimizer::CheckDiskSpace(const std::string& path) {
    std::filesystem::space_info s = std::filesystem::space(path);
    std::cout << "Free space: " << s.free / (1024 * 1024) << "MB\n";
}

// Method to run the file optimizer
void FileOptimizer::RunFileOptimizer(const std::string& path) {
    CheckFileSystemIntegrity(path);
    ManagePermissions(path);
    CheckDiskSpace(path);
}
