package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	// open file
	fl, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error opening file:",err)
		return
	}
	defer fl.Close()

	// read content
	scanner := bufio.NewScanner(fl)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// handle error
	if err := scanner.Err(); err != nil {
		fmt.Println("error handling:",err)
		return
	}

	// create output file
	flOutput, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:",err)
		return
	}
	defer flOutput.Close()

	// rad content of file input.txt
	scanner = bufio.NewScanner(fl)
	for scanner.Scan() {
		// modify 
		modifiedText := scanner.Text() + "Edited"
		_, err := flOutput.WriteString(modifiedText + "\n")
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
		fmt.Printf("original: %s, Modified: %s \n", scanner.Text(), modifiedText)
	}

	// handle error
	if err := scanner.Err(); err != nil {
		fmt.Println("error handling file:", err)
		return
	}
}