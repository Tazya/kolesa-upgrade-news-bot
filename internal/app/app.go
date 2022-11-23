package app

import (
	"fmt"
	"kolesa-upgrade-team/delivery-bot/cmd/bot"
	"kolesa-upgrade-team/delivery-bot/config"
	"kolesa-upgrade-team/delivery-bot/http/handlers"
	"kolesa-upgrade-team/delivery-bot/internal/models"
	"log"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Run(config *config.Config) {
	if config.Mysql.DbPassword == "" {
		password, err := os.ReadFile("config/DbPassword.txt")
		if err != nil {
			log.Fatal(err)
		}
		config.Mysql.DbPassword = string(password)
	}

	router := http.NewServeMux()
	db, err := gorm.Open(mysql.Open(getDsn(config)), &gorm.Config{})
	if err != nil {
		fmt.Println("Ошибка при подключении к БД")
		return
	}

	if config.Bot.BotToken == "" {
		token, err := os.ReadFile("config/token.txt")
		if err != nil {
			log.Fatal(err)
		}
		config.Bot.BotToken = string(token)
	}
	handler := &handlers.Handler{
		Bot:  bot.InitBot(config.Bot.BotToken),
		User: &models.UserModel{Db: db},
	}
	handlers.InitRoutes(router, handler)

	server := &http.Server{
		Addr:         ":" + config.Http.Port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Соединение не установлено %v", err)
		}
	}()
	log.Println("Запуск начался")
	handler.Bot.Handle("/start", handler.StartHandler)
	handler.Bot.Handle("/hello", handler.HelloHandler)
	handler.Bot.Start()
}
func getDsn(config *config.Config) string {
	Dsn := config.Mysql.DbUser + ":" + config.Mysql.DbPassword + "@tcp(" + config.Mysql.DbHost + ":" + config.Mysql.DbPort + ")/" + config.Mysql.DbName

	return Dsn
}
