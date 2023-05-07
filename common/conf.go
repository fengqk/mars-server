package common

type (
	Server struct {
		Ip   string `yaml:"ip"`
		Port int    `yaml:"port"`
	}

	Mysql struct {
		Ip           string `yaml:"ip"`
		Port         int    `yaml:"port"`
		Name         string `yaml:"name"`
		User         string `yaml:"user"`
		Password     string `yaml:"password"`
		MaxIdleConns int    `yaml:"maxIdleConns"`
		MaxOpenConns int    `yaml:"maxOpenConns"`
	}

	Redis struct {
		Ip       string `yaml:"ip"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
	}
)
