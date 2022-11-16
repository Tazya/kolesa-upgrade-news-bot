package bot

import (
	"kolesa-upgrade-team/delivery-bot/config"
	"kolesa-upgrade-team/delivery-bot/internal/models"
	"log"
	"os"
	"sync"
	"time"

	"gopkg.in/telebot.v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UpgradeBot struct {
	Bot   *telebot.Bot
	Users *models.UserModel
}

func (bot *UpgradeBot) StartHandler(ctx telebot.Context) error {
	newUser := models.User{
		Name:       ctx.Sender().Username,
		TelegramId: ctx.Chat().ID,
		FirstName:  ctx.Sender().FirstName,
		LastName:   ctx.Sender().LastName,
		ChatId:     ctx.Chat().ID,
	}
	existUser, err := bot.Users.FindOne(ctx.Chat().ID)
	if err != nil {
		log.Printf("Ошибка получения пользователя %v", err)
	}

	if existUser == nil {
		err := bot.Users.Create(newUser)
		if err != nil {
			log.Printf("Ошибка создания пользователя %v", err)
		}
	}
	return ctx.Send("Привет, я дружелюбный бот. Мои команды /hello")
}

func (bot *UpgradeBot) HelloHandler(ctx telebot.Context) error {
	return ctx.Send("Привет, " + ctx.Sender().FirstName + " =^_^=")
}

func InitBot(token string) *telebot.Bot {
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}
	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatalf("Ошибка при инициализации бота %v", err)
	}

	return bot
}

func Run(config *config.Config, wg *sync.WaitGroup) {
	db, err := gorm.Open(sqlite.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к БД %v", err)
	}

	if config.BotToken == "" {
		token, err := os.ReadFile("config/token.txt")
		if err != nil {
			log.Fatalf("Ошибка при чтенни токен файла %v", err)
		}
		config.BotToken = string(token)
	}

	upgradeBot := UpgradeBot{
		Bot:   InitBot(config.BotToken),
		Users: &models.UserModel{Db: db},
	}

	upgradeBot.Bot.Handle("/start", upgradeBot.StartHandler)
	upgradeBot.Bot.Handle("/hello", upgradeBot.HelloHandler)
	upgradeBot.Bot.Start()
}
