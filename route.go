package main

import (
	"io"
	"log"
	"net"
)

type route struct {
	src string
	dst string
}

func (r route) handle(src net.Conn) {
	defer src.Close()

	dst, err := net.Dial("tcp", r.dst)
	if err != nil {
		log.Printf("Error connecting to dstination: %s\n", r.dst)
		return
	}

	defer dst.Close()

	log.Printf("Routing %s -> %s", r.src, r.dst)

	go io.Copy(src, dst)
	io.Copy(dst, src)

	log.Printf("Connection closed: %s -> %s", r.src, r.dst)
}

func (r route) listen() {
	log.Printf("Creating listener: %s -> %s\n", r.src, r.dst)

	l, err := net.Listen("tcp", r.src)
	if err != nil {
		log.Fatalf("Error setting up listener: %s -> %s: %s", r.src, r.dst, err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s", err)
		}
		go r.handle(conn)
	}
}
