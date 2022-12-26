package menu

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
