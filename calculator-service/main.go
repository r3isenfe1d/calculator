package main

import (
	"context"
	"flag"
	"log"
	"net"
	"strings"

	"calculator/proto"

	"google.golang.org/grpc"
)

func main() {
	// Read port and option from command arguments
	port, option := readCommandArgs()

	// Create a server and register that
	s := grpc.NewServer()
	srv := &server{option: option}
	databus.RegisterDatabusServiceServer(s, srv)

	// Listen port...
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to %v", err)
	} else {
		log.Printf("Server listenning at %v", lis.Addr())
	}

	// Serve listeninng
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

type server struct {
	option string
	databus.UnimplementedDatabusServiceServer
}

func (s *server) Send(ctx context.Context, req *databus.SendRequest) (*databus.SendResponse, error) {
	var result float32
	switch s.option {
	case "add":
		result = req.GetPrm1() + req.GetPrm2()
	case "sub":
		result = req.GetPrm1() - req.GetPrm2()
	case "mul":
		result = req.GetPrm1() * req.GetPrm2()
	case "div":
		result = req.GetPrm1() / req.GetPrm2()
	}
	log.Println(req.GetPrm1(), s.option, req.GetPrm2(), "=", result)
	return &databus.SendResponse{Result: result}, nil
}

func readCommandArgs() (string, string) {
	// Parse command arguments check the number of arguments
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatalln("Fail: not enough arguments")
	}

	port := ":"
	// If port looks like localhost:8080
	if strings.Contains(flag.Arg(0), ":") {
		port += strings.Split(flag.Arg(0), ":")[1]
		log.Println(port)
	} else {
		port += flag.Arg(0)
	}

	// Check option exists
	option := flag.Arg(1)
	switch option {
	case "mul", "sub", "add", "div":
		log.Printf("Option: %s", option)
	default:
		log.Fatalf("Failed option: %s", option)
	}

	return port, option
}
