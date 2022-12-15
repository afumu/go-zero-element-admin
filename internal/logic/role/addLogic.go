package role

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zouchangfu/go-zero-element-admin/internal/common/errx"
	"github.com/zouchangfu/go-zero-element-admin/internal/common/utils"
	"github.com/zouchangfu/go-zero-element-admin/internal/model"
	"github.com/zouchangfu/go-zero-element-admin/internal/svc"
	"github.com/zouchangfu/go-zero-element-admin/internal/types"
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

func (l *AddLogic) Add(req *types.RoleAddReq) error {

	role := model.SysRole{}
	userId := utils.GetLoginUserId(l.ctx)
	role.CreatedBy = userId
	role.UpdatedBy = userId
	if err := copier.Copy(&role, &req); err != nil {
		return errx.NewErrCode(errx.ServerCommonError)
	}
	if err := l.svcCtx.RoleDao.Save(&role).Error; err != nil {
		return errx.NewErrCode(errx.DbError)
	}

	return nil
}
