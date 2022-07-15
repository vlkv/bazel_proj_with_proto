package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	pb "github.com/vlkv/bazel_proj_with_proto/proto"
)

func GetFoo() string {
	return "foo"
}

func main() {
	mp := &pb.MyProto{}
	fmt.Printf("Hello %s, also: %v", GetFoo(), proto.CompactTextString(mp))
}
