package utils

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// ReadFile Read file to []byte .using ioutil.ReadFile()
func ReadFile(path string) string {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("read file failed!", err)
		return ""
	}
	return string(f)
}

// Open Open -> Read to file -> append []byte
func Open(path string) string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Opend file failed!", err)
		return ""
	}

	defer f.Close()
	var chunk []byte
	buf := make([]byte, 1024)

	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read buf failed", err)
			return ""
		}

		if n == 0 {
			break
		}

		chunk = append(chunk, buf[:n]...)
	}

	return string(chunk)
}

// OpenNewReader Opend -> NewReader ->append
func OpenNewReader(path string) string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Opend file failed!", err)
		return ""
	}

	defer f.Close()

	r := bufio.NewReader(f)
	var chunk []byte
	buf := make([]byte, 1024)

	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("Read failed!", err)
			panic(err)
		}
		if 0 == n {
			break
		}

		chunk = append(chunk, buf[:n]...)
	}
	return string(chunk)
}

// OpenReadALl Opend -> ReadAll
func OpenReadAll(path string) string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Open file failed!", err)
		return ""
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd failed!", fd)
		return ""
	}

	return string(fd)
}
