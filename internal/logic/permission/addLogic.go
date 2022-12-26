package permission

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zouchangfu/go-zero-element-admin/internal/common/errx"
	"github.com/zouchangfu/go-zero-element-admin/internal/common/utils"
	"github.com/zouchangfu/go-zero-element-admin/internal/model"

	"github.com/zouchangfu/go-zero-element-admin/internal/svc"
	"github.com/zouchangfu/go-zero-element-admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.PermissionAddReq) error {
	permission := model.SysPermission{}
	userId := utils.GetLoginUserId(l.ctx)
	permission.CreatedBy = userId
	permission.UpdatedBy = userId
	if err := copier.Copy(&permission, &req); err != nil {
		return errx.NewErrCode(errx.ServerCommonError)
	}
	if err := l.svcCtx.PermissionDao.Save(&permission).Error; err != nil {
		return errx.NewErrCode(errx.DbError)
	}
	return nil
}
