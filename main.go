package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var inputFile string
	var outputFile string

	flag.StringVar(&inputFile, "input", "input.txt", "The name of the text file to read")
	flag.StringVar(&outputFile, "output", "output.csv", "The name of the csv file to write output to")
	flag.Parse()

	textFile, _ := os.Open(inputFile)

	newFile, _ := os.Create(outputFile)

	lines := bufio.NewScanner(textFile)
	for lines.Scan() {
		txt := lines.Text()
		if !strings.Contains(txt, " - ") && !strings.Contains(txt, " = ") {
			continue
		}

		txt = strings.ReplaceAll(txt, " = ", " - ")
		txt = strings.TrimSpace(txt)
		txt = strings.ReplaceAll(txt, "\"", "\"\"")
		parts := strings.Split(txt, " - ")
		parts[0] = fmt.Sprintf("\"%v\"", parts[0])
		parts[1] = fmt.Sprintf("\"%v\"", parts[1])
		txt = strings.Join(parts, ",")
		partsRev := []string{parts[1], parts[0]}
		txtRev := strings.Join(partsRev, ",")

		newFile.WriteString(txt + ",\n")
		newFile.WriteString(txtRev + ",\n")
	}
}
