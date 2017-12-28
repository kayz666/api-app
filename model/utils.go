package model

import (
	"github.com/satori/go.uuid"
	"github.com/kayz666/api-app/utils"
)



func StringToPassword(p string) string{
	return utils.StrToMD5v1(p)
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