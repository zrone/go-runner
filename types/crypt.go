package types

type AbstractCrypt interface {
	Build(data CryptDataConfig)
	Compare() bool
	BuildPrefixCryptSign()
}

type CryptDataConfig struct {
	Symbol  string
	Message Message
	Headers map[string][]string
	Project Project
	Token   string
	Sign    string
	Payload string
}

type Project struct {
	Secret string
	Path   string
}

func HandleBuildPrefixCryptSign(c AbstractCrypt) {
	c.BuildPrefixCryptSign()
}

func HandleBuild(c AbstractCrypt, data CryptDataConfig) {
	c.Build(data)
}

func HandleCompare(c AbstractCrypt) bool {
	return c.Compare()
}
