package main

import "fmt"

func reverseBits(num uint32) uint32 {
	var res uint32 = 0
	for i := 0; i < 32; i++ {
		bit := (num >> i) & 1
		res += (bit << (31 - i))
	}
	return res
}

func main() {
	var num uint32 = 43261596
	fmt.Println(reverseBits(num))
}

// output: 964176192
