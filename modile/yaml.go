package module
// Yaml struct of yaml
type Yaml struct {
	Mysql struct {
		User string `yaml:"user"`
		Host string `yaml:"host"`
		Password string `yaml:"password"`
		Port string `yaml:"port"`
		Name string `yaml:"name"`

	}
	Crontab struct {
		Period string `yaml:"period"`
	}

	LogFile struct {
		FileName string `yaml:"filename"`
	}
}
