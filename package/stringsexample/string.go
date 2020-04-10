package stringsexample

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

// JudgeBool example
func JudgeBool() {
	fmt.Println(strings.Contains("Hello World", "Wor"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	// fmt.Println(strings.ContainsRune("Hello World", ))
	fmt.Println(strings.EqualFold("Go", "go"))
	fmt.Println(strings.HasPrefix("Hello World", "Hello"))
	fmt.Println(strings.HasSuffix("Hello World", "rld"))
}

// StatCount example
func StatCount() {
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("five", ""))
}

// Search example
func Search() {
	fmt.Println(strings.Index("Hello World", "Worl"))
	fmt.Println(strings.IndexAny("chicken", "aeiouy"))
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))
	fmt.Println(strings.IndexByte("Helo", 'l'))

	// 单引号可以表示rane?
	fmt.Println(strings.IndexRune("chicken", 'k'))

	fmt.Println(strings.LastIndex("go gopher", "go"))
	fmt.Println(strings.LastIndexAny("go gopher", "oh"))

	f := func(c rune) bool {
		return c == 'm'
	}
	fmt.Println(strings.LastIndexFunc("go go m pher", f))

}

// Join example
func Join() {
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))

	fmt.Println("ba" + strings.Repeat("na", 2))

	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz波 ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))

	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))

	// ["a," "b," "c"]
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))

	// -l no limit
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", -1))

	fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))
	fmt.Printf("Fields are: %q\n", strings.FieldsFunc("  foo1;bar2,baz3...", func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}))
}

// Replace example
func Replace() {

	// Map
	f := func(r rune) rune {
		return r + 1
	}
	fmt.Println(strings.Map(f, "Hello"))

	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))

	fmt.Println(strings.ToLower("Gopher"))

	// 单词首字母大写
	fmt.Println(strings.Title("her royal highness"))
	// 所有单词都大写
	fmt.Println(strings.ToTitle("loud noises"))
	fmt.Println(strings.ToTitle("хлеб"))

	fmt.Println(strings.ToUpper("Gopher"))

}

// Truncate example
func Truncate() {
	fmt.Printf("%q\n", strings.Trim(" !!! Achtung! Achtung! !!! ", "! "))
	fmt.Printf("%q\n", strings.TrimFunc(" !!! Achtung! Achtung! !!! ", func(r rune) bool {
		return r == '!' || r == ' '
	}))

	fmt.Printf("%q\n", strings.TrimLeft(" !!! Achtung! Achtung! !!! ", "! "))
	fmt.Printf("%q\n", strings.TrimLeftFunc(" !!! Achtung! Achtung! !!! ", func(r rune) bool {
		return r == '!' || r == ' '
	}))

	fmt.Printf("%q\n", strings.TrimRight(" !!! Achtung! Achtung! !!! ", "! "))
	fmt.Printf("%q\n", strings.TrimRightFunc(" !!! Achtung! Achtung! !!! ", func(r rune) bool {
		return r == '!' || r == ' '
	}))

	fmt.Println(strings.TrimPrefix("Hello World", "Hel"))
	fmt.Println(strings.TrimSuffix("Hello, goodbye, etc!", "goodbye, etc!"))

	fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n"))

}

// Reader example
func Reader() {
	strr := strings.NewReader("我Hello World")
	fmt.Println(strr.Len())

	b := make([]byte, 20)
	// n, _ := strr.Read(b)
	// fmt.Printf("%d, %s\n", n, b[:n])

	n, _ := strr.ReadAt(b, 2)
	fmt.Printf("%d, %s\n", n, b[:n])

	// b, _ := strr.ReadByte()
	// fmt.Printf(string(b))
	// strr.UnreadByte()

	// r, i, _ := strr.ReadRune()
	// fmt.Printf("%s, %d", string(r), i)
	// strr.UnreadRune()

	// n, _ := strr.WriteTo(os.Stdout)
	// fmt.Println("\n", n)

	reprr := strings.NewReplacer("<", "&lt", ">", "&gt")
	fmt.Println(reprr.Replace("<Root>Helo</Root>"))

	reprr.WriteString(os.Stdout, "<root>")

}

// Builder example
func Builder() {
	var b strings.Builder
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&b, "%d...", i)
	}
	b.WriteString("ignition")
	fmt.Println(b.String())
}
