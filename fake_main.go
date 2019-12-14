package main

// import (
// 	"fmt"
//
// 	api "github.com/solo-io/protoc-gen/api"
// )
//
// func main() {
// 	nested := &api.Nested{
// 		DoubleRange: &api.DoubleRange{
// 			Start: 10,
// 			End:   12,
// 		},
// 	}
// 	nested2 := &api.Nested{
// 		DoubleRange: &api.DoubleRange{
// 			Start: 13,
// 			End:   14,
// 		},
// 	}
// 	nested3 := &api.Nested{
// 		DoubleRange: nil,
// 		Hello:       []string{"hello"},
// 	}
//
// 	hash1, err := nested.Hash(nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	hash2, err := nested2.Hash(nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	hash3, err := nested3.Hash(nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	hash4, err := nested3.Hash(nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(hash1)
// 	fmt.Println(hash2)
// 	fmt.Println(hash3)
// 	fmt.Println(hash4)
// }
