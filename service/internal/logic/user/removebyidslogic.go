package user

import (
	"context"
	"github.com/zouchangfu/go-zero-element-admin/common/errx"

	"github.com/zouchangfu/go-zero-element-admin/service/internal/svc"
	"github.com/zouchangfu/go-zero-element-admin/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveByIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveByIdsLogic {
	return &RemoveByIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveByIdsLogic) RemoveByIds(req *types.FormParamIds) error {
	if err := l.svcCtx.UserModel.DeleteByIds(l.ctx, req.Ids); err != nil {
		return errx.NewErrCode(errx.DbError)
	}
	return nil
}
