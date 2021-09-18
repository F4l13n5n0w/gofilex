package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func sendFile(conn net.Conn, filepath string) {
	file, err := os.Open(filepath)
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
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("conn.Write err:", err)
			return
		}
	}
}

func recvFile(conn net.Conn, outFile string) {
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	filename := string(buf[:n])
	fmt.Println("filename:", filename)
	if filename != "" {
		_, err = conn.Write([]byte("ok"))
		if err != nil {
			fmt.Println("conn.Write err:", err)
			return
		}
	} else {
		return
	}

	fmt.Println(filename)
	file, err := os.Create(outFile)
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}

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
	}
}

func main() {

	if len(os.Args) != 5 {
		fmt.Println("[!] Usage:")
		fmt.Println("    ./gofile server(-l)/client(-n) listen_ip:port put/get filename")
		fmt.Println("    E.g. ./gofile -l 127.0.0.1:7777 get output.txt")
		fmt.Println("    or   ./gofile -l 127.0.0.1:7777 put input.txt")
		fmt.Println("    E.g. ./gofile -n 127.0.0.1:7777 get output.txt")
		fmt.Println("    or   ./gofile -n 127.0.0.1:7777 put input.txt")
		return
	}

	list := os.Args
	strSrvCli := list[1]
	strConn := list[2]
	strCtrl := list[3]
	fileName := list[4]

	if strSrvCli == "-l" {
		fmt.Printf("[+] Listen on %s", strConn)
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

			_, err = conn.Write([]byte(filename))
			if err != nil {
				fmt.Println("conn.Write err", err)
				return
			}
			buf := make([]byte, 4096)

			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("conn.Read err", err)
				return
			}

			if string(buf[:n]) == "ok" {
				sendFile(conn, fileName)
			}
		}
	}

	if strSrvCli == "-n" {
		conn, err := net.Dial("tcp", strConn)
		if err != nil {
			fmt.Println("net.Dialt err", err)
			return
		}

		if strCtrl == "put" {
			fileInfo, err := os.Stat(fileName)
			if err != nil {
				fmt.Println("os.Stat err", err)
				return
			}
			filename := fileInfo.Name()

			_, err = conn.Write([]byte(filename))
			if err != nil {
				fmt.Println("conn.Write err", err)
				return
			}
			buf := make([]byte, 4096)

			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("conn.Read err", err)
				return
			}

			if string(buf[:n]) == "ok" {
				sendFile(conn, fileName)
			}
		}

		if strCtrl == "get" {
			recvFile(conn, fileName)
		}
	}

}
