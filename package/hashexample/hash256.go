package hashexample

import (
	"crypto/sha256"
	"fmt"
)

// Hash256 example
func Hash256() {

	h := sha256.New()
	h.Write([]byte("Hello "))
	h.Write([]byte("World"))

	fmt.Printf("%x\n", h.Sum(nil))

	fmt.Printf("%x\n", sha256.Sum256([]byte("Hello World")))
}
