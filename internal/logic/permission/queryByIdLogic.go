package permission

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zouchangfu/go-zero-element-admin/internal/common/errx"

	"github.com/zouchangfu/go-zero-element-admin/internal/svc"
	"github.com/zouchangfu/go-zero-element-admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryByIdLogic {
	return &QueryByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryByIdLogic) QueryById(req *types.FormParamId) (resp *types.PermissionResp, err error) {
	permission, sqlResult := l.svcCtx.PermissionDao.GetById(req.Id)
	if sqlResult.Error != nil {
		return nil, errx.NewErrCode(errx.DbError)
	}

	permissionResp := types.PermissionResp{}
	if err := copier.Copy(&permissionResp, &permission); err != nil {
		return nil, errx.NewErrCode(errx.DbError)
	}
	return &permissionResp, nil
}
