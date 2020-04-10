package compressexample

import (
	"bytes"
	"compress/flate"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

var rawData string = "Hello World, Hello World, Hello World 123"

// 使用被压缩数据中的一些高频词，作为压缩字典，可以提高压缩率
var dict string = "Hello World , 123"

// Compress example
func Compress() {
	buf := new(bytes.Buffer)
	fw, err := flate.NewWriter(buf, flate.BestCompression)
	if err != nil {
		log.Fatal(err)
	}
	_, err = fw.Write([]byte(rawData))
	if err != nil {
		log.Fatal(err)
	}
	fw.Close()
	fmt.Printf("Compress Data: %x", buf.String())
}

// Uncompress example
func Uncompress() {
	buf := new(bytes.Buffer)
	compressHex := "f248cdc9c95708cf2fca49d151c0c55130343206040000ffff"

	cbytes, _ := hex.DecodeString(compressHex)

	freader := flate.NewReader(bytes.NewReader(cbytes))

	_, err := io.Copy(buf, freader)
	if err != nil {
		log.Fatal(err)
	}

	err = freader.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Raw Data: %s", buf)
}

// CompressWithDict example
func CompressWithDict() {
	buf := new(bytes.Buffer)
	fw, err := flate.NewWriterDict(buf, flate.BestCompression, []byte(dict))
	if err != nil {
		log.Fatal(err)
	}
	_, err = fw.Write([]byte(rawData))
	if err != nil {
		log.Fatal(err)
	}
	fw.Close()
	fmt.Printf("Compress Data: %x", buf.String())
}

// UncompressWithDict example
func UncompressWithDict() {

	buf := new(bytes.Buffer)

	compressHex := "4212d051c0c5012904040000ffff"

	cbytes, _ := hex.DecodeString(compressHex)

	freader := flate.NewReaderDict(bytes.NewReader(cbytes), []byte(dict))

	_, err := io.Copy(buf, freader)
	if err != nil {
		log.Fatal(err)
	}

	err = freader.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Raw Data: %s", buf)
}
