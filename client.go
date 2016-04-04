package main

import (
	"log"
	"net"
	"time"
)

func broadcastUdp(addr string) {
	udpAddr, err := net.ResolveUDPAddr("udp", addr) 	
	if err != nil {
		log.Fatal(err)
	}

	udpBroadcast, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer udpBroadcast.Close()

	for {
		udpBroadcast.Write([]byte("Not master"))
		time.Sleep(1000 * time.Millisecond)
		log.Println("Hei")
	}
}

func main() {
	log.Println(getMyIP())
	doneChannel := make(chan bool, 1)
	ipListChannel := make(chan []string, 1)
	doneChannel <- true
	//port := ":20010"
	broadcastAddr := "129.241.187.255:20060"

	go broadcastUdp(broadcastAddr)
	//go listenUdp(port, ipListChannel)

	log.Println(<-ipListChannel)
	<-doneChannel
}
