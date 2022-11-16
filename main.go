package main

import (
	"flag"
	"kolesa-upgrade-team/delivery-bot/cmd/bot"
	"kolesa-upgrade-team/delivery-bot/config"
	"kolesa-upgrade-team/delivery-bot/internal/app"
	"log"
	"sync"

	"github.com/BurntSushi/toml"
)

func main() {

	port := flag.String("port", "8888", "HTTP port")
	DbHost := flag.String("host", "localhost", "host name")
	configPath := flag.String("config", "config/local.toml", "Path to config file")
	flag.Parse()

	config := &config.Config{}
	_, err := toml.DecodeFile(*configPath, config)

	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}

	config.Port = *port
	config.DbHost = *DbHost

	var wg sync.WaitGroup
	wg.Add(2)
	go app.Run(config, &wg)
	go bot.Run(config, &wg)
	wg.Wait()
}
