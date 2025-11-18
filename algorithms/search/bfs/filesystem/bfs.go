// Breadth-First Search (BFS) - File Finder
// Time: O(n) - where n is total number of files and directories
// Space: O(w) - where w is maximum width (number of directories at one level)

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FindFileBFS(dirName string, fileName string) (string, bool) {
	// Initialize the queue with starting directory
	queue := []string{dirName}

	for len(queue) > 0 {
		dir := queue[0]   // take the first folder from the queue (FIFO)
		queue = queue[1:] // dequeue

		// In Unix-like OS everything (including directories) is a file,
		// but os.ReadDir separates files and directories
		files, _ := os.ReadDir(dir)
		for _, file := range files {
			fullPath := filepath.Join(dir, file.Name())
			if file.IsDir() {
				queue = append(queue, fullPath) // is directory — add it to the queue to explore later
			} else if strings.Contains(file.Name(), fileName) {
				return fullPath, true // is a file and matches our search
			}
		}
	}
	return "", false
}

func main() {
	// Placeholder: replace with actual directory path and file name
	path, found := FindFileBFS("/path/to/dir", "file.txt")
	if found {
		fmt.Println(path)
	}
}

/*
How BFS file search works:

Explores directories level by level (breadth-first traversal).
Uses a queue to defer subdirectory exploration until all items at current level are checked.

Parameters:
  - dirName: root folder to start search
  - fileName: name or part of the file name to search for

Returns:
  - string: full path to the found file, or empty string if not found
  - bool: true if file was found, false otherwise

Used Go standard library:

os.ReadDir(dir)
  - Reads the contents of a directory
  - Returns []fs.DirEntry and an error

filepath.Join(dir, file)
  - Joins path segments into a normalized path string
  - Example: filepath.Join("pics", "a.png") → "pics/a.png"

item.Name()
  - Method of fs.DirEntry interface
  - Returns the name of the file or directory

item.IsDir()
  - Method of fs.DirEntry interface
  - Returns true if the entry is a directory, false if it's a file

strings.Contains(str, substr)
  - Checks whether substr is within str
  - Returns true if substr is found, false otherwise
*/
