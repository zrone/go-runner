package types

type Config struct {
	Domain     string
	Host       string
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
	Secret   string         `json:"secret"`
	Path     string         `json:"path" validate:"required"`
	Option   uint8          `json:"option" validate:"required"` // 部署类型 1 自动化部署 2 计划发布
	Auth     Authentication `gorm:"embedded" validate:"required"`
	IsDelete bool           `json:"is_delete"`
}

// TASK log
type TaskLog struct {
	ID        int64  `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Symbol    string `json:"symbol"`
	Committer string `json:"committer"` // 触发者
	Version   string `json:"version"`   // 版本
	Uuid      string `json:"uuid"`
	State     string `json:"state"`
	CreateAt  int64  `json:"create_at"`
	EndAt     int64  `json:"end_at"`
	Type      uint8  `json:"type"`
	Args      string `json:"-"`
}

// manager
type Manager struct {
	ID       int64  `json:"id" gorm:"primaryKey;autoIncrement:true"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
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
	Environment map[string]string `yaml:"environment" json:"environment,omitempty"`
	Prepare     []string          `yaml:"prepare" json:"prepare,omitempty"`
	Script      []string          `yaml:"script" json:"script,omitempty"`
	Release     []string          `yaml:"release" json:"release,omitempty"`
	Webpath     string            `yaml:"webpath" json:"webpath,omitempty"`
	Only        []string          `yaml:"only" json:"only,omitempty"`
}

type TaskParams struct {
	Environment  map[string]string
	BeforeScript []string
	Script       []string
	AfterScript  []string
}
