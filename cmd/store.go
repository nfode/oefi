package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var cacheDir = os.Getenv("HOME") + "/.cache/oefi"
var cacheFile = cacheDir + "/cache"

func init() {
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		os.Mkdir(cacheDir, os.ModePerm)
	}
}

func writeToCache(stations []string) {
	f, err := os.OpenFile(cacheFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	builder := strings.Builder{}
	for _, station := range stations {
		builder.WriteString(fmt.Sprintf("%v\n", station))
	}

	if _, err := f.WriteString(builder.String()); err != nil {
		log.Println(err)
	}
}

func findStations(search string) []string {
	if _, err := os.Stat(cacheFile); os.IsNotExist(err) {
		return []string{}
	}

	var result []string
	file, err := os.Open(cacheFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, search) {
			result = append(result, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
