Useful info on google protocol buffers
===================================================
About google protocol buffers: https://developers.google.com/protocol-buffers/

protobufs go installation:

1. Install the standard C++ implementation of protocol buffers from https://developers.google.com/protocol-buffers/

2. Of course, install the Go compiler and tools from https://golang.org/

3. Grab the code from the repository and install the proto package. The simplest way is to run:
 > go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
 
The compiler plugin, protoc-gen-go, will be installed in $GOBIN, defaulting to $GOPATH/bin.
It must be in your $PATH for the protocol compiler, protoc, to find it.