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

	listOfURLsToCheck := loadFile("urls.txt")

	for _, element := range listOfURLsToCheck {
		resp, err := http.Get(element)
		if err != nil {
			log.Fatal("There was a problem getting the URL : ", element)
		}
		responseCodes[element] = resp.StatusCode
	}

	saveFile(responseCodes)
}

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
