package main

import (
	"calculator/proto"
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ip, prm1, prm2 := readCommandArgs()

	// Connect to server
	conn, err := grpc.Dial(
		ip, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}
	defer conn.Close()

	// Create a client and send request
	c := databus.NewDatabusServiceClient(conn)
	res, err := c.Send(context.Background(),
		&databus.SendRequest{Prm1: prm1, Prm2: float32(prm2)})
		
	if err != nil {
		log.Fatalf("Could not to count: %v", err)
	}

	fmt.Println("Result:", res.GetResult())
}

func readCommandArgs() (string, float32, float32) {
	// Parse command arguments check the number of arguments
	flag.Parse()
	if flag.NArg() < 3 {
		log.Fatalln("Fail: not enough arguments")
	}

	ip := flag.Arg(0)
	
	// Convert params to float32
	prm1, err := strconv.ParseFloat(flag.Arg(1), 32)
	if err != nil {
		log.Fatalf("Failed to convert arguments %v\n", err)
	}
	
	prm2, err := strconv.ParseFloat(flag.Arg(2), 32)
	if err != nil {
		log.Fatalf("Failed to convert arguments %v\n", err)
	}
	
	return ip, float32(prm1), float32(prm2)
}
