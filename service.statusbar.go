package main

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

type StatusBarService struct{}

func (s *StatusBarService) Quit() {
	application.Get().Quit()
}
