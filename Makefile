#
# by PJS
#
# Notes:
# 		Download and install recompiled https://github.com/google/protobuf/releases
#		Placed in ~/go/src/www*/gRPC 
#		This gets you protoc
#

all: ./proto/log_it.pb.go
	( cd server ; go build )
	( cd client ; go build )

# From: http://www.grpc.io/docs/tutorials/basic/go.html
# Generates log_it.pb.go
./proto/log_it.pb.go: ./proto/log_it.proto
	 protoc -I proto/ proto/log_it.proto --go_out=plugins=grpc:proto

run_server:
	./server/server & 

clean:
	rm client/client
	rm server/server

