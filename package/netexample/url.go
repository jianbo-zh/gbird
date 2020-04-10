package netexample

import (
	"fmt"
	"log"
	"net/url"
)

// URLParse example
func URLParse() {
	u, err := url.Parse("http://tp@bing.com:80/search?q=dotnet")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Hostname())
	fmt.Println(u.Port())
	fmt.Println(u.EscapedPath())
	fmt.Println(u.Query())
	fmt.Println(u.RequestURI())

	fmt.Println(u.IsAbs())

	u.Scheme = "https"
	u.Host = "google.com"
	q := u.Query()
	q.Set("q", "golang")
	fmt.Println(q.Encode())
	u.RawQuery = q.Encode()
	fmt.Println(u)
}

// URLValues example
func URLValues() {
	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")

	// "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
	fmt.Println(v.Encode())
	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"])
	fmt.Println(v)
}
