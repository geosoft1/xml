package xml

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

type A struct {
	XMLName string `xml:"xml"`
	A       int    `xml:"a"`
	B       string `xml:"b"`
}

var a A

func Test(t *testing.T) {
	x := "<xml><a>1</a><b>a</b></xml>"
	Decode(strings.NewReader(x), &a)
	fmt.Println(a)
	Encode(os.Stdout, &a)
}
