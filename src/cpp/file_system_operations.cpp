#include <iostream>
#include <fstream>
#include <filesystem>
#include <string>
#include <vector>
#include "file_optimization.h" 

class FileSystemOperations {
public:
    // Create a new file
    void CreateFile(const std::string& filename) {
        std::ofstream file(filename);
        if (!file) {
            std::cerr << "Error in creating file: " << filename << '\n';
        }
        file.close();
    }

    // Delete a file
    void DeleteFile(const std::string& filename) {
        if (remove(filename.c_str()) != 0) {
            std::cerr << "Error in deleting file: " << filename << '\n';
        }
    }

    // Rename a file
    void RenameFile(const std::string& oldName, const std::string& newName) {
        if (std::rename(oldName.c_str(), newName.c_str()) != 0) {
            std::cerr << "Error in renaming file: " << oldName << '\n';
        }
    }

    // Move a file
    void MoveFile(const std::string& source, const std::string& destination) {
        if (std::rename(source.c_str(), destination.c_str()) != 0) {
            std::cerr << "Error in moving file: " << source << '\n';
        }
    }

    // List all files in a directory
    void ListFiles(const std::string& path) {
        for (const auto& entry : std::filesystem::directory_iterator(path)) {
            std::cout << entry.path() << '\n';
        }
    }

    // Write data to a file
    void WriteToFile(const std::string& filename, const std::string& data) {
        std::ofstream file(filename, std::ios::app); // open in append mode
        if (!file) {
            std::cerr << "Error opening file: " << filename << '\n';
            return;
        }
        file << data;
        file.close();
    }

    // Read data from a file
    std::vector<std::string> ReadFromFile(const std::string& filename) {
        std::ifstream file(filename);
        if (!file) {
            std::cerr << "Error opening file: " << filename << '\n';
            return std::vector<std::string>{};
        }

        std::vector<std::string> lines;
        std::string line;
        while (std::getline(file, line)) {
            lines.push_back(line);
        }
        file.close();

        return lines;
    }

    // Get file size
    std::uintmax_t GetFileSize(const std::string& filename) {
        try {
            return std::filesystem::file_size(filename);
        } catch (std::filesystem::filesystem_error& e) {
            std::cerr << e.what() << '\n';
            return 0;
        }
    }
};

int main(int argc, char* argv[]) {
    // Create objects of both classes
    FileSystemOperations operations;
    FileOptimizer optimizer;

    if (argc != 2) {
        std::cerr << "Usage: " << argv[0] << " path\n";
        return 1;
    }

    std::string path = argv[1];

    // Execute functions from FileOptimizer
    optimizer.CheckFileSystemIntegrity(path);
    optimizer.ManagePermissions(path);
    optimizer.CheckDiskSpace(path);

    // Execute functions from FileSystemOperations
    operations.CreateFile("test.txt");
    operations.WriteToFile("test.txt", "Hello, world!");
    auto lines = operations.ReadFromFile("test.txt");
    for (const auto& line : lines) {
        std::cout << line << '\n';
    }
    std::cout << "Size of file: " << operations.GetFileSize("test.txt") << " bytes" << '\n';
    operations.DeleteFile("test.txt");
    operations.RenameFile("test1.txt", "test2.txt");
    operations.MoveFile("test2.txt", "dir/test2.txt");
    operations.ListFiles("./");

    return 0;
}
