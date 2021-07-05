package types

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type GithubCrypt struct {
	CryptDataConfig
}

func DiscoverGithubCrypt(cryptDataConfig CryptDataConfig) AbstractCrypt {
	return &GithubCrypt{
		cryptDataConfig,
	}
}

func (c *GithubCrypt) Build() {
	if len(c.CryptDataConfig.Headers["x-hub-signature-256"]) > 0 {
		c.Token = c.CryptDataConfig.Headers["x-hub-signature-256"][0]
	}
}

func (c *GithubCrypt) Compare() bool {
	return c.Sign == c.Token
}

func (c *GithubCrypt) BuildPrefixCryptSign() {
	m := hmac.New(sha256.New, []byte(c.Project.Secret))
	m.Write([]byte(c.Payload))
	c.Sign = fmt.Sprintf("sha256=%s", hex.EncodeToString(m.Sum(nil)))
}

func (c *GithubCrypt) GetCryptDataConfig() CryptDataConfig {
	return c.CryptDataConfig
}
