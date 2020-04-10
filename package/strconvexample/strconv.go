package strconvexample

import (
	"fmt"
	"strconv"
)

// Append example
func Append() {
	buf := []byte("Hello World")
	newBuf := strconv.AppendBool(buf, false)

	fmt.Printf("AppendBool: %s -> %s\n", buf, newBuf)
	fmt.Printf("AppendFloat: %s -> %s\n", buf, strconv.AppendFloat(buf, 100000.23456, 'g', 4, 32))
	fmt.Printf("AppendInt: %s -> %s\n", buf, strconv.AppendInt(buf, 123, 16))
	fmt.Printf("AppendUint: %s -> %s\n", buf, strconv.AppendUint(buf, 123, 8))

	fmt.Printf("AppendQuote: %s -> %s\n", buf, strconv.AppendQuote(buf, "T"))
	fmt.Printf("AppendQuoteToASCII: %s -> %s\n", buf, strconv.AppendQuoteToASCII(buf, "a"))

	fmt.Printf("AppendQuoteRune: %s -> %s\n", buf, strconv.AppendQuoteRune(buf, 'T'))
	fmt.Printf("AppendQuoteRuneToASCII: %s -> %s\n", buf, strconv.AppendQuoteRuneToASCII(buf, 'a'))

}

// Format example
func Format() {

	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatInt(123, 8))
	fmt.Println(strconv.FormatUint(123, 8))
	fmt.Println(strconv.FormatFloat(123.234, 'e', 2, 32))

	fmt.Println(strconv.Itoa(123))

}

// Other example
func Other() {

	// 可以用单反引号表示，且没有其它控制符，除制表符以外
	fmt.Println(strconv.CanBackquote("Hello' World"))

	fmt.Printf("%t\n", strconv.IsPrint('T'))
}

// Parse example
func Parse() {

	t, _ := strconv.ParseBool("T")
	fmt.Printf("%t\n", t)

	f, _ := strconv.ParseFloat("1.234", 32)
	fmt.Printf("%f\n", f)

	i, _ := strconv.ParseInt("1234", 16, 32)
	fmt.Printf("%d\n", i)

	ui, _ := strconv.ParseUint("1234", 16, 32)
	fmt.Printf("%d\n", ui)

	d, _ := strconv.Atoi("12321224")
	fmt.Printf("%d\n", d)
}

// Quote example
func Quote() {

	fmt.Println(strconv.Quote("Hello World"))
	fmt.Println(strconv.QuoteToASCII("Hello"))

	fmt.Println(strconv.QuoteRune('a'))
	fmt.Println(strconv.QuoteRuneToASCII('a'))

	s, _ := strconv.Unquote("\"Hello World\"")
	fmt.Println(s)
}
