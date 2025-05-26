package globals

type Config struct {
	DB struct {
		Type   string      `yaml:"type"`
		Config MysqlConfig `yaml:"config"`
	}
	Redis struct {
		RedisConfig
	}
	App App `yaml:"app"`
}
type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	MaxCon   int    `yaml:"maxCon"`
	MaxIdle  int    `yaml:"maxIdle"`
}

type RedisConfig struct {
}

type App struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}
