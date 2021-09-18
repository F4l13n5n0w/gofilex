package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("hello")

	file, err := os.Open("WiresharkPortable64_3.5.1rc0-213-gb81192d312c9.paf.exe")
	if err != nil {
		fmt.Println("os.Open err", err)
		return
	}
	buf := make([]byte, 4096)

	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			fmt.Println("sent !!")
			return
		}
		if err != nil {
			fmt.Println("file.Read err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
		break
	}
}
