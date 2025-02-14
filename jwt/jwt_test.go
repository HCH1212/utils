package jwt

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test(t *testing.T) {
	var id uint
	id = uint(rand.Int())
	fmt.Println(id)
	token, refreshToken, err := GetToken(id)
	fmt.Println(token, refreshToken, err)
	res, err := ParseAccessToken(token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.ID, err)
	res, err = ParseRefreshToken(refreshToken)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.ID, err)
}
