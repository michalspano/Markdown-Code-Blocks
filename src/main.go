package main

import (
    "bufio"
    "encoding/json"
    "os"
    "strings"
)

func main() {
    
    // Parse all command line arguments
    var args = os.Args[1:]

    // Exit with a non-zero status code if length of 'args' is not equal to 2
    if len(args) != 2 {
        panic("Expected 2 argument")
    }
    
    // Define output file's path
    writeFilePath := args[1]

    // Make sure that the file is empty
    if _, err := os.Stat(writeFilePath); err == nil {
        err := os.Remove(writeFilePath)
        if err != nil {
            return 
        }
    }

    // Open the file for writing
    writeFile, err := os.Create(writeFilePath)
    if err != nil {
        panic(err)
    }

    /*
    Load possible languages supported by Flavored Markdown
    Source: https://rdmd.readme.io/docs/code-blocks
    */

    fileExtBuff := readJsonFile("lang/lang.json")

    // Load files from the specified directory
    dirFiles := getFiles(args[0])

    for _, file := range dirFiles {

        // Ignore output file
        if file == writeFilePath {
            continue
        }

        // Validate file extensions
        currentFileExt := getFileExtension(file)
        toWrite := checkFileExtension(currentFileExt, fileExtBuff)

        // Process writing to the output file
        if toWrite {
            _, _ = writeFile.WriteString("## " + file + " `." + currentFileExt + "`\n\n")
            _, _ = writeFile.WriteString("```" + currentFileExt + "\n" + readFile(file) + "```\n\n")
        }
    }

    // Close opened files
    closeErr := writeFile.Close()
    if closeErr != nil {
        panic(closeErr)
    }
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
    for val := range files {
        files[val] = path + "/" + files[val]
    }
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

// Read contents of the json file to a map (a.k.a. dictionary)
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

// Create a function to check if the file extension is in the map of extensions -> return type bool
func checkFileExtension(fileExtension string, fileExtensions map[string][]string) bool {
    for _, value := range fileExtensions {
        // Iterate over multiple extensions for a given language
        for _, ext := range value {
            if ext == fileExtension {
                return true
            }
        }
    }
    return false
}
