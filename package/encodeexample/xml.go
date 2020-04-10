package encodeexample

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

// XMLMarshal example
func XMLMarshal() {
	type Address struct {
		City, State string
	}
	type Person struct {
		XMLName   xml.Name `xml:"person"`
		ID        int      `xml:"id,attr"`
		FirstName string   `xml:"name>first"`
		LastName  string   `xml:"name>last"`
		Age       int      `xml:"age"`
		Height    float32  `xml:"height,omitempty"`
		Married   bool
		Address
		Comment string `xml:",comment"`
	}

	v := &Person{ID: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}

	output, err := xml.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
}

// XMLUnmarshal example
func XMLUnmarshal() {
	type Email struct {
		Where string `xml:"where,attr"`
		Addr  string
	}
	type Address struct {
		City, State string
	}
	type Result struct {
		XMLName xml.Name `xml:"Person"`
		Name    string   `xml:"FullName"`
		Phone   string
		Email   []Email  `xml:"email"`
		Groups  []string `xml:"Group>Value"`
		Address
	}
	v := Result{Name: "none", Phone: "none"}

	data := `
		<Person>
			<FullName>Grace R. Emlin</FullName>
			<Company>Example Inc.</Company>
			<email where="home">
				<Addr>gre@example.com</Addr>
			</email>
			<email where='work'>
				<Addr>gre@work.com</Addr>
			</email>
			<Group>
				<Value>Friends</Value>
				<Value>Squash</Value>
			</Group>
			<City>Hanga Roa</City>
			<State>Easter Island</State>
		</Person>
	`
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("XMLName: %#v\n", v.XMLName)
	fmt.Printf("Name: %q\n", v.Name)
	fmt.Printf("Phone: %q\n", v.Phone)
	fmt.Printf("Email: %v\n", v.Email)
	fmt.Printf("Groups: %v\n", v.Groups)
	fmt.Printf("Address: %v\n", v.Address)
}

// XMLDecoder example
func XMLDecoder() {
	input := `
	<Person>
		<FirstName>FN</FirstName>
		<LastName>LN</LastName>
	</Person>`
	inputReader := strings.NewReader(input)

	decoder := xml.NewDecoder(inputReader)
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		// 开始标签
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name: %s\n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
			}
		// 结束标签
		case xml.EndElement:
			fmt.Printf("Token of '%s' end\n", token.Name.Local)
		// 文本内容
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
		default:
			// ...
		}
	}
}
