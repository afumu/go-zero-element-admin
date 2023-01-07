package menu

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zouchangfu/go-zero-element-admin/internal/common/errx"

	"github.com/zouchangfu/go-zero-element-admin/internal/svc"
	"github.com/zouchangfu/go-zero-element-admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserMenuLogic {
	return &QueryUserMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryUserMenuLogic) QueryUserMenu() (resp *types.UserMenuResp, err error) {
	var r = &types.UserMenuResp{}
	menus, sqlResult := l.svcCtx.MenuDao.GetUserMenu()
	if sqlResult.Error != nil {
		return nil, errx.NewErrCode(errx.DbError)
	}
	for _, v := range menus {
		var menuResp types.MenuResp
		copier.Copy(&menuResp, &v)
		menuResp.CreatedAt = v.CreatedAt.UnixMilli()
		menuResp.UpdatedAt = v.UpdatedAt.UnixMilli()
		r.Menus = append(r.Menus, &menuResp)
	}
	return r, nil
}
