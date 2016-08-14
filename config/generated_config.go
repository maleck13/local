package config

type Config struct {
	Admin struct {
		Auth string `json:"auth"`
		User string `json:"user"`
	} `json:"admin"`
	ClientDir string `json:"clientDir"`
	Database  struct {
		Addrs    []string `json:"addrs"`
		Database string   `json:"database"`
		Password string   `json:"password"`
		Timeout  int      `json:"timeout"`
		Username string   `json:"username"`
	} `json:"database"`
	Google struct {
		ClientID     string `json:"clientId"`
		ValidatorURL string `json:"validatorUrl"`
	} `json:"google"`
	PprofEnabled bool `json:"pprof_enabled"`
}

//THIS IS A GENERATED FILE. USE scripts/generate_config to regen
