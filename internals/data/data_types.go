package data

// repo package
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

// watch package
type Chart struct {
	Name string `json:"name"`
	Prev string `json"prev"`
	Curr string `json"curr"`
}

type AppSpec struct {
	Stage       string `yaml:"stage"`
	Environment string `yaml:"environment"`
	Name        string `yaml:"name"`
	Namespace   string `yaml:"namespace"`
	RepoURL     string `yaml:"repoUrl"`
	Project     string `yaml:"project"`
	UseHelm     bool   `yaml:"useHelm"`
	Plugin      struct {
		Name string `yaml:"name"`
		Env  []struct {
			Name  string `yaml:"name"`
			Value string `yaml:"value"`
		} `yaml:"env"`
	} `yaml:"plugin,omitempty"`
}

type App map[string]AppSpec
