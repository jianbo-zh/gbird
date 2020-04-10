package encodeexample

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// HEXEncoder example
func HEXEncoder() {
	s := "Hello world"
	w := hex.NewEncoder(os.Stdout)
	w.Write([]byte(s))
}

// HEXDecoder example
func HEXDecoder() {
	s := "48656c6c6f20776f726c64"
	buf := make([]byte, len(s))
	r := strings.NewReader(s)
	reader := hex.NewDecoder(r)
	n, _ := reader.Read(buf)
	fmt.Printf("%s", buf[:n])
}

// HEXEncode example
func HEXEncode() {
	s := "Hello world"
	fmt.Printf("%d EncodeLen: %d\n", len(s), hex.EncodedLen(len(s)))
	fmt.Printf("%s EncodeToString: %s\n", s, hex.EncodeToString([]byte(s)))

	buf := make([]byte, hex.EncodedLen(len(s)))
	n := hex.Encode(buf, []byte(s))
	fmt.Printf("%s Encode: %s\n", s, buf[:n])
}

// HEXDecode example
func HEXDecode() {
	s := "48656c6c6f20776f726c64"
	fmt.Printf("%d DecodedLen: %d\n", len(s), hex.DecodedLen(len(s)))

	d, _ := hex.DecodeString(s)
	fmt.Printf("%s DecodeString: %s\n", s, d)
}

// HEXDump example
func HEXDump() {
	s := "Hello World"
	fmt.Printf("%s HEXDump: \n%s", s, hex.Dump([]byte(s)))

	// Dumper
	wc := hex.Dumper(os.Stdout)
	wc.Write([]byte(s))
	wc.Close()
}
