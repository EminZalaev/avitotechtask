package config

type Config struct {
	Server
	DataBase
}

type Server struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
}

type DataBase struct {
	DbHost   string `yaml:"DbHost"`
	DbPort   string `yaml:"DbPort"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	DbName   string `yaml:"DbName"`
	SSLmode  string `yaml:"SSLmode"`
}
