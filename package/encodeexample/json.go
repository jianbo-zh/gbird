package encodeexample

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type website struct {
	Name    string `json:"name"`
	URL     string `json:"url,omitempty"`
	Comment string `json:"-"`
}

// JSONMarshal example
func JSONMarshal() {
	bi, _ := json.Marshal(123)
	fmt.Printf("Marshal Int: %s \n", string(bi))

	bs, _ := json.Marshal("Hello World!")
	fmt.Printf("Marshal String: %s \n", string(bs))

	barr, _ := json.Marshal([...]int{1, 2, 3})
	fmt.Printf("Marshal String: %s \n", string(barr))

	bmap, _ := json.Marshal(map[string]string{"key": "value"})
	fmt.Printf("Marshal String: %s \n", string(bmap))

	bstruct, _ := json.MarshalIndent([]website{
		{"baidu", "https://www.baidu.com", "xxx"},
		{"google", "https://www.google.com", "xxx"},
		{"sina", "", "xxx"},
	}, "", "  ")
	fmt.Printf("Marshal String: %s \n", string(bstruct))
}

// JSONUnMarshal example
func JSONUnMarshal() {

	var bi int
	var bs string
	var barr []int
	var bmap map[string]string
	var bstruct website

	json.Unmarshal([]byte("123"), &bi)
	fmt.Printf("%v \n", bi)

	// 序列化后的字符串是由双引号的
	json.Unmarshal([]byte("\"Hello World\""), &bs)
	fmt.Printf("%v \n", bs)

	json.Unmarshal([]byte("[1, 2, 3]"), &barr)
	fmt.Printf("%v \n", barr)

	json.Unmarshal([]byte(`{"key": "value"}`), &bmap)
	fmt.Printf("%v \n", bmap)

	err := json.Unmarshal([]byte(`{
		"name":"baidu",
		"url":"https://www.baidu.com"
	}`), &bstruct)

	if err != nil {
		log.Fatal("can not unmarshal")
	}

	fmt.Printf("%+v", bstruct)
}

// JSONEncode example
func JSONEncode() {

	var buf bytes.Buffer

	json.NewEncoder(&buf).Encode(123)
	buf.WriteTo(os.Stdout)

	json.NewEncoder(&buf).Encode("Hello World!")
	buf.WriteTo(os.Stdout)

	json.NewEncoder(&buf).Encode([...]int{1, 2, 3})
	buf.WriteTo(os.Stdout)

	json.NewEncoder(&buf).Encode(map[string]string{"key": "value"})
	buf.WriteTo(os.Stdout)

	json.NewEncoder(&buf).Encode(website{"baidu", "https://www.baidu.com", "comments"})
	buf.WriteTo(os.Stdout)
}

// JSONDecode example
func JSONDecode() {

	var (
		bi      int
		bs      string
		barr    []int
		bmap    map[string]string
		bstruct website
	)

	json.NewDecoder(strings.NewReader("123")).Decode(&bi)
	fmt.Printf("%v \n", bi)

	json.NewDecoder(strings.NewReader(`"Hello World"`)).Decode(&bs)
	fmt.Printf("%v \n", bs)

	json.NewDecoder(strings.NewReader("[1, 2, 3]")).Decode(&barr)
	fmt.Printf("%v \n", barr)

	json.NewDecoder(strings.NewReader(`{"key": "value"}`)).Decode(&bmap)
	fmt.Printf("%v \n", bmap)

	json.NewDecoder(strings.NewReader(`
	{
		"name": "baidu", 
		"url": "https://www.baidu.com",
		"comment": "comment ..."
	}`)).Decode(&bstruct)
	fmt.Printf("%+v \n", bstruct)
}
