package tests

import (
	"kolesa-upgrade-team/delivery-bot/cmd/bot"
	"kolesa-upgrade-team/delivery-bot/config"
	"testing"
)

func TestGetDsn(t *testing.T) {
	config := &config.Config{
		Mysql: config.Mysql{
			DbUser:     "root",
			DbPassword: "qwerty",
			DbHost:     "176.0.0.1",
			DbPort:     "1234",
			DbName:     "Upgrade",
		},
	}

	got := bot.GetDsn(config)
	want := "root:qwerty@tcp(176.0.0.1:1234)/Upgrade"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
