package types

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

type GiteeCrypt struct {
	CryptDataConfig
	Timestamp string
}

func (c *GiteeCrypt) Build(data CryptDataConfig) {
	if len(data.Headers["X-Gitee-Token"]) > 0 {
		c.Token = data.Headers["X-Gitee-Token"][0]
	}

	if len(data.Headers["X-Gitee-Timestamp"]) > 0 {
		c.Timestamp = data.Headers["X-Gitee-Timestamp"][0]
	}
}

func (c *GiteeCrypt) Compare() bool {
	return c.Sign == c.Token
}

func (c *GiteeCrypt) BuildPrefixCryptSign() {
	prefixCryptString := []byte(fmt.Sprintf(`%s
%s`, c.Timestamp, c.Project.Secret))
	m := hmac.New(sha256.New, []byte(c.Project.Secret))
	m.Write(prefixCryptString)
	c.Sign = base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(m.Sum(nil))))
	fmt.Println(c.Sign)
}
