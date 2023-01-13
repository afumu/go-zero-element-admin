package menu

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zouchangfu/go-zero-element-admin/internal/common/errx"
	"github.com/zouchangfu/go-zero-element-admin/internal/model"

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

	menus, sqlResult := l.svcCtx.MenuDao.GetUserMenu()
	if sqlResult.Error != nil {
		return nil, errx.NewErrCode(errx.DbError)
	}

	// 构建响应菜单
	resp = buildMenuResp(menus)

	// 构建子菜单
	childrenMenus := BuildChildrenMenus(resp.Menus)
	resp.Menus = childrenMenus

	return resp, nil
}

func buildMenuResp(menus []*model.SysMenu) *types.UserMenuResp {
	var r = &types.UserMenuResp{}
	for _, v := range menus {
		var menuResp types.MenuResp
		copier.Copy(&menuResp, &v)
		r.Menus = append(r.Menus, &menuResp)
	}
	return r
}

func BuildChildrenMenus(menus []*types.MenuResp) []*types.MenuResp {
	menuMap := make(map[int64]*types.MenuResp)
	for _, menu := range menus {
		menuMap[menu.Id] = menu
	}

	for _, menu := range menus {
		if menu.ParentId != 0 {
			parentMenu := menuMap[menu.ParentId]
			parentMenu.Children = append(parentMenu.Children, menu)
		}
	}

	var subMenus []*types.MenuResp
	for _, menu := range menus {
		if menu.ParentId == 0 {
			subMenus = append(subMenus, menu)
		}
	}
	return subMenus
}
