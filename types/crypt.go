package types

type AbstractCrypt interface {
	Build()
	Compare() bool
	BuildPrefixCryptSign()
	GetCryptDataConfig() CryptDataConfig
}

type CryptDataConfig struct {
	Symbol        string
	Message       Message
	Headers       map[string][]string
	Project       Project
	Token         string
	Sign          string
	Payload       string
	InternalDeloy InternalDeploy
}

type Project struct {
	Secret string
	Path   string
	Auth   Authentication
}
