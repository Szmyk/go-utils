package files

import (
  "os"
  "strings"
  "bufio"
)

func readLines(path string) []string {
	file, _ := os.Open(path)

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func readFile(path string) string {
	filerc, _ := os.Open(path)

	buf := new(bytes.Buffer)
	buf.ReadFrom(filerc)
	content := buf.String()

	return content
}
