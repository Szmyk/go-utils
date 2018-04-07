package files

import (
  "os"
  "bufio"
  "bytes"
  "fmt"
)

func ReadLines(path string) []string {
	file, _ := os.Open(path)

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func ReadFile(path string) string {
	filerc, _ := os.Open(path)

	buf := new(bytes.Buffer)
	buf.ReadFrom(filerc)
	content := buf.String()

	return content
}

func WriteLines(lines []string, path string) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	for _, line := range lines {
		fmt.Fprintln(w, line)
	}

	return w.Flush()
}
