package main

// ---------------------------------------------------------
// EXERCISE: Find and write the names of subdirectories to a file
//
//  Create a program that can get multiple directory paths from
//  the command-line, and prints only their subdirectories into a
//  file named: dirs.txt
//
//
//  1. Get the directory paths from command-line
//
//  2. Append the names of subdirectories inside each directory
//     to a byte slice
//
//  3. Write that byte slice to dirs.txt file
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Please provide directory paths
//
//   go run main.go dir/ dir2/
//
//   cat dirs.txt
//
//     dir/
//             subdir1/
//             subdir2/
//
//     dir2/
//             subdir1/
//             subdir2/
//             subdir3/
// ---------------------------------------------------------

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please provide directory paths")
		return
	}

	var toWrite []byte
	for _, directories := range args {

		files, err := ioutil.ReadDir(directories)
		if err != nil {
			log.Fatal(err)
		}
		toWrite = append(toWrite, directories...)
		toWrite = append(toWrite, '\n')

		for _, file := range files {
			if os.FileInfo.IsDir(file) {
				toWrite = append(toWrite, '\t', '\t')
				toWrite = append(toWrite, file.Name()...)
				toWrite = append(toWrite, '/', '\n')
			}

		}
		toWrite = append(toWrite, '\n')
	}
	err := ioutil.WriteFile("dirs.txt", toWrite, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

}
