package menu

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zouchangfu/go-zero-element-admin/internal/common/errx"
	"github.com/zouchangfu/go-zero-element-admin/internal/svc"
	"github.com/zouchangfu/go-zero-element-admin/internal/types"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List() (resp *types.MenuListResp, err error) {
	allMenus, sqlResult := l.svcCtx.MenuDao.ListAll()
	if sqlResult.Error != nil {
		return nil, errx.NewErrCode(errx.DbError)
	}
	var menus []*types.MenuResp
	for _, u := range allMenus {
		menuResp := types.MenuResp{}
		copier.Copy(&menuResp, &u)
		tm1 := time.UnixMilli(u.CreatedAt.UnixMilli())
		menuResp.CreatedAt = tm1.Format("2006-01-02 15:04:05")

		if u.UpdatedAt.UnixMilli() > 0 {
			tm2 := time.UnixMilli(u.UpdatedAt.UnixMilli())
			menuResp.UpdatedAt = tm2.Format("2006-01-02 15:04:05")
		}

		menus = append(menus, &menuResp)
	}

	childrenMenus := BuildChildrenMenus(menus)
	return &types.MenuListResp{MenuList: childrenMenus}, nil
}
