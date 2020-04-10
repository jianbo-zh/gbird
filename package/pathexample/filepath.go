package pathexample

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// Filepath example
func Filepath() {

	a, _ := filepath.Abs("a/b/c/d")
	fmt.Println(a)

	fmt.Println(filepath.FromSlash("a/b/c"))

	// 匹配 D:\Projects\gbird 目录下不以 ".", "f", "g", "h"开头的文件
	list, err := filepath.Glob(`D:\Projects\gbird\[^.fgh]*`)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range list {
		fmt.Println(v)
	}

	fmt.Printf("%t\n", filepath.HasPrefix("/a/b/c", "/a/b"))

	// windows => true，filepath 会考虑文件系统
	fmt.Printf("%t\n", filepath.IsAbs(`D:\a\b\c\d`))

	// windows => false
	fmt.Printf("%t\n", path.IsAbs(`D:\a\b\c\d`))

	fmt.Printf("%s\n", filepath.Join("a", "b", "c"))

	paths := []string{
		"/a/b/c",
		"/b/c",
		"./b/c",
	}
	base := "/a"

	fmt.Println("On Windows:")
	for _, targetpath := range paths {
		relatepath, err := filepath.Rel(base, targetpath)
		fmt.Printf("%q: %q %v\n", targetpath, relatepath, err)
	}

	fmt.Println("On Unix:", filepath.SplitList("/a/b/c:/usr/bin"))
	fmt.Println("On Windows:", filepath.SplitList("/a/b/c;/usr/bin"))

	// replacing each separator character in path with a slash ('/')
	fmt.Println(filepath.ToSlash(`D:\a\b\c\d`))

	fmt.Println(filepath.VolumeName(`D:\a\b\c\d`))

	// 查找所有包括txt文件的目录，最多10个
	n := 0
	filepath.Walk("D:\\", func(path string, info os.FileInfo, err error) error {
		if n >= 10 {
			return filepath.SkipDir
		}
		mbool, err := filepath.Match("*.txt", info.Name())
		if mbool == true {
			fmt.Println(filepath.Dir(path))
			n++
			return filepath.SkipDir
		}

		return err
	})
}
