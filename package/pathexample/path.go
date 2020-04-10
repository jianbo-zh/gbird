package pathexample

import (
	"fmt"
	"path"
)

// Path example
func Path() {

	fmt.Println(path.Base(""))
	fmt.Println(path.Base("//"))
	fmt.Println(path.Base("/Hello/World"))

	paths := []string{
		"a/c",
		"a//c",
		"a/c/.",
		"a/c/b/..",
		"/../a/c",
		"/../a/b/../././/c",
	}
	for _, p := range paths {
		fmt.Printf("Clean(%q) = %q\n", p, path.Clean(p))
	}

	fmt.Printf("Dir(%q) = %q\n", "/a/b/c", path.Dir("/a/b/c"))

	fmt.Printf("Ext(%q) = %q\n", "/a/b/c/bar.css.txt", path.Ext("/a/b/c/bar.css.txt"))

	fmt.Printf("IsAbs(%q) = %t\n", "a/b/c/d", path.IsAbs("a/b/c/d"))
	fmt.Printf("IsAbs(%q) = %t\n", "/a/b/c/d", path.IsAbs("/a/b/c/d"))
	fmt.Printf("IsAbs(%q) = %t\n", "/a/b/../c/d", path.IsAbs("/a/b/../c/d"))

	fmt.Println(path.Join("a", "b", "c"))

	m, _ := path.Match("a/*/c", "a/b/c")
	fmt.Printf(`Match("a/*/c", "abc") = %t`+"\n", m)

	dir, file := path.Split("a/b/c/myfile.css")
	fmt.Printf("Split(%q), -> Dir:%q , File:%q", "a/b/c/myfile.css", dir, file)

}
