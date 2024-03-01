package main

import (
	"bytes"
	"log"
	"net"
	"time"
)

func main() {
	go func() {
		time.Sleep(2 * time.Second)
		writeFile()
	}()
	fs := FileServer{}
	fs.Run()
}

type FileServer struct{}

func (fs *FileServer) Run() {
	l, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("file server running on port 3000")

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
		}
		go fs.ReedLoop(conn)
	}
}

func (fs *FileServer) ReedLoop(conn net.Conn) {
	//Memory issues tha data is more than 4 bytes
	buf := make([]byte, 4)
	log.Println("got connection from :", conn.RemoteAddr())
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("recive %d bytes form remote host", n)
	log.Println(string(buf))
	log.Println(buf)
}

func writeFile() {
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.Buffer{}
	buf.WriteString("Hello world")
	buf.WriteString("abdelhadi is here")
	n, err := conn.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote %d bytes to remote host", n)
}
