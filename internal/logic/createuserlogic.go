// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"github.com/sshehrozali/basic-service/internal/svc"
	"github.com/sshehrozali/basic-service/internal/model"
	"github.com/sshehrozali/basic-service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req types.CreateUserRequest) (string, error) {

	newUser := &model.User{
		Name: req.Name,
		Email: req.Email,
	} 

	err := l.svcCtx.DB.Create(newUser).Error
	if err != nil {
		return "error while saving new user to database", err
	}

	return "new user created successfully", nil
}
