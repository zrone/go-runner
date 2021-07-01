package types

type GitlabCrypt struct {
	CryptDataConfig
}

func (c *GitlabCrypt) Build(data CryptDataConfig) {
	if len(data.Headers["x-gitlab-token"]) > 0 {
		c.Token = data.Headers["x-gitlab-token"][0]
	}
}

func (c *GitlabCrypt) Compare() bool {
	return c.Sign == c.Token
}

func (c *GitlabCrypt) BuildPrefixCryptSign() {
	c.Sign = c.Project.Secret
}
