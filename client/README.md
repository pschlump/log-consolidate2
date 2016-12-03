Log Consolidator Client
=======================

This simple code is meant to consolidate logs.  I have an application that has bash, Go (golang), C++, C and Java
code with logs.   Most of the logs are written to files on multiple machines.  The files are replaced with named
pipes and this client then reads the named pipes, takes the data and used gRPC to send all of the log data to
a single server where it is written into a log file.

	$ client --file example.log &

to read from a named pipe.  The bash script can log directly with:

	$ client --msg "some message to log" --severity 2

The Go code directly makes calls to a *logAMessage* function.


