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

func (c *GithubCrypt) Build(data CryptDataConfig) {
	if len(data.Headers["x-hub-signature-256"]) > 0 {
		c.Token = data.Headers["x-hub-signature-256"][0]
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
