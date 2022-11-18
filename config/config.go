package config

type http struct {
	Port string
}

type bot struct {
	BotToken string
}

type mysql struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
}

type Config struct {
	Http  http
	Bot   bot
	Mysql mysql
}
