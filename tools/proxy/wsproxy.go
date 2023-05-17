package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"

	"github.com/gorilla/websocket"
)

var g_tcpAddress string
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func copyWebscoket2TcpWorker(dst net.Conn, src *websocket.Conn, doneCh chan<- bool) {
	defer func() {
		doneCh <- true
	}()

	for {
		_, message, err := src.ReadMessage()
		if err != nil {
			log.Fatalln("copyWebscoket2TcpWorker read message failed")
			return
		}

		_, err = dst.Write([]byte(message))
		if err != nil {
			log.Fatalln("copyWebscoket2TcpWorker write message failed")
			return
		}
	}
}

func copyTcp2WebsocketWorker(dst *websocket.Conn, src net.Conn, doneCh chan<- bool) {
	defer func() {
		doneCh <- true
	}()

	buff := make([]byte, 1024)
	for {
		n, err := src.Read(buff)
		if err != nil {
			log.Fatalln("copyTcp2WebsocketWorker read message failed")
			return
		}

		err = dst.WriteMessage(websocket.BinaryMessage, buff[:n])
		if err != nil {
			log.Fatalln("copyTcp2WebsocketWorker write message failed")
			return
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln("upgrade failed")
		return
	}

	conn, err := net.Dial("tcp", g_tcpAddress)
	if err != nil {
		log.Fatalln("conn failed")
		return
	}

	doneCh := make(chan bool)

	go copyWebscoket2TcpWorker(conn, ws, doneCh)
	go copyTcp2WebsocketWorker(ws, conn, doneCh)

	<-doneCh
	conn.Close()
	ws.Close()
	<-doneCh
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s -host=ip:port -port=listen_port [optional]\noption:\n", os.Args[1])
	flag.PrintDefaults()
}

func parseArgs() (string, int, string, string, error) {
	var host string
	var port int
	var certFile string
	var keyFile string

	flag.StringVar(&host, "host", "", "gameserver tcpaddr ip:port")
	flag.IntVar(&port, "port", 0, "websocket listen port")
	flag.StringVar(&certFile, "tlscert", "", "TLS cert file path")
	flag.StringVar(&keyFile, "tlskey", "", "TLS key file path")
	flag.Usage = usage
	flag.Parse()

	if host == "" || port == 0 {
		return host, port, certFile, keyFile, fmt.Errorf("host or port error")
	}

	return host, port, certFile, keyFile, nil
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	tcpAddress, port, certFile, keyFile, err := parseArgs()
	if err != nil {
		usage()
		log.Fatalln(err)
		os.Exit(1)
	}

	g_tcpAddress = tcpAddress
	portString := fmt.Sprintf(":%d", port)
	log.Printf("start server on port: %d,game server: %s", port, tcpAddress)

	if certFile != "" && keyFile != "" {
		err = http.ListenAndServeTLS(portString, certFile, keyFile, nil)
	} else {
		err = http.ListenAndServe(portString, nil)
	}

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
