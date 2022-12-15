package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	errx2 "github.com/zouchangfu/go-zero-element-admin/internal/common/errx"
	"github.com/zouchangfu/go-zero-element-admin/internal/model"
	"github.com/zouchangfu/go-zero-element-admin/internal/svc"
	"github.com/zouchangfu/go-zero-element-admin/internal/types"
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
		return errx2.NewErrCode(errx2.ServerCommonError)
	}

	if err := l.svcCtx.UserDao.UpdateById(&user).Error; err != nil {
		return errx2.NewErrCode(errx2.DbError)
	}

	return nil
}
