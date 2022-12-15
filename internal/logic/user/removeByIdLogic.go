package user

import (
	"context"
	errx2 "github.com/zouchangfu/go-zero-element-admin/internal/common/errx"
	"github.com/zouchangfu/go-zero-element-admin/internal/svc"
	"github.com/zouchangfu/go-zero-element-admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveByIdLogic {
	return &RemoveByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveByIdLogic) RemoveById(req *types.FormParamId) error {
	if err := l.svcCtx.UserDao.RemoveById(req.Id).Error; err != nil {
		return errx2.NewErrCode(errx2.DbError)
	}
	return nil
}
