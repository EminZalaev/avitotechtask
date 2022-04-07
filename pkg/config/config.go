package config

func GetConfig() *Config {
	return &Config{
		Server:   getServerConfig(),
		DataBase: getDataBaseConfig(),
	}
}

func getDataBaseConfig() DataBase {
	return DataBase{
		DbHost:   "localhost",
		DbPort:   "5432",
		User:     "postgres",
		Password: "1234",
		DbName:   "postgres",
		SSLmode:  "disable",
	}
}

func getServerConfig() Server {
	return Server{
		Host: "localhost",
		Port: ":8080",
	}
}
