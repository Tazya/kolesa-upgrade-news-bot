package config

type Config_server struct {
	Port string
}

type Config_bot struct {
	Env      string
	BotToken string
	Dsn      string
}
