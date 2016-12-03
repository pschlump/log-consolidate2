// Simple gRPC server that consolidates log files.
package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	pb "github.com/pschlump/log-consolidate2/proto" // pb "google.golang.org/grpc/examples/chat/proto"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "testdata/server1.pem", "The TLS cert file")
	keyFile  = flag.String("key_file", "testdata/server1.key", "The TLS key file")
	port     = flag.Int("port", 10000, "The server port")
	logFile  = flag.String("log_file", "log/out.log", "The log output file")
)

type logItServer struct {
	fo *os.File
}

// LogMessage returns the feature at the given point. (PJS)
func (s *logItServer) IAmAlive(ctx context.Context, in *pb.LogData) (*pb.LogSuccess, error) {

	// fmt.Fprintf(s.fo, "%d: %s\n", in.Severity, in.Data) // xyzzy - add time stamp etc.
	// xyzzy - add output destination file etc.

	// No feature was found, return an unnamed feature
	return &pb.LogSuccess{Status: "success", Msg: ""}, nil
}

// LogMessage returns the feature at the given point. (PJS)
func (s *logItServer) LogMessage(ctx context.Context, in *pb.LogData) (*pb.LogSuccess, error) {

	fmt.Fprintf(s.fo, "%d: %s\n", in.Severity, in.Data) // xyzzy - add time stamp etc.
	// xyzzy - add output destination file etc.

	// No feature was found, return an unnamed feature
	return &pb.LogSuccess{Status: "success", Msg: ""}, nil
}

var invalidMode = errors.New("Invalid Mode")

func Fopen(fn string, mode string) (file *os.File, err error) {
	file = nil
	if mode == "r" {
		file, err = os.Open(fn) // For read access.
	} else if mode == "w" {
		file, err = os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	} else if mode == "a" {
		file, err = os.OpenFile(fn, os.O_RDWR|os.O_APPEND, 0660)
		if err != nil {
			file, err = os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		}
	} else {
		err = invalidMode
	}
	return
}

func newServer(fn string) *logItServer {
	s := new(logItServer)
	// PJS - xyzzy - if no path, then create the log path
	fo, err := Fopen(fn, "a")
	if err != nil {
		fmt.Printf("Error: (fatal) unable to open %s for append, %s\n", fn, err)
		os.Exit(1)
	}
	s.fo = fo
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			grpclog.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterLogItServer(grpcServer, newServer(*logFile))
	grpcServer.Serve(lis)
}
