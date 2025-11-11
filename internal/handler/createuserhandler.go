// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sshehrozali/basic-service/internal/types"

	"github.com/sshehrozali/basic-service/internal/logic"
	"github.com/sshehrozali/basic-service/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewCreateUserLogic(r.Context(), svcCtx)

		// Parse request body into CreateUserRequest struct
		var req types.CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httpx.ErrorCtx(r.Context(), w, errors.New("server error while parsing request body"))	// error parsing request body
			return
		}

		if req.Email == "" || req.Name == "" {
			httpx.ErrorCtx(r.Context(), w, errors.New("request invalid or not correct")) // missing required fields
			return
		}

		_, err := l.CreateUser(req)

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}

		httpx.OkJsonCtx(r.Context(), w, &types.CreateUserResponse{
			Name:    req.Name,
			Email:   req.Email,
			Message: "User created successfully",
		})
	}
}
