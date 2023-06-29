package main

import (
	"github.com/davecgh/go-spew/spew"
	_ "github.com/leetm4n/rpc-error-proto-mapper-go/api/proto/rpc/errormapper/v1"
	_ "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func main() {

	pb := &anypb.Any{
		TypeUrl: "type.googleapis.com/google.rpc.BadRequest",
		Value:   []byte{},
	}

	// pb, err := anypb.New(&errdetails.BadRequest{})
	// if err != nil {
	// 	panic(err)
	// }

	spew.Dump(pb, pb.Value)

	dst, err := anypb.UnmarshalNew(pb, proto.UnmarshalOptions{})

	if err != nil {
		panic(err)
	}

	spew.Dump(dst)

}
