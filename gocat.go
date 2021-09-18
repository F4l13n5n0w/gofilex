package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

func sendFile(conn net.Conn, filepath, filesize string) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("os.Open err", err)
		return
	}
	buf := make([]byte, 4096)

	i := 1
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
		_, err = conn.Write(buf[:n])
		fmt.Printf("%d / %s\n", (n * i), filesize)
		i++
		if err != nil {
			fmt.Println("conn.Write err:", err)
			return
		}
	}
}

func recvFile(conn net.Conn, outFile string) {
	buf := make([]byte, 4096)

	file, err := os.Create(outFile)
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}

	i := 1
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("received !!")
			break
		}
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		file.Write(buf[:n])
		fmt.Printf("[+] recving %d\n", (n * i))
		i++
	}
}

func main() {

	if len(os.Args) != 5 {
		fmt.Println("[!] Usage:")
		fmt.Println("    ./gobaba server(-l)/client(-n) listen_ip:port put/get filename")
		fmt.Println("    E.g. ./gobaba -l 127.0.0.1:7777 get output.txt")
		fmt.Println("    or   ./gobaba -l 127.0.0.1:7777 put input.txt")
		fmt.Println("    E.g. ./gobaba -n 127.0.0.1:7777 get output.txt")
		fmt.Println("    or   ./gobaba -n 127.0.0.1:7777 put input.txt")
		return
	}

	list := os.Args
	strSrvCli := list[1]
	strConn := list[2]
	strCtrl := list[3]
	fileName := list[4]

	if strSrvCli == "-l" {
		fmt.Printf("[+] Listen on %s\n", strConn)
		listener, err := net.Listen("tcp", strConn)
		if err != nil {
			fmt.Println("net.Listen err:", err)
			return
		}

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("lis.Accept err:", err)
			return
		}

		if strCtrl == "get" {
			recvFile(conn, fileName)
		}

		if strCtrl == "put" {
			fileInfo, err := os.Stat(fileName)
			if err != nil {
				fmt.Println("os.Stat err", err)
				return
			}
			filename := fileInfo.Name()
			filesize := strconv.FormatInt(fileInfo.Size(), 10)

			fmt.Println("[+] Sending file:")
			fmt.Printf("    filename: %s\n", filename)
			fmt.Printf("    filensize: %s \n", filesize)

			sendFile(conn, fileName, filesize)
		}
	}

	if strSrvCli == "-n" {
		conn, err := net.Dial("tcp", strConn)
		if err != nil {
			fmt.Println("net.Dialt err", err)
			return
		}

		if strCtrl == "get" {
			recvFile(conn, fileName)
		}

		if strCtrl == "put" {
			fileInfo, err := os.Stat(fileName)
			if err != nil {
				fmt.Println("os.Stat err", err)
				return
			}
			filename := fileInfo.Name()
			filesize := strconv.FormatInt(fileInfo.Size(), 10)

			fmt.Println("[+] Sending file:")
			fmt.Printf("    filename: %s\n", filename)
			fmt.Printf("    filensize: %s \n", filesize)

			sendFile(conn, fileName, filesize)
		}
	}

}
