package archiveexample

import (
	"archive/zip"
	"bytes"
	"compress/flate"
	"fmt"
	"io"
	"log"
	"os"
)

// Zip example
func Zip() {
	buf := new(bytes.Buffer)

	zw := zip.NewWriter(buf)

	zw.RegisterCompressor(zip.Deflate, func(w io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(w, flate.BestCompression)
	})

	files := []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}

	for _, file := range files {
		f, err := zw.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}

		f.Write([]byte(file.Body))
	}

	err := zw.Close()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile("file.zip", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	buf.WriteTo(f)
}

// Unzip example
func Unzip() {
	r, err := zip.OpenReader("file.zip")
	if err != nil {
		log.Fatal(err)
	}

	r.RegisterDecompressor(zip.Deflate, func(r io.Reader) io.ReadCloser {
		return flate.NewReader(r)
	})

	defer r.Close()

	for _, f := range r.File {
		fmt.Printf("Contents of %s: \n", f.Name)

		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}

		io.CopyN(os.Stdout, rc, int64(f.UncompressedSize64))

		rc.Close()
		fmt.Println("\n ")
	}

}
