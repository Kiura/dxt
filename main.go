package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	shellCMD := read()
	// fmt.Println(1, string(shellCMD))
	cmd := exec.Command("sh", "-c", " "+string(shellCMD))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func read() []byte {
	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage:")
		fmt.Println("cat yourfile | dxt")
		return nil
	}

	// if info.Size() > 0 {
	reader := bufio.NewReader(os.Stdin)
	return readP(reader)
	// }

	// return nilwhich dlv
}

func readP(reader *bufio.Reader) []byte {
	file := []byte{'#', '!', '/', 'b', 'i', 'n', '/', 'b', 'a', 's', 'h', '\n'}
	input := []byte{}
	for {
		var err error
		i, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		i = bytes.Replace(i, []byte{'"'}, []byte("\\\""), -1)
		prepend := []byte{'e', 'c', 'h', 'o', ' ', '"'}
		i = append(prepend, i...)
		input = append(input, i...)
		input = append(input, []byte{'"', '\n'}...)
	}
	file = append(file, input...)
	return file
}
