PROTO	:= $(wildcard *.proto)
PBSRC	:= $(PROTO:.proto=.pb.go)

%.pb.go: %.proto
	protoc $< --go_out=plugins=grpc:.

all: $(PBSRC)
