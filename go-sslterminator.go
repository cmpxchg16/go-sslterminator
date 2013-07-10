package main

import (
	"crypto/tls"
	"log"
	"net"
	"io"
	"runtime"
	"flag"
)
    
var localAddress string
var backendAddress string
var certificatePath string
var keyPath string

func init() {
	flag.StringVar(&localAddress, "l", ":443", "local address")
	flag.StringVar(&backendAddress, "b", ":80", "backend address")
	flag.StringVar(&certificatePath, "c", "server.pem", "SSL certificate path")
	flag.StringVar(&keyPath, "k", "server.key", "SSL key path")
}
    
func main() {
	flag.Parse()
	
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	cert, err := tls.LoadX509KeyPair(certificatePath, keyPath)
	if err != nil {
		log.Fatalf("error in tls.LoadX509KeyPair: %s", err)
	}
	
	config := tls.Config{Certificates: []tls.Certificate{cert}, 
						 CipherSuites: []uint16{tls.TLS_RSA_WITH_AES_256_CBC_SHA}}
	
	listener, err := tls.Listen("tcp", localAddress, &config)
	if err != nil {
		log.Fatalf("error in tls.Listen: %s", err)
	}
	
	log.Printf("local server on: %s, backend server on: %s", localAddress, backendAddress)
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error in listener.Accept: %s", err)
			break
		}
		
		go handle(conn)
	}
}

func handle(clientConn net.Conn) {
	tlsconn, ok := clientConn.(*tls.Conn)
	if ok {
		
		err := tlsconn.Handshake()
		if err != nil {
			log.Printf("error in tls.Handshake: %s", err)
		}
		
		backendConn, err := net.Dial("tcp", backendAddress)
		if err != nil {
			log.Printf("error in net.Dial: %s", err)
		}

		go Tunnel(clientConn, backendConn)
        go Tunnel(backendConn, clientConn)
	}
}

func Tunnel(from, to io.ReadWriteCloser) {
    io.Copy(from, to)
    from.Close()
    to.Close()
}
