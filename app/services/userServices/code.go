package userServices

import (
	"context"
	"fmt"
	"math/rand"
	"time"
	"xinmiao/config/database/redis"
)

var ctx = context.Background()

func GenRandomCode(length int) string {
	const charset = "0123456789"

	// 创建本地的随机数生成器
	randSource := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(randSource)

	// 生成验证码
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[randGenerator.Intn(len(charset))]
	}
	return string(code)
}

func SetCode(key, value string) {
	// 将验证码放到redis中，以ID+使用场景为键，以code为值，有效期5分钟
	redis.RedisClient.Set(ctx, key, value, 5*time.Minute)
}

func GetCode(key string) (string, error) {
	// 获取redis中的验证码
	val, err := redis.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func DelCode(key string) {
	// 删除redis中的验证码
	redis.RedisClient.Del(ctx, key)
}

func SendCodeByMail(ID, mailto, category string) error {
	code := GenRandomCode(4)
	subject := fmt.Sprintf("%s验证码", category)
	body := fmt.Sprintf("您的验证码为：%s，5分钟内有效，请尽快使用。请勿泄漏他人。", code)
	err := SendMail(mailto, subject, body)
	if err != nil {
		return err
	}
	// 将验证码放到redis中
	SetCode(fmt.Sprintf("%s-%s", ID, category), code)
	return nil
}

func CheckCode(ID, category, code string) (bool, error) {
	rightCode, err := GetCode(fmt.Sprintf("%s-%s", ID, category))
	if err != nil {
		return false, err
	}
	if rightCode == code {
		SetCode(fmt.Sprintf("%s-%s", ID, category), "Accepted")
		return true, nil
	} else if rightCode == "Accepted" {
		return true, nil
	} else if rightCode != code {
		return false, nil
	}
	return false, nil
}
