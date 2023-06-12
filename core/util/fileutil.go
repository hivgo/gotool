package util

import (
	"io/ioutil"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func ReadTextFile(name string) (string, error) {
	bytes, err := ReadFile(name)
	return string(bytes), err
}
