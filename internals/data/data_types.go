package data

type Repository struct {
	Name        string      `json:"name"`
	Host        string      `json:"host"`
	Private     bool        `json:"private"`
	Credentials Credentials `json:"credentials,omitempty"`
}

type Credentials struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
