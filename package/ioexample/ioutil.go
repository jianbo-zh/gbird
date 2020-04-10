package ioexample

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// ReadFile example
func ReadFile() {
	content, err := ioutil.ReadFile("testdata/hello")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)
}

// WriteFile example
func WriteFile() {
	message := []byte("Hello, Gophers!")
	err := ioutil.WriteFile("testdata/hello", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// ReadAll example
func ReadAll() {
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)
}

// ReadDir example
func ReadDir() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

// TempDir example
func TempDir() {
	content := []byte("temporary file's content")
	// 使用系统默认临时文件目录，并创建以example开头的随机命名的临时文件夹
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}
}

// TempDirSuffix example
func TempDirSuffix() {
	parentDir := os.TempDir()
	logsDir, err := ioutil.TempDir(parentDir, "*-logs")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(logsDir) // clean up

	// Logs can be cleaned out earlier if needed by searching
	// for all directories whose suffix ends in *-logs.
	globPattern := filepath.Join(parentDir, "*-logs")
	matches, err := filepath.Glob(globPattern)
	if err != nil {
		log.Fatalf("Failed to match %q: %v", globPattern, err)
	}

	for _, match := range matches {
		if err := os.RemoveAll(match); err != nil {
			log.Printf("Failed to remove %q: %v", match, err)
		}
	}
}

// TempFile example
func TempFile() {
	content := []byte("temporary file's content")
	// prefix: example
	// tmpfile, err := ioutil.TempFile("", "example")
	tmpfile, err := ioutil.TempFile("", "example.*.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		tmpfile.Close()
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
