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
	Symbol   string         `gorm:"primaryKey;autoIncrement:false" json:"symbol" validate:"required"`
	Name     string         `json:"name" validate:"required"`
	Secret   string         `json:"secret" validate:"required"`
	Path     string         `json:"path" validate:"required"`
	Auth     Authentication `gorm:"embedded" validate:"required"`
	IsDelete bool           `json:"is_delete"`
}

// TASK log
type TaskLog struct {
	Symbol   string
	Uuid     string
	State    string
	CreateAt string
	EndAt    string
	Args     string
}

type Authentication struct {
	Scheme int    `json:"scheme" validate:"required"`
	User   string `json:"user" validate:"required"`
	Host   string `json:"host" validate:"required,ipv4"`
	Port   int    `json:"port" validate:"required"`
	Pwd    string `json:"-" validate:"required"`
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
