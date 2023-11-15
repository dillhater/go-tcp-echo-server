package main

import (
	"flag"
	"log"
	"net"
	"strconv"
)

func main() {
	port := flag.Int("port", 8282, "Port to accept connections on")
	host := flag.String("host", "127.0.0.1", "Host to bind to")

	flag.Parse()

	listen, err := net.Listen("tcp", *host+":"+strconv.Itoa(*port))
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Listening to connections on port " + strconv.Itoa(*port))
	defer listen.Close()

	for {
		con, err := listen.Accept()
		if err != nil {
			log.Panicln(err)
		}
		go handleRequest(con)
	}
}

func handleRequest(con net.Conn) {
	log.Println("Accepting new connection")
	defer con.Close()
	defer log.Println("Close connection")
	buf := make([]byte, 1024)
	size, err := con.Read(buf)
	if err != nil {
		log.Panicln(err)
	}
	data := buf[:size]
	log.Println("Read from client: ", string(data))
	con.Write(data)
}
