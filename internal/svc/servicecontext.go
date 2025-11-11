// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"github.com/sshehrozali/basic-service/internal/config"
	"gorm.io/gorm"
	"github.com/sshehrozali/basic-service/internal/model"
)

type ServiceContext struct {
	Config config.Config
	DB *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB: model.InitDB(c.Mysql.DSN),
	}
}
