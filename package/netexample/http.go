package netexample

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

// SimpleClient example
func SimpleClient() {
	resp, err := http.Get("http://www.baidu.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", string(body))
}

// CustomClient example
func CustomClient() {
	trans := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Transport: trans,
		Timeout:   30 * time.Second,
	}

	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", string(body))
}

type countHandler struct {
	mu sync.Mutex // guards n
	n  int
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

// SimpleServer example
func SimpleServer() {

	http.Handle("/foo", new(countHandler))

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	// func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// CustomServer example
func CustomServer() {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        new(countHandler),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

// FileServer example
func FileServer() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("D:\\Projects\\gbird"))))
}
