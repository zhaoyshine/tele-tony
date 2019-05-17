package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func LoadFile(path string) int {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	i, _ := strconv.Atoi(string(content))
	return i
}

func WriteFile(path string, record int) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	t := strconv.Itoa(record)
	file.WriteString(t)
}
