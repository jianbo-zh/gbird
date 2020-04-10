package encodeexample

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
)

// BASE64Encode example
func BASE64Encode() {

	bi := base64.StdEncoding.EncodedLen(11)
	fmt.Printf("EncodedLen: 11 => %d", bi)

	// 标准的base64和urlbase64区别：标准的base64用 "-"和"_" 替换了 "+"和"/" 字符
	bstrraw := base64.RawStdEncoding.EncodeToString([]byte("Hello World"))
	fmt.Printf("RawStdEncoding: Hello World! => %s \n", bstrraw)

	bstr := base64.StdEncoding.EncodeToString([]byte("Hello World"))
	fmt.Printf("StdEncoding: Hello World! => %s \n", bstr)

	burlraw := base64.RawURLEncoding.EncodeToString([]byte("https://www.baidu.com/path/to?q=1&a=b"))
	fmt.Printf("RawStdEncoding: https://www.baidu.com/path/to?q=1&a=b => %s \n", burlraw)

	burl := base64.URLEncoding.EncodeToString([]byte("https://www.baidu.com/path/to?q=1&a=b"))
	fmt.Printf("StdEncoding: https://www.baidu.com/path/to?q=1&a=b => %s \n", burl)

	rawStr := "Hello World"
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(rawStr)))
	base64.StdEncoding.Encode(buf, []byte(rawStr))
	fmt.Printf("%s", buf)

	buf2 := make([]byte, len(buf))
	n, _ := base64.StdEncoding.Decode(buf2, buf)
	fmt.Printf("%s-%d \n", buf2[:n], n)
}

// BASE64Encoder example
func BASE64Encoder() {

	// 对输出流进行编码
	stdoutEncoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	stdoutEncoder.Write([]byte("Hello World"))
	stdoutEncoder.Close()
}

// BASE64Decoder example
func BASE64Decoder() {

	rawStr := "Hello World"
	base64Str := base64.StdEncoding.EncodeToString([]byte(rawStr))

	bStr := strings.NewReader(base64Str)
	strReader := base64.NewDecoder(base64.StdEncoding, bStr)
	rBuf := make([]byte, base64.StdEncoding.DecodedLen(len(base64Str)))
	n, err := strReader.Read(rBuf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s-%d \n", rBuf[:n], n)
}
