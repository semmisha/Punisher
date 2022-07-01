package repository

import (
	"bufio"
	"log"
	"os"
)

func FileToVar(path string) []string {
	var variable []string
	file, err := os.OpenFile(path, os.O_RDONLY, 0555)
	if err != nil {
		log.Fatalf("Unable to open file: %v, error: %v\n", path, err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		variable = append(variable, scanner.Text())

	}
	//fmt.Println(variable)
	return variable
}
