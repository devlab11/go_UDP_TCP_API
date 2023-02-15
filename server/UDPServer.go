package server

import (
	"fmt"
	"net"	
	"log"
	"strconv"	
	"encoding/hex"
	"go_testtask/config"
)

type UDPServer struct {
	Host    string	
	Port 	int
	Listener net.Listener
	onPacket   func(server *UDPServer, addr *net.UDPAddr, packet []byte)
	Connection net.Conn
	}

var (
	G_UDPServer *UDPServer
)

func InitUDPServer() {
	var (
		udpServ *UDPServer
	)
	udpServ = &UDPServer{
		Host: 	config.G_Config.ServerAddress,		
		Port:	config.G_Config.UDPPort,
	}
	G_UDPServer = udpServ
	return
}	

func (server *UDPServer) CreateListener() (*UDPServer) {
	addr := fmt.Sprintf("%v:%v", G_UDPServer.Host, strconv.Itoa(G_UDPServer.Port))
	udpAddr, err := net.ResolveUDPAddr("udp", addr)

	if err != nil {
		log.Fatal("Wrong UDP addres: ", err.Error())
	}
	conn, err := net.ListenUDP("udp", udpAddr)

	//defer conn.Close()

	if err != nil {		
		log.Fatal("Conn UDP is not created: ", err.Error())	
	}

	log.Println("Server " + "udp" + " is srarted: " + G_UDPServer.Host + ":" +  strconv.Itoa(G_UDPServer.Port))

	for {
		buf := make([]byte, 4096)
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal("Error with reading UDP connection", err.Error())
			
		}		
		log.Println("Received UDP packet:", hex.EncodeToString(buf[0:n]), "to the ", server.Connection.LocalAddr().String(), clientAddr)
		//server.onPacket(server)
	}
}

// func (server *UDPServer) Listen() {	
// 	for {
// 		buf := make([]byte, 4096)
// 		n, clientAddr, err := server.Connection.ReadFromUDP(buf)
// 		if err != nil {
// 			log.Fatal("Error with reading UDP connection", err.Error())
// 			return
// 		}		
// 		log.Println("Received UDP packet:", hex.EncodeToString(buf[0:n]), "to the ", server.Connection.LocalAddr().String())
// 		//server.onPacket(server)
// 	}		
// }


// func Response(udpServer net.PacketConn, addr G_UDPServer.Host, buf []byte) {
// 	time := time.Now().Format(time.ANSIC)
// 	responseStr := fmt.Sprintf("time received: %v. Your message: %v!", time, string(buf))

// 	udpServer.WriteTo([]byte(responseStr), addr)
// }

// func doReceiveMessage(conn net.Conn) {
// 	var (
// 		sess 		 *Session
// 		sessinfo	 SessionInfo
// 		sessionId 	 string
// 		readChannel  chan []byte
// 		writeChannel chan *channelData
// 		readData     []byte
// 		passData     *channelData
// 	)

// 	G_UDPServer.Connects[conn.RemoteAddr().String()] = 0
// }
