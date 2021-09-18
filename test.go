package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/schollz/progressbar"
)

func main() {

	fileName := "/root/codes/WiresharkPortable_3.4.8.paf.exe"

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("os.Stat err", err)
		return
	}
	//filename := fileInfo.Name()
	filesize := strconv.FormatInt(fileInfo.Size(), 10)

	fmt.Println("[+] Sending file:")
	fmt.Printf("    filename: %s\n", fileName)
	fmt.Printf("    filensize: %s \n", filesize)

	//bar := pb.StartNew(int(fileInfo.Size()))
	bar := progressbar.DefaultBytes(fileInfo.Size(), "downloading")

	file, err := os.Open(fileName)
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
		//fmt.Println(string(buf[:n]))
		bar.Add(n)
		time.Sleep(time.Millisecond)
		//break
	}

}
