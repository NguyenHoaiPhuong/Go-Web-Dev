package services

import (
	"github.com/NguyenHoaiPhuong/Go-Web-Dev/020_AWS/config"
	"github.com/NguyenHoaiPhuong/Go-Web-Dev/020_AWS/services/awss3"
)

// Service implements all methods in the Actions interface
type Service struct {
	S3 awss3.IService
}

// Init :
func (svc *Service) Init() {
	conf := config.GlobalConfig()
	svc.S3 = awss3.NewService(conf.S3config)
}
