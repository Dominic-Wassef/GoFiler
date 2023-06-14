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

    // Copy a file
    void CopyFile(const std::string& source, const std::string& destination) {
        try {
            std::filesystem::copy(source, destination, std::filesystem::copy_options::overwrite_existing);
        } catch (std::filesystem::filesystem_error& e) {
            std::cerr << "Error occurred while copying file: " << e.what() << '\n';
        }
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
            std::cerr << "Error occurred while getting file size: " << e.what() << '\n';
            return 0;
        }
    }
};

int main(int argc, char* argv[]) {
    if (argc < 2) {
        std::cerr << "Usage: " << argv[0] << " <operation> <parameters>\n"
                  << "Available operations:\n"
                  << "optimize <path>\n"
                  << "create <filename>\n"
                  << "delete <filename>\n"
                  << "rename <oldname> <newname>\n"
                  << "move <source> <destination>\n"
                  << "list <path>\n"
                  << "write <filename> <data>\n"
                  << "read <filename>\n"
                  << "getsize <filename>\n"
                  << "copy <source> <destination>\n";
        return 1;
    }

    std::string operation = argv[1];

    // Create objects of both classes
    FileSystemOperations operations;
    FileOptimizer optimizer;
    if ((operation == "rename" || operation == "move" || operation == "write" || operation == "copy") && argc < 4) {
        std::cerr << "Error: Operation " << operation << " requires two arguments.\n";
        return 1;
    }
    if ((operation == "create" || operation == "delete" || operation == "list" || operation == "read" || operation == "getsize") && argc < 3) {
        std::cerr << "Error: Operation " << operation << " requires one argument.\n";
        return 1;
    }
    if (operation == "optimize") {
        std::string path = argv[2];
        optimizer.RunFileOptimizer(path);
    }
    else if (operation == "create") {
        operations.CreateFile(argv[2]);
    }
    else if (operation == "delete") {
        operations.DeleteFile(argv[2]);
    }
    else if (operation == "rename") {
        operations.RenameFile(argv[2], argv[3]);
    }
    else if (operation == "move") {
        operations.MoveFile(argv[2], argv[3]);
    }
    else if (operation == "list") {
        operations.ListFiles(argv[2]);
    }
    else if (operation == "write") {
        operations.WriteToFile(argv[2], argv[3]);
    }
    else if (operation == "read") {
        auto lines = operations.ReadFromFile(argv[2]);
        for (const auto& line : lines) {
            std::cout << line << '\n';
        }
    }
    else if (operation == "getsize") {
        std::cout << "Size of file: " << operations.GetFileSize(argv[2]) << " bytes" << '\n';
    }
    else if (operation == "copy") {
        operations.CopyFile(argv[2], argv[3]);
    }
    else {
        std::cerr << "Invalid operation.\n";
        return 1;
    }

    return 0;
}