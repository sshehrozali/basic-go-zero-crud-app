// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"github.com/sshehrozali/basic-service/internal/logic"
	"github.com/sshehrozali/basic-service/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewCreateUserLogic(r.Context(), svcCtx)
		err := l.CreateUser()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
