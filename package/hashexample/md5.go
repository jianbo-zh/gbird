package hashexample

import (
	"crypto/md5"
	"fmt"
	"io"
)

// Md5 example
func Md5() {
	h := md5.New()
	io.WriteString(h, "The fog is getting thicker!")
	io.WriteString(h, "And Leon's getting laaarger!")
	fmt.Printf("%x\n", h.Sum(nil))

	fmt.Printf("md5.Sum: %s -> %x", "Hello World", md5.Sum([]byte("Hello World")))
}
