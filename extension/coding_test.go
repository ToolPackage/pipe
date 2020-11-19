package extension

import (
	"fmt"
	"github.com/vipally/binary"
	"testing"
)

func TestCoding(t *testing.T) {
	out := binary.NewEncoder(0)
	out.Uint8(5)
	out.String("123456")
	bytes := out.Buffer()
	in := binary.NewDecoder(bytes)
	fmt.Println(in.Uint8())
	fmt.Println(in.String())
}
