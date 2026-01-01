package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Write(d []string) {
	file, err := os.OpenFile("results.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Ошибка открытия файла: %v\n", err)
		return
	}
	defer file.Close()
	for _, f := range d {
		_, err := file.WriteString(f)
		if err != nil {
			fmt.Printf("Ошибка при записи в файл: %v\n", err)
			return
		}
	}
}

func HasFile(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	return file, err
}

func ReadUrl(fileName string) []string {
	f, _ := HasFile(fileName)
	if f == nil {
		return nil
	}
	var urls []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		urls = append(urls, strings.TrimSpace(scanner.Text()))
	}
	return urls
}
