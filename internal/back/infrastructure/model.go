package infrastructure

type ActuatorInfo struct {
	Git   GitInfo   `json:"git"`
	Build BuildInfo `json:"build"`
}

type GitInfo struct {
	Branch string        `json:"branch"`
	Commit CommitGitInfo `json:"commit"`
}

type CommitGitInfo struct {
	Id   string `json:"id"`
	Time string `json:"time"`
}

type BuildInfo struct {
	Group    string `json:"group"`
	Artifact string `json:"artifact"`
	Name     string `json:"name"`
	Version  string `json:"version"`
	Time     string `json:"time"`
}

type Parameters struct {
	Envs []EnvParameter `yaml:"envs"`
}

type EnvParameter struct {
	Name          string         `yaml:"name"`
	CacheDuration int            `yaml:"cache-duration-seconds"`
	Apps          []AppParameter `yaml:"apps"`
}

type AppParameter struct {
	Name           string `yaml:"name"`
	Url            string `yaml:"url"`
	CompareWithEnv string `yaml:"compare-with-env"`
}
