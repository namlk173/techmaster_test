package helper

import (
	"bufio"
	"os"
	"strings"
)

func GetStringFromFile(fileName string) ([]string, error) {
	var lines []string

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		line = strings.Join(strings.Fields(line), " ")
		if line != "" {
			// Standardized data read from file text.
			flag := true
			a := strings.Index(line, "a)")
			b := strings.Index(line, "b)")
			c := strings.Index(line, "c)")
			if a != -1 && a != 0 {
				lines = append(lines, string(line[0:a-1]))
				lines = append(lines, string(line[a:]))
				flag = false
			}
			if b != -1 && b != 0 {
				lines = append(lines, string(line[0:b-1]))
				lines = append(lines, string(line[b:]))
				flag = false
			}
			if c != -1 && c != 0 {
				lines = append(lines, string(line[0:c-1]))
				lines = append(lines, string(line[c:]))
				flag = false
			}
			if flag {
				lines = append(lines, line)
			}
		}
	}
	return lines, nil
}
