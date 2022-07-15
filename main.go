package main

import (
	"fmt"

	"example.com/m/v2/github.com/GrayFlash/proto_test"
	"google.golang.org/protobuf/types/known/anypb"
	// "github.com/GrayFlash/proto_test"
)

func main() {
	fmt.Println("Hello world!")
	db := &proto_test.Dashboard{}
	any, err := anypb.New(db)
	if err != nil {
		fmt.Println(err)
	}
	data := proto_test.Asset{Data: any}
	fmt.Println(data.Data)

	one := proto_test.AssetOne{Data: &proto_test.AssetOne_Dashboard{Dashboard: &proto_test.Dashboard{}}}
	fmt.Println(one.Data)
}
