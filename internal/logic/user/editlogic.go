package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zouchangfu/go-zero-element-admin/common/errx"
	"github.com/zouchangfu/go-zero-element-admin/internal/svc"
	"github.com/zouchangfu/go-zero-element-admin/internal/types"
	"github.com/zouchangfu/go-zero-element-admin/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditLogic {
	return &EditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditLogic) Edit(req *types.UserEditReq) error {
	user := model.SysUser{}
	if err := copier.Copy(&user, &req); err != nil {
		return errx.NewErrCode(errx.ServerCommonError)
	}

	if err := l.svcCtx.UserDao.UpdateById(&user).Error; err != nil {
		return errx.NewErrCode(errx.DbError)
	}

	return nil
}