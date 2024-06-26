package config

type Config struct {
	Redis RedisConfig `yaml:"Redis"`
	Web   WebConfig   `yaml:"Web"`
	Bsz   BszConfig   `yaml:"Bsz"`
}

type RedisConfig struct {
	Address    string `yaml:"Address"`
	Password   string `yaml:"Password"`
	Database   int    `yaml:"Database"`
	TLS        bool   `yaml:"TLS"`
	Prefix     string `yaml:"Prefix"`
	MaxIdle    int    `yaml:"MaxIdle"`
	MaxActive  int    `yaml:"MaxActive"`
	MinIdle    int    `yaml:"MinIdle"`
	MaxRetries int    `yaml:"MaxRetries"`
}

type WebConfig struct {
	Address string `yaml:"Address"`
	Cors    string `yaml:"Access-Control-Allow-Origin"`
	Debug   bool   `yaml:"Debug"`
	Log     bool   `yaml:"Log"`
}

type BszConfig struct {
	JwtSecret string `yaml:"JwtSecret"`
}
