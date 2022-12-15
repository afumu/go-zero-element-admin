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

func (l *AddLogic) Add(req *types.UserAddReq) error {
	user := model.SysUser{}
	if err := copier.Copy(&user, &req); err != nil {
		return errx2.NewErrCode(errx2.ServerCommonError)
	}

	if err := l.svcCtx.UserDao.Save(&user).Error; err != nil {
		return errx2.NewErrCode(errx2.DbError)
	}

	return nil
}
