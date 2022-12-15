package role

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zouchangfu/go-zero-element-admin/internal/common/errx"

	"github.com/zouchangfu/go-zero-element-admin/internal/svc"
	"github.com/zouchangfu/go-zero-element-admin/internal/types"

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

func (l *ListLogic) List() (resp *types.RoleListResp, err error) {
	allRoles, sqlResult := l.svcCtx.RoleDao.ListAll()
	if sqlResult.Error != nil {
		return nil, errx.NewErrCode(errx.DbError)
	}

	var roles []*types.RoleResp
	for _, u := range allRoles {
		roleResp := types.RoleResp{}
		copier.Copy(&roleResp, &u)
		roleResp.CreatedAt = u.CreatedAt.UnixMilli()
		roleResp.UpdatedAt = u.UpdatedAt.UnixMilli()
		roles = append(roles, &roleResp)
	}

	return &types.RoleListResp{RoleList: roles}, nil
}
