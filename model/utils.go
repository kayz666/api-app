package model


import (
	"github.com/satori/go.uuid"
	"github.com/kayz666/api-app/utils"
	"fmt"

)
func RandomAccount()string{
	return fmt.Sprintf("user_%s%s","10",uuid.NewV1().String()[:6])
}


func MakePasswd(p string) string{
	return utils.Encrypt(p,"AES_128")
}

func RandomUUID() string{
	return uuid.NewV4().String()
}

func RandomAPIKey() string{
	return utils.StrToMD5v2(uuid.NewV1().String()[:20])
}

func RandomSN() string{
	return uuid.NewV4().String()
}