package main

import "fmt"

// Hello struct
type Hello struct {
	it int
	na string
}

func main() {
	amap := make(map[*Hello]bool)

	hel := &Hello{1, "me"}
	amap[hel] = true

	fmt.Printf("%v", amap)
}
