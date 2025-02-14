package password

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	pass1 := "password1"
	pass2 := "password1"
	hashPaas1, _ := HashPassword(pass1)
	fmt.Println(VerifyPassword(hashPaas1, pass2))
}
