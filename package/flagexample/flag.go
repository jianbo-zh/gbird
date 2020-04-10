package flagexample

import (
	"flag"
	"fmt"
	"net/url"
)

type urlValue struct {
	URL *url.URL
}

func (v urlValue) String() string {
	if v.URL != nil {
		return v.URL.String()
	}
	return ""
}

func (v urlValue) Set(s string) error {
	if u, err := url.Parse(s); err != nil {
		return err
	} else {
		*v.URL = *u
	}
	return nil
}

// FlagSetVar1 example
func FlagSetVar1() {
	mdDir := flag.String("src-dir", "./", "markdown file root dir")
	htmlDir := flag.String("des-dir", "./htmldir", "markdown generate html dir")
	flag.Parse()
	fmt.Println(*mdDir, *htmlDir)

}

// FlagSetVar2 example
func FlagSetVar2() {
	var mdDir, htmlDir string
	flag.StringVar(&mdDir, "src-dir", "./", "markdown file root dir")
	flag.StringVar(&htmlDir, "des-dir", "./htmldir", "markdown generate html dir")
	flag.Parse()
	fmt.Println(mdDir, htmlDir)
}

// FlagSetVar3 example
func FlagSetVar3() {
	var u = &url.URL{}
	fs := flag.NewFlagSet("ExampleValue", flag.ExitOnError)
	fs.Var(&urlValue{u}, "url", "URL to parse")

	fs.Parse([]string{"-url", "https://golang.org/pkg/flag/"})
	fmt.Printf(`{scheme: %q, host: %q, path: %q}`, u.Scheme, u.Host, u.Path)
}
