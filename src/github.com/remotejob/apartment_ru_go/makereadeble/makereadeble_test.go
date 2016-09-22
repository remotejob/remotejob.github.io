package makereadeble

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/remotejob/apartment_ru_go/mgenerator"
)

func TestMakereadeble(t *testing.T) {

	buf := bytes.NewBuffer(nil)

	f, _ := os.Open("/tmp/blog.txt")

	io.Copy(buf, f) // Error handling elided for brevity.
	f.Close()

	mtext0 := mgenerator.Generate(buf.Bytes())

	// fmt.Println(mtext0)
	// fmt.Println("---------------")
	// mtext1 := mgenerator.Generate(buf.Bytes())

	// fmt.Println(mtext1)

	title, contents, mcontents := Makehuman(mtext0)

	fmt.Println(title)
	fmt.Println("----------")
	fmt.Println(contents)

	fmt.Println("-----------")
	fmt.Println(mcontents)
}
