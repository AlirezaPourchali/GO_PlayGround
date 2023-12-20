package main

import (
	"fmt"
	"hash/fnv"
)

func fnv32(y uint32) uint32 {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprint(y))) // only string input
	return h.Sum32()
}

func main() {
	//var input uint32 = 32
	//result := fnv32(input)
	//fmt.Printf("FNV hash of %v is: %d\n", input, result)
	//const MaxUint = ^uint32(0)
	fmt.Println(fnv32(899998))
}
