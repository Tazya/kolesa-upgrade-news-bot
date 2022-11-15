package main

import (
	"flag"
	"kolesa-upgrade-team/delivery-bot/cmd/bot"
	"kolesa-upgrade-team/delivery-bot/config"
	"kolesa-upgrade-team/delivery-bot/internal/app"
	"kolesa-upgrade-team/delivery-bot/internal/models"
	"log"
	"os"
	"sync"

	"github.com/BurntSushi/toml"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func run_server(port *string, wg *sync.WaitGroup) {
	config_server := &config.Config_server{}
	config_server.Port = *port

	if err := app.Run(config_server); err != nil {
		log.Fatalf("%v", err)
	}
}

func run_bot(configPath *string, wg *sync.WaitGroup) {

	config_bot := &config.Config_bot{}
	_, err := toml.DecodeFile(*configPath, config_bot)

	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}

	db, err := gorm.Open(sqlite.Open(config_bot.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Ошибка подключения к БД %v", err)
	}

	if config_bot.BotToken == "" {
		token, err := os.ReadFile("config/token.txt")
		if err != nil {
			log.Fatal(err)
		}
		config_bot.BotToken = string(token)
	}

	upgradeBot := bot.UpgradeBot{
		Bot:   bot.InitBot(config_bot.BotToken),
		Users: &models.UserModel{Db: db},
		Tasks: &models.TaskModel{Db: db},
	}

	upgradeBot.Bot.Handle("/start", upgradeBot.StartHandler)
	upgradeBot.Bot.Handle("/hello", upgradeBot.HelloHandler)

	upgradeBot.Bot.Start()
}

func main() {

	port := flag.String("port", "8888", "HTTP port")
	configPath := flag.String("config", "config/local.toml", "Path to config file")
	flag.Parse()

	var wg sync.WaitGroup
	wg.Add(2)
	go run_server(port, &wg)
	go run_bot(configPath, &wg)
	wg.Wait()
}
