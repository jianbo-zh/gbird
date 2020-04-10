package compressexample

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"strings"
)

// Gzip example
func Gzip() {
	rawData := "Hello World Hello World"
	buf := new(bytes.Buffer)
	// gw := gzip.NewWriter(buf)
	gw, _ := gzip.NewWriterLevel(buf, gzip.BestCompression)

	_, err := gw.Write([]byte(rawData))
	if err != nil {
		log.Fatal(err)
	}
	gw.Close()

	fmt.Printf("%x", buf.Bytes())
}

// UnGzip example
func UnGzip() {
	gzipDataHex := "1f8b08000000000002fff248cdc9c95708cf2fca4951406203020000ffff8a50f98117000000"
	gr, err := gzip.NewReader(hex.NewDecoder(strings.NewReader(gzipDataHex)))
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, gr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", buf.String())
}
