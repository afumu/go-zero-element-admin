package permission

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

func (l *ListLogic) List() (resp *types.PermissionListResp, err error) {
	allPermissions, sqlResult := l.svcCtx.PermissionDao.ListAll()
	if sqlResult.Error != nil {
		return nil, errx.NewErrCode(errx.DbError)
	}

	var permissions []*types.PermissionResp
	for _, u := range allPermissions {
		permissionsResp := types.PermissionResp{}
		copier.Copy(&permissionsResp, &u)
		permissionsResp.CreatedAt = u.CreatedAt.UnixMilli()
		permissionsResp.UpdatedAt = u.UpdatedAt.UnixMilli()
		permissions = append(permissions, &permissionsResp)
	}

	return &types.PermissionListResp{PermissionList: permissions}, nil
}
