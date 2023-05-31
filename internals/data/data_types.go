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
}

type AppSpec struct {
	Stage       string   `yaml:"stage"`
	Environment string   `yaml:"environment"`
	Name        string   `yaml:"name"`
	Namespace   string   `yaml:"namespace"`
	RepoURL     string   `yaml:"repoUrl"`
	Project     string   `yaml:"project"`
	UseHelm     bool     `yaml:"useHelm"`
	ValueFiles  []string `yaml:"valueFiles"`
	Plugin      struct {
		Name string `yaml:"name"`
		Env  []struct {
			Name  string `yaml:"name"`
			Value string `yaml:"value"`
		} `yaml:"env"`
	} `yaml:"plugin"`
	Versions  []string `yaml:"versions,omitempty"`
	ChartRepo string   `yaml:"chartRepo,omitempty"`
}

type AppMeta map[string]AppSpec

// detect
type App struct {
	DirectoryPath   string
	Name            string
	Namespace       string
	ValueFiles      []string
	ChartRepository string
	Version         []string
}

type Requirement struct {
	Dependencies []struct {
		Name       string `yaml:"name"`
		Version    string `yaml:"version"`
		Repository string `yaml:"repository"`
	} `yaml:"dependencies"`
}
