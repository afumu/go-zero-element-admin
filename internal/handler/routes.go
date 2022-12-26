// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	menu "github.com/zouchangfu/go-zero-element-admin/internal/handler/menu"
	permission "github.com/zouchangfu/go-zero-element-admin/internal/handler/permission"
	role "github.com/zouchangfu/go-zero-element-admin/internal/handler/role"
	user "github.com/zouchangfu/go-zero-element-admin/internal/handler/user"
	"github.com/zouchangfu/go-zero-element-admin/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/list",
				Handler: user.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/add",
				Handler: user.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/queryById",
				Handler: user.QueryByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/user/edit",
				Handler: user.EditHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/user/removeById",
				Handler: user.RemoveByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/user/removeByIds",
				Handler: user.RemoveByIdsHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/role/list",
				Handler: role.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/role/add",
				Handler: role.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/role/queryById",
				Handler: role.QueryByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/role/edit",
				Handler: role.EditHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/role/removeById",
				Handler: role.RemoveByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/role/removeByIds",
				Handler: role.RemoveByIdsHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/menu/list",
				Handler: menu.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/menu/add",
				Handler: menu.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/menu/queryById",
				Handler: menu.QueryByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/menu/edit",
				Handler: menu.EditHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/menu/removeById",
				Handler: menu.RemoveByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/menu/removeByIds",
				Handler: menu.RemoveByIdsHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/permission/list",
				Handler: permission.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/permission/add",
				Handler: permission.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/permission/queryById",
				Handler: permission.QueryByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/permission/edit",
				Handler: permission.EditHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/permission/removeById",
				Handler: permission.RemoveByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/permission/removeByIds",
				Handler: permission.RemoveByIdsHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)
}
