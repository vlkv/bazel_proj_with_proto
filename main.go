package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	pb "gitlab.com/myorg/myproj/proto"
)

func GetFoo() string {
	return "foo"
}

func main() {
	mp := &pb.MyProto{}
	fmt.Printf("Hello %s, also: %v", GetFoo(), proto.CompactTextString(mp))
}
