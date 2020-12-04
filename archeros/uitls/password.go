package utils

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	NumStr  = "0123456789"
	CharStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SpecStr = "+=-@#~,.[]()!%^*$"
)

func GenPassword(length int) string {
	//随机种子
	rand.Seed(time.Now().UnixNano())

	//初始化密码切片
	var passwd []byte = make([]byte, length, length)
	//源字符串
	var sourceStr string

	sourceStr = fmt.Sprintf("%s%s%s", NumStr, CharStr, SpecStr)

	//遍历，生成一个随机index索引,
	for i := 0; i < length; i++ {
		index := rand.Intn(len(sourceStr))
		passwd[i] = sourceStr[index]
	}
	return string(passwd)
}
