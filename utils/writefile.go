package utils

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// IOWrite OpenFile -> io.WriteString
func IOWrite(filename string) {
	str := "This is IOWrite test.\n"
	var f *os.File
	var err error
	if CheckFileExist(filename) {
		f, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0777)
		if err != nil {
			fmt.Println("file open fail", err)
			return
		}
	} else {
		f, err = os.Create(filename)
		if err != nil {
			fmt.Println("file cread fail", err)
			return
		}
	}

	n, err := io.WriteString(f, str)
	if err != nil {
		fmt.Println("write fail", err)
		return
	}
	fmt.Println("Write 2 File", n)
}

// IoutilWriteFile -
func IoutilWriteFile(filename string) {
	str := "This is IoutilWriteFile test.\n"

	var d = []byte(str)
	err := ioutil.WriteFile(filename, d, 0666)
	if err != nil {
		fmt.Println("WriteFile fail", err)
		return
	}

	fmt.Println("Write 2 File Successed")
}

// FileWrite -
func FileWrite(filename string) {
	str := "This is FileWrite test.\n"
	var d = []byte(str)

	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Create fail", err)
		return
	}
	defer f.Close()

	n, err := f.Write(d)
	fmt.Printf("Write %d bytes", n)
	n, err = f.WriteString("write")
	fmt.Printf("Write %d bytes", n)
	f.Sync()

}

func BufilWrite(filename string) {
	str := "This is FileWrite test.\n"
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("create file", err)
	}

	w := bufio.NewWriter(f)
	n, err := w.WriteString(str)
	fmt.Printf("write %d bytes", n)

	w.Flush()
	f.Close()
}

func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
