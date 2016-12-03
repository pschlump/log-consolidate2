# Description

The log server and client demonstrate how to use grpc go libraries to
perform log consolidation.

Please refer to [gRPC Basics: Go] (http://www.grpc.io/docs/tutorials/basic/go.html) for more information.

See the definition of the route guide service in proto/route_guide.proto.

# Run the sample code

To compile and run the server, assuming you are in the root of the route_guide
folder, i.e., .../examples/route_guide/, simply:

```sh
$ make
```

Will will compile both client and server.

```sh
$ ./server/server &
```

to star the server, then

```sh
$ ./client/client --msg "this is a message that gets logged"
```

to send a simple message.

# Optional command line flags

Read the README.md in ./client

