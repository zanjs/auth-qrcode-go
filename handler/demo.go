package handler

import (
	"city6/au/sign"
	"encoding/hex"
	"fmt"

	"github.com/kataras/iris"
)

// Sing is
type Sing struct {
	Sign string `json:"sign"`
	AKey string `json:"akey"`
}

// SingDemoT is
func SingDemoT(ctx iris.Context) {

	// sum := sign.MD5.Encrypt([]byte(`zanjs@55#`))
	// sumStr := hex.EncodeToString(sum)
	dax := &Sing{}
	dax.Sign = ctx.FormValue("sign")
	dax.AKey = ctx.FormValue("akey")

	if dax.AKey == "" || dax.Sign == "" {
		fmt.Println(dax)
		ResponseBad(ctx, "where is go")
		return
	}

	sum2 := sign.MD5.EncryptWithSalt([]byte(`zanjser@`), []byte(`skasjsssskkdaj`))
	sumStr2 := hex.EncodeToString(sum2)

	ResponseJSON(ctx, sumStr2)
}

func isSignAuth() {

}
