package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var responseCodes = make(map[string]int)

	// snag the file name/path for the list of URLs to load
	args := os.Args

	// load the file into memory
	listOfURLsToCheck := loadFile(args[1])

	// this will iterate over the list of URLs and check the response code of each
	// using the http.Get package
	for _, element := range listOfURLsToCheck {
		resp, err := http.Get(element)
		if err != nil {
			log.Fatal("There was a problem getting the URL : ", element)
		}
		responseCodes[element] = resp.StatusCode
	}

	// save off the map of response codes to a CSV for further use
	saveFile(responseCodes)
}

// loadFile will load a basic file line by line into memory and return it as a string slice
func loadFile(path string) []string {
	var tempData []string
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("The file could not be found.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tempData = append(tempData, scanner.Text())
	}

	return tempData
}

// saveFile will check if the default file exists in the CWD
// and if it does, remove it and recreate it with the data we want the user to have
func saveFile(responseCodes map[string]int) {
	if _, err := os.Stat("ResponseCodes.csv"); os.IsExist(err) {
		err := os.Remove("ResponseCodes.csv")
		if err != nil {
			log.Fatal("There was a problem removing the file from the directory.")
		}
	}

	file, err := os.Create("ResponseCodes.csv")

	if err != nil {
		log.Fatal("Error creating CSV file.")
	}

	defer file.Close()

	w := bufio.NewWriter(file)

	for index, element := range responseCodes {
		w.WriteString(index + "," + strconv.Itoa(element) + "\n")
	}

	w.Flush()
}
