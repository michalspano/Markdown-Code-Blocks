package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Print all command-line arguments without the first one
    var args = os.Args[1:]
    fmt.Println(args)

    // Exit with a non-zero status code if length of 'args' is not equal to 1
    if len(args) != 2 {
        panic("Expected 2 argument")
    }

    writeFilePath := args[1]

    // Make sure that the file is empty
    if _, err := os.Stat(writeFilePath); err == nil {
        os.Remove(writeFilePath)
    }

    // Open a file to write into
    writeFile, err := os.Create(writeFilePath)
    if err != nil {
        panic(err)
    }

    fileExtPath := "src/lang.json"
    fileExtBuff := readJsonFile(fileExtPath)

    dirFiles := getFiles(args[0])

    for _, file := range dirFiles {

        if file == "out.md" {
            continue
        }

        currentFileExt := getFileExtension(file)
        fmt.Println("Current file extension: ", currentFileExt)
        toWrite := checkFileExtension(currentFileExt, fileExtBuff)

        if toWrite {
            writeFile.WriteString("## " + file + " `." + currentFileExt + "`\n\n")
            writeFile.WriteString("```" + currentFileExt + "\n" + readFile(file) + "```\n\n")
        }
    }

    // Close opened files
    writeFile.Close()
}

// Define a function to get names of the files from the directory
func getFiles(path string) []string {
    // Create a slice to store the names of the files
    var files []string

    // Open the directory
    dir, err := os.Open(path)
    if err != nil {
        panic(err)
    }

    // Read the directory
    files, err = dir.Readdirnames(0)
    if err != nil {
        panic(err)
    }

    // Close the directory
    Err := dir.Close()
    if Err != nil {
        return nil
    }

    // Return the slice of file names
    return files
}


// Read the file and put its contents in a string
func readFile(fileName string) string {
    // Open the file
    file, err := os.Open(fileName)
    if err != nil {
        panic(err)
    }

    // Create a scanner to read the file
    scanner := bufio.NewScanner(file)

    // Create a string to store the contents of the file
    var contents string

    // Read the file
    for scanner.Scan() {
        contents += scanner.Text() + "\n"
    }

    // Close the file
    Err := file.Close()
    if Err != nil {
        return ""
    }

    // Return the contents of the file
    return contents
}

// Read contents of the json file
func readJsonFile(path string) map[string][]string {
    // Open the file
    file, err := os.Open(path)
    if err != nil {
        panic(err)
    }

    // Create a decoder to read the file
    decoder := json.NewDecoder(file)

    // Create a map to store the contents of the file
    var fileExt map[string][]string

    // Decode the file
    err = decoder.Decode(&fileExt)
    if err != nil {
        panic(err)
    }

    // Close the file
    Err := file.Close()
    if Err != nil {
        return nil
    }

    // Print the contents of the file
    return fileExt
}

// Function to get the extension of the file
func getFileExtension(fileName string) string {
    // Get the last index of the period
    index := strings.LastIndex(fileName, ".")

    // Return the extension of the file
    return fileName[index+1:]
}

// Get name of the file without file extension
func getFileName(fileName string) string {
    // Get the last index of the period
    index := strings.LastIndex(fileName, ".")

    // Return the name of the file
    return fileName[:index]
}

// Create a function to check if the file extension is in the map of extensions
func checkFileExtension(fileExtension string, fileExtensions map[string][]string) bool {
    for _, value := range fileExtensions {
        for _, ext := range value {
            if ext == fileExtension {
                return true
            }
        }
    }
    return false
}
