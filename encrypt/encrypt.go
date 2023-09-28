package encrypt

import (
	`crypto/md5`
	`crypto/sha1`
	`crypto/sha256`
	`encoding/base64`
	`fmt`
	`net/url`
	
	`github.com/google/uuid`
)

// Md5 加密字符串
//  @param value string 要加密的字符串
//  @return string  加密后的字符串
func Md5(value string) string {
	h := md5.Sum([]byte(value))
	md5String := fmt.Sprintf("%x", h)
	return md5String
}

// Password 加密密码
//  password 要加密的密码
//  passwords 加密后的密码
func Password(password string) (passwords string) {
	passwords = Md5(url.QueryEscape(base64.StdEncoding.EncodeToString([]byte(password))))
	return
}

// PasswdSHA1 SHA1密码加密
func PasswdSHA1(s string) (p string) {
	o := sha1.New()
	o.Write([]byte(s))
	p = base64.StdEncoding.EncodeToString(o.Sum(nil))
	return
}

// PasswdSHA256 SHA256密码加密
func PasswdSHA256(s string) (p string) {
	o := sha256.New()
	o.Write([]byte(s))
	p = base64.StdEncoding.EncodeToString(o.Sum(nil))
	return
}

// UUID 生成UUID
//  @return string uuid 字符串
func UUID() string {
	return uuid.New().String()
}
