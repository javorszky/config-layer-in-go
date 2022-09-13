package service

import (
	"time"

	"github.com/javorszky/config-layer/config"
)

type App struct {
	c config.AppConfig
}

func New(c config.AppConfig) App {
	return App{c: c}
}

func (a App) Start() {
	time.Sleep(3 * time.Minute)
}
