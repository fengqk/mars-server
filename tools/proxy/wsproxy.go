package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/beevik/etree"
	"github.com/gorilla/websocket"
)

var g_tcpAddress string
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func forwardWeb(dst net.Conn, src *websocket.Conn) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("%s: forward from web failed: %s", time.Now().Format(time.Stamp), err)
		}
		if src != nil {
			src.Close()
		}
		if dst != nil {
			dst.Close()
		}
	}()

	for {
		if (src == nil) || (dst == nil) {
			log.Printf("forwardWeb conn is null")
			return
		}

		_, message, err := src.ReadMessage()
		if err != nil {
			log.Printf("forwardWeb read message failed")
			return
		}

		_, err = dst.Write(message)
		if err != nil {
			log.Printf("forwardWeb write message failed")
			return
		}
	}
}

func forwardTcp(dst *websocket.Conn, src net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("%s: forward from tcp failed: %s", time.Now().Format(time.Stamp), err)
		}
		if src != nil {
			src.Close()
		}
		if dst != nil {
			dst.Close()
		}
	}()

	buff := make([]byte, 1024)
	for {
		if (src == nil) || (dst == nil) {
			log.Printf("forwardTcp conn is null")
			return
		}

		n, err := src.Read(buff[0:])
		if err != nil {
			log.Printf("forwardTcp read message failed")
			return
		}

		err = dst.WriteMessage(websocket.BinaryMessage, buff[0:n])
		if err != nil {
			log.Printf("forwardTcp write message failed")
			return
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade failed")
		return
	}

	conn, err := net.Dial("tcp", g_tcpAddress)
	if err != nil {
		log.Printf("conn failed")
		return
	}

	go forwardWeb(conn, ws)
	go forwardTcp(ws, conn)
}

func argsUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s -host=ip:port -port=listen_port [optional]\noption:\n", os.Args[1])
	flag.PrintDefaults()
}

func argsParse() (string, int, string, string, error) {
	var host string
	var port int
	var certFile string
	var keyFile string

	flag.StringVar(&host, "host", "", "gameserver tcpaddr ip:port")
	flag.IntVar(&port, "port", 0, "websocket listen port")
	flag.StringVar(&certFile, "tlscert", "", "TLS cert file path")
	flag.StringVar(&keyFile, "tlskey", "", "TLS key file path")
	flag.Usage = argsUsage
	flag.Parse()

	if host == "" || port == 0 {
		return host, port, certFile, keyFile, fmt.Errorf("host or port error")
	}

	return host, port, certFile, keyFile, nil
}

func xmlsParse() (string, int, string, string, error) {
	var host string
	var port int
	var certFile string
	var keyFile string

	doc := etree.NewDocument()
	if err := doc.ReadFromFile("./srvconf.xml"); err != nil {
		log.Printf("open srvconf failed")
		return host, port, certFile, keyFile, fmt.Errorf("open srvconf failed")
	}

	root := doc.SelectElement("srvconf")
	connElement := root.FindElement("./conn")
	connAddr := connElement.FindElement("./addr1").Text()
	connPort := connElement.FindElement("./export1").Text()
	host = connAddr + ":" + connPort

	wsElement := root.FindElement("./websocket")
	port, _ = strconv.Atoi(wsElement.FindElement("port").Text())
	certFile = wsElement.FindElement("certfile").Text()
	keyFile = wsElement.FindElement("keyfile").Text()

	return host, port, certFile, keyFile, nil
}

type LogFile struct {
	fp string
}

func (lf *LogFile) Write(p []byte) (int, error) {
	f, err := os.OpenFile(lf.fp, os.O_CREATE|os.O_APPEND, 0x666)
	defer f.Close()

	if err != nil {
		return -1, err
	}
	return f.Write(p)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//lf := LogFile{fp: "log.txt"}
	//log.SetOutput(&lf)

	//tcpAddress, port, certFile, keyFile, err := argsParse()
	tcpAddress, port, certFile, keyFile, err := xmlsParse()
	if err != nil {
		argsUsage()
		log.Fatalln(err)
	}

	g_tcpAddress = tcpAddress
	portString := fmt.Sprintf(":%d", port)
	log.Printf("start server on port: %d, game server: %s", port, tcpAddress)

	http.HandleFunc("/", handler)

	if certFile != "" && keyFile != "" {
		err = http.ListenAndServeTLS(portString, certFile, keyFile, nil)
	} else {
		err = http.ListenAndServe(portString, nil)
	}

	if err != nil {
		log.Fatalln(err)
	}
}
