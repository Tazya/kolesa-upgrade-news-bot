package config

type Http struct {
	Port string
}

type Bot struct {
	BotToken string
}

type Mysql struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
}

type Config struct {
	Http  Http
	Bot   Bot
	Mysql Mysql
}
