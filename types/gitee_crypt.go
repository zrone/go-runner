package types

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

type GiteeCrypt struct {
	CryptDataConfig
	Timestamp string
}

func DiscoverGiteeCrypt(cryptDataConfig CryptDataConfig) AbstractCrypt {
	return &GiteeCrypt{
		CryptDataConfig: cryptDataConfig,
	}
}

func (c *GiteeCrypt) Build() {
	if len(c.CryptDataConfig.Headers["X-Gitee-Token"]) > 0 {
		c.Token = c.CryptDataConfig.Headers["X-Gitee-Token"][0]
	}

	if len(c.CryptDataConfig.Headers["X-Gitee-Timestamp"]) > 0 {
		c.Timestamp = c.CryptDataConfig.Headers["X-Gitee-Timestamp"][0]
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
	c.Sign = base64.StdEncoding.EncodeToString(m.Sum(nil))
}

func (c *GiteeCrypt) GetCryptDataConfig() CryptDataConfig {
	return c.CryptDataConfig
}
