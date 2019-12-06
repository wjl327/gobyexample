package main

type AppService interface {
	Install() string
	UnInstall() string
}

func NewCoreAppService() *CoreAppService {
	return &CoreAppService{}
}

type CoreAppService struct{}

func (cs *CoreAppService) Install() string {
	return "core app Install"
}

func (cs *CoreAppService) UnInstall() string {
	return "core app UnInstall"
}
