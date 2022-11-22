package bot

import (
	"kolesa-upgrade-team/delivery-bot/config"
	"testing"
)

func TestGetDsn(t *testing.T) {
	config := &config.Config{
		DB: config.DB{
			User:     "root",
			Password: "qwerty",
			Host:     "176.0.0.1",
			Port:     "1234",
			Name:     "Upgrade",
		},
	}

	got := getDsn(config)
	want := "root:qwerty@tcp(176.0.0.1:1234)/Upgrade"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
