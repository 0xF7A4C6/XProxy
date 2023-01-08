package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// fixed: "bufio.Scanner: token too long"
// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
func ReadLines(path string) ([]string, error) {
	f, err := os.Open(fmt.Sprintf("data/%s", path))	
	if HandleError(err) {
		return nil, err
	}

	defer f.Close()

	r := bufio.NewReader(f)
	bytes, lines := []byte{}, []string{}

	for {
		line, isPrefix, err := r.ReadLine()
		if HandleError(err) {
			break
		}

		bytes = append(bytes, line...)
		if !isPrefix {
			str := strings.TrimSpace(string(bytes))
			
			if len(str) > 0 {
				lines = append(lines, str)
				bytes = []byte{}
			}
		}
	}

	if len(bytes) > 0 {
		lines = append(lines, string(bytes))
	}

	return lines, nil
}

func AppendFile(FileName string, Content string) {
	File, err := os.OpenFile(fmt.Sprintf("data/%s", FileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if HandleError(err) {
		return
	}

	_, err = File.WriteString(Content + "\n")
	if HandleError(err) {
		return
	}

	File.Close()
}

func RemoveDuplicateStr(strSlice []string) []string {
    allKeys, list := make(map[string]bool), []string{}

    for _, item := range strSlice {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }

    return list
}

func RemoveLine(path string, line string) {
	file, err := os.Open(fmt.Sprintf("data/%s", path))
	if HandleError(err) {
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		if scanner.Text() != line {
			lines = append(lines, scanner.Text())
		}
	}

	file, err = os.OpenFile(fmt.Sprintf("data/%s", path), os.O_TRUNC|os.O_WRONLY, 0600)
	if HandleError(err) {
		return
	}

	defer file.Close()
	for _, line := range lines {
		if _, err = file.WriteString(line + "\n"); err != nil {
			panic(err)
		}
	}
}

func InSlice(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}

	return false
}