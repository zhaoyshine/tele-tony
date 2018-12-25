package fileOperation

import (
	"io/ioutil"
	"fmt"
	"os"
)

func ReadFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	return string(content)
}

func WriteFile(filename string, content string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(content)
}
