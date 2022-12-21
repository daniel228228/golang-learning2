package config

type Config struct {
	ServicePort string `env_config:"SERVICE_PORT"`
	Log         struct {
		File         string `env_config:"LOG_FILE"`
		WriteConsole bool   `env_config:"LOG_WRITE_CONSOLE"`
		WriteFile    bool   `env_config:"LOG_WRITE_FILE"`
		ConsoleLevel string `env_config:"LOG_CONSOLE_LEVEL"`
		FileLevel    string `env_config:"LOG_FILE_LEVEL"`
	}
	DB struct {
		Host string `env_config:"DB_HOST"`
		Port string `env_config:"DB_PORT"`
		User string `env_config:"DB_USER"`
		Pass string `env_config:"DB_PASS"`
		Name string `env_config:"DB_NAME"`
	}
}
