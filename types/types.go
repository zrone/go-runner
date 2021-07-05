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
	Uuid   string
	Secret string
	Path   string         // template path
	Auth   Authentication `gorm:"embedded"`
}

type Authentication struct {
	Scheme int64
	User   string
	Host   string
	Port   int64
	Pwd    string
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

type RunnerCi struct {
	BeforeScript []string `json:"before_script,omitempty"`
	Script       []string `json:"script,omitempty"`
	AfterScript  []string `json:"after_script,omitempty"`
	Webpath      string   `json:"webpath,omitempty"`
	Only         []string `json:"only,omitempty"`
}

type TaskParams struct {
	BeforeScript []string
	Script       []string
	AfterScript  []string
}
