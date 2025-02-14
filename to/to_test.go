package to

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	a := "hello world"
	b := "true"
	c := 123456
	d := 1.6653
	e := "1.625"
	fmt.Println(ToBool(b))
	fmt.Println(ToBytes(a))
	fmt.Println(ToString(c))
	fmt.Println(ToInt(d))
	fmt.Println(ToFloat(e))
	fmt.Println(ToJson(a))
}
