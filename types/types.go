package types

type Config struct {
	Port       string
	MysqlDNS   string
	TimeZone   string
	RedisDNS   string
	WorkNumber int
	QueueDNS   string
	QueueDb    int
}

// 自动化部署配置
type InternalDeploy struct {
	Symbol string
	Secret string
	Path   string
	Auth   Authentication `gorm:"embedded"`
}

type Authentication struct {
	Scheme int64
	User   string
	Host   string
	Port   int64
}

type Request struct {
}

type Response struct {
	Code    int64                  `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

func (t *InternalDeploy) TableName() string {
	return "internal_deploy"
}
