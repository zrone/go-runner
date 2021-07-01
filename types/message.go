package types

type Message struct {
	Ref        string     `json:"ref"`
	Repository Repository `json:"repository"`
}

type Repository struct {
	HtmlUrl string `json:"html_url"`
}
