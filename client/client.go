// Package main implements a simple logging client using gRPC.
// Based on google's sample
// Copyright (C) Philip Schlump, 2015-2016.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	pb "github.com/pschlump/log-consolidate2/proto" // pb "google.golang.org/grpc/examples/chat/proto"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "testdata/ca.pem", "The file containning the CA root cert file")
	serverHostPort     = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
	logSeverity        = flag.Int("severity", 1, "The severity of the log message")
	logMsg             = flag.String("msg", "dummy-message", "The message to log")
	readFile           = flag.String("file", "", "The named pipe/file to read continuously if specified")
	verbose            = flag.Bool("verbose", false, "Show status and other messages")
)

func logAMessage(client pb.LogItClient, logSeverity int, logMsg string) {
	msg := &pb.LogData{Severity: int32(logSeverity), Data: logMsg}
	ok, err := client.LogMessage(context.Background(), msg)
	if err != nil {
		grpclog.Fatalf("%v.LogMessage(_) = _, %v: ", client, err)
	}
	if *verbose || ok.Status != "success" {
		grpclog.Println(ok)
	}
}

func readFIFO(client pb.LogItClient, fn string) {

	file, err := os.OpenFile(fn, os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// infinite loop
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			if *verbose {
				fmt.Printf("%s\n", line)
			}
			logAMessage(client, *logSeverity, string(line))
		}
	}

}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		var sn string
		if *serverHostOverride != "" {
			sn = *serverHostOverride
		}
		var creds credentials.TransportCredentials
		if *caFile != "" {
			var err error
			creds, err = credentials.NewClientTLSFromFile(*caFile, sn)
			if err != nil {
				grpclog.Fatalf("Failed to create TLS credentials %v", err)
			}
		} else {
			creds = credentials.NewClientTLSFromCert(nil, sn)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(*serverHostPort, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewLogItClient(conn)

	if *readFile == "" {
		logAMessage(client, *logSeverity, *logMsg)
	} else {
		readFIFO(client, *readFile)
	}

}

/* vim: set noai ts=4 sw=4: */
