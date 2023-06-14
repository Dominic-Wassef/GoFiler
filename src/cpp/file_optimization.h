#ifndef FILE_OPTIMIZATION_H
#define FILE_OPTIMIZATION_H

#include <string>
#include <filesystem>

class FileOptimizer {
public:
    void CheckFileSystemIntegrity(const std::string& fs);
    void RepairFileSystem(const std::string& fs);
    void ManagePermissions(const std::string& path);
    void CheckDiskSpace(const std::string& path);
    void RunFileOptimizer(const std::string& path);
};

#endif  // FILE_OPTIMIZATION_H
