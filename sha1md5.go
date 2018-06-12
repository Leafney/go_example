package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

// 对字符串进行MD5哈希
func md5Str(data string) string {
	m := md5.New()
	m.Write([]byte(data))
	my_md5 := m.Sum(nil)
	return fmt.Sprintf("%x", my_md5)
}

// 对字符串进行MD5哈希
func md5Str2(data string) string {
	my_md5 := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", my_md5)
}

// 对字符串进行SHA1哈希
func sha1Str(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	my_sha1 := h.Sum(nil)
	// 使用`%x`来将散列结果格式化为16进制的字符串
	return fmt.Sprintf("%x", my_sha1)
}

func hmacsha1Str(key string, data string) string {
	h_key := []byte(key)
	mac := hmac.New(sha1.New, h_key)
	mac.Write([]byte(data))
	my_mac := mac.Sum(nil)
	// 使用`%x`来将散列结果格式化为16进制的字符串
	return fmt.Sprintf("%x", my_mac)
}

func main() {
	s := "hello world"
	fmt.Println(md5Str(s))
	fmt.Println(md5Str2(s))
	fmt.Println(sha1Str(s))
	fmt.Println(hmacsha1Str("abc", s))

	// 5eb63bbbe01eeed093cb22bb8f5acdc3
	// 5eb63bbbe01eeed093cb22bb8f5acdc3
	// 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
	// faf32544e39b2c626bd8c17cd6c54d79ba86d8a0
}
