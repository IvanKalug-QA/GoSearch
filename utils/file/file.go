package file

import (
	"fmt"
	"os"
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
