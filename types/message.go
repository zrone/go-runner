package types

const (
	NOTIFICATION_WORK_SERVER = "notification_work_server"
	UNKNOWN_SYMBOL           = "Unknown symbol"
)

type Message struct {
	Ref        string     `json:"ref"`
	After      string     `json:"after"`
	UserName   string     `json:"user_name"`
	Repository Repository `json:"repository"`
}

type Repository struct {
	HtmlUrl string `json:"html_url"`
	SshUrl  string `json:"ssh_url"`
}

type Task struct {
	Symbol string
	Branch string `json:"branch"`
}

type ErrorMessage struct {
	Type   string
	TaskId string
}
