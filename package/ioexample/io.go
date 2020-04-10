package ioexample

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// ReadFull example
func ReadFull() {
	s := make([]byte, 10)
	str := "Hello World"
	n, err := io.ReadFull(strings.NewReader(str), s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read Len: %d, Content: %s", n, s)
}

// Pipe example
func Pipe() {
	r, w := io.Pipe()
	go func() {
		fmt.Fprint(w, "some text to be read\n")
		w.Close()

	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	fmt.Print(buf.String())
}

// SectionReader example
func SectionReader() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 5, 16)

	if _, err := s.Seek(10, io.SeekStart); err != nil {
		log.Fatal(err)
	}

	fmt.Println(s.Size())

	buf := make([]byte, 6)
	if _, err := s.Read(buf); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", buf)
}

// LimitReader example
func LimitReader() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := io.LimitReader(r, 4)

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}

// MultiReader example
func MultiReader() {
	r1 := strings.NewReader("first reader ")
	r2 := strings.NewReader("second reader ")
	r3 := strings.NewReader("third reader\n")
	r := io.MultiReader(r1, r2, r3)

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

// TeeReader example
func TeeReader() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	var buf bytes.Buffer
	tee := io.TeeReader(r, &buf)

	printall := func(r io.Reader) {
		b, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", b)
	}

	printall(tee)
	printall(&buf)
}
