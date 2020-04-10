package bytesexample

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

// Buffer example
func Buffer() {
	buffer := bytes.NewBuffer([]byte("Hello"))
	fmt.Printf("Len: %d, Caption: %d\n", buffer.Len(), buffer.Cap())
	buffer.Write([]byte(" World World World World World"))
	// fmt.Printf("Len: %d, Caption: %d\n", buffer.Len(), buffer.Cap())
	buffer.WriteByte('W')
	buffer.WriteString("Yes Ok")
	// buffer.WriteTo(os.Stdout)
	// fmt.Printf("Len: %d, Caption: %d", buffer.Len(), buffer.Cap())

	// fmt.Printf("%s", buffer.Next(5))
	// fmt.Printf("%s", buffer.Next(5))

	// for buffer.Len() != 0 {
	// 	fmt.Printf("%s\n", buffer.Next(5))
	// }

	// fmt.Printf("Len: %d, Caption: %d", buffer.Len(), buffer.Cap())

	// b := make([]byte, 100)
	// buffer.Read(b)
	// fmt.Printf("%s", b)
	// fmt.Printf("%s", buffer)
	// fmt.Printf("Len: %d, Caption: %d", buffer.Len(), buffer.Cap())

	// a, _ := buffer.ReadByte()
	// fmt.Printf("%s", string(a))

	// for true {
	// 	if line, err := buffer.ReadBytes(' '); err == nil {
	// 		fmt.Printf("%s\n", line)
	// 	} else {
	// 		log.Print(string(line))
	// 		log.Print(err)
	// 		break
	// 	}
	// }
}

// ReadFrom test
func ReadFrom() {
	buffer := bytes.NewBufferString("Hello World")
	reader := strings.NewReader("Yes Ok!")
	buffer.ReadFrom(reader)
	buffer.WriteTo(os.Stdout)
}

// ReadRune example
func ReadRune() {
	buffer := bytes.NewBufferString("我是谁，我在哪？")

	r, s, err := buffer.ReadRune()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s, %d", string(r), s)
}

// Truncate example
func Truncate() {
	buffer := bytes.NewBufferString("Hello World")

	fmt.Printf("Len: %d, Caption: %d\n", buffer.Len(), buffer.Cap())
	// buffer.Truncate(0)
	buffer.Reset()
	fmt.Printf("Len: %d, Caption: %d\n", buffer.Len(), buffer.Cap())
}

// ReadByte example
func ReadByte() {
	buffer := bytes.NewBufferString("Hello World")
	fmt.Printf("Len: %d, Caption: %d\n", buffer.Len(), buffer.Cap())

	b, _ := buffer.ReadByte()
	fmt.Printf("ReadByte: %s\n", string(b))
	fmt.Printf("Len: %d, Caption: %d\n", buffer.Len(), buffer.Cap())

	err := buffer.UnreadByte()
	if err != nil {
		log.Fatal(err)
	}

	// err2 := buffer.UnreadByte()
	// if err2 != nil {
	// 	// bytes.Buffer: UnreadByte: previous operation was not a successful read
	// 	log.Fatal(err2)
	// }

	fmt.Printf("Len: %d, Caption: %d\n", buffer.Len(), buffer.Cap())
}

// Write example
func Write() {
	buffer := bytes.NewBufferString("")
	buffer.Write([]byte("Hello World"))
	buffer.WriteByte('T')
	buffer.WriteRune([]rune("波")[0])
	buffer.WriteString("World")
	buffer.WriteTo(os.Stdout)
}

// ToLower example
func ToLower() {
	rawStr := "Hello world"
	fmt.Printf("ToLower: %s -> %s\n", rawStr, bytes.ToLower([]byte(rawStr)))
	fmt.Printf("ToTitle: %s -> %s\n", rawStr, bytes.ToTitle([]byte(rawStr)))
	fmt.Printf("ToUpper: %s -> %s\n", rawStr, bytes.ToUpper([]byte(rawStr)))
	fmt.Printf("Trim: %s -> %s\n", rawStr, bytes.Trim([]byte(rawStr), "Hd"))
	fmt.Printf("Trim: %s -> %s\n", rawStr, bytes.TrimFunc([]byte(rawStr), func(r rune) bool {
		return true
	}))
}
