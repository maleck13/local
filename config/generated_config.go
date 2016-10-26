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
	Files struct {
		Aws struct {
			Enabled bool `json:"enabled"`
		} `json:"aws"`
		Local struct {
			DirPath string `json:"dirPath"`
			Enabled bool   `json:"enabled"`
		} `json:"local"`
	} `json:"files"`
	Google struct {
		ClientID     string `json:"clientId"`
		ValidatorURL string `json:"validatorUrl"`
	} `json:"google"`
	Jwt struct {
		Expire string `json:"expire"`
		Secret string `json:"secret"`
	} `json:"jwt"`
	PprofEnabled bool `json:"pprof_enabled"`
	Sendgrid     struct {
		APIKey  string `json:"apiKey"`
		Enabled bool   `json:"enabled"`
	} `json:"sendgrid"`
	SiteHost string `json:"site_host"`
}
//THIS IS A GENERATED FILE. USE scripts/generate_config to regen
