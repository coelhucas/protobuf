package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/coelhucas/protobuf/generated"
	"google.golang.org/protobuf/proto"
)

const (
	HOST = "localhost"
)

var (
	port = flag.Int("port", 6969, "The server port")
)

func handleIncomingRequest(conn net.Conn) {
	log.Println("Incoming request...")
    // store incoming data
    buffer := make([]byte, 1024)
	
    _, err := conn.Read(buffer)
    if err != nil {
        log.Fatal(err)
    }
	
    time := time.Now().Format("Monday, 02-Jan-06 15:04:05 MST")

	player := &pb.Player {
		Name: "Potato",
		Id: 1,
		Position: &pb.Player_Position {
			X: 1,
			Y: 2,
		},
	}

	out, err := proto.Marshal(player)

	if err != nil {
		log.Fatalln("Failed to encode Player:", err)
	}

	log.Println("Encoded Player:", out)
	length, err := conn.Write(out)
	
	log.Printf("Sent encoded player data, length %d bytes", length)
	
    conn.Write([]byte(time))
    conn.Close()
}


func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", HOST, *port))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
			}

		handleIncomingRequest(conn)
		}


}
