package archiveexample

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

// Tar example
func Tar() {
	buf := new(bytes.Buffer)

	tw := tar.NewWriter(buf)

	files := []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence."},
	}

	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Size: int64(len(file.Body)),
		}

		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}

		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatal(err)
		}
	}

	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile("file.tar", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	buf.WriteTo(f)
	f.Close()
}

// Untar example
func Untar() {

	file, err := os.OpenFile("file.tar", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	r := tar.NewReader(file)

	for {
		hdr, err := r.Next()
		if err == io.EOF {
			break
		}

		fmt.Printf("Contents of %s:\n", hdr.Name)

		if _, err := io.Copy(os.Stdout, r); err != nil {
			log.Fatal(err)
		}

		fmt.Println("\n ")
	}
}
