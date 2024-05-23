package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type FileServer struct {
}

func (fs *FileServer) start() {
	ln, err := net.Listen("tcp", ":3000")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go fs.readLoop(conn)
	}
}

func (fs *FileServer) readLoop(conn net.Conn) {
	// simulando um valor de um arquivo,
	// na logica retorna uma lista de byetes com o tamanho do arquivo
	buf := make([]byte, 2048)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		file := buf[:n]

		fmt.Println(file)
		fmt.Printf("Received %d bytes over the network\n", n)
	}
}

func sendFile(size int) error {
	file := make([]byte, size)

	_, err := io.ReadFull(rand.Reader, file)

	if err != nil {
		return err
	}

	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		return err
	}

	n, err := io.Copy(conn, bytes.NewReader(file))
	// n, err := conn.Write(file)
	if err != nil {
		return err
	}

	fmt.Printf("Written %d bytes over the network\n", n)
	return nil

}

func main() {
	go func() {
		time.Sleep(2 * time.Second)
		sendFile(8000)
	}()
	server := &FileServer{}
	server.start()
}
