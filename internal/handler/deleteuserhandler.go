// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"github.com/sshehrozali/basic-service/internal/logic"
	"github.com/sshehrozali/basic-service/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDeleteUserLogic(r.Context(), svcCtx)
		err := l.DeleteUser()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
