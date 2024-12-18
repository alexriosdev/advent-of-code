package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func StrToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
