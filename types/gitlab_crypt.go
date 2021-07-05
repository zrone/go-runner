package types

type GitlabCrypt struct {
	CryptDataConfig
}

func DiscoverGitlabCrypt(cryptDataConfig CryptDataConfig) AbstractCrypt {
	// 默认 master 分支找不到 ref
	if cryptDataConfig.Message.Ref == "" {
		cryptDataConfig.Message.Ref = "refs/heads/master"
	}
	return &GitlabCrypt{
		cryptDataConfig,
	}
}

func (c *GitlabCrypt) Build() {
	if len(c.CryptDataConfig.Headers["x-gitlab-token"]) > 0 {
		c.Token = c.CryptDataConfig.Headers["x-gitlab-token"][0]
	}
}

func (c *GitlabCrypt) Compare() bool {
	return c.Sign == c.Token
}

func (c *GitlabCrypt) BuildPrefixCryptSign() {
	c.Sign = c.Project.Secret
}

func (c *GitlabCrypt) GetCryptDataConfig() CryptDataConfig {
	return c.CryptDataConfig
}
