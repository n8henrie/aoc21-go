package files

import (
	"fmt"
	"io/ioutil"
)

func InputForDay(day int) (string, error) {
	path := fmt.Sprintf("d%02d/input.txt", day)
	input, err := ReadFile(path)
	if err != nil {
		return "", err
	}
	return input, nil
}

func ReadFile(path string) (string, error) {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(input), nil
}
