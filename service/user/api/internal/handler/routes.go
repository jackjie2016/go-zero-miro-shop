// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "zero-mal/service/user/api/internal/handler/user"
	"zero-mal/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.IsLogin},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/users/login",
					Handler: user.LoginHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/users/register",
					Handler: user.RegisterHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/user/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.IsAdmin},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/users/:userId",
					Handler: user.GetUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/users/edit",
					Handler: user.EditUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/users/create",
					Handler: user.CreateUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/users/logout",
					Handler: user.LogoutHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/user/api/v1"),
	)
}
