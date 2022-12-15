package user

import (
	"context"
	"github.com/zouchangfu/go-zero-element-admin/internal/common/errx"
	"github.com/zouchangfu/go-zero-element-admin/internal/svc"
	"github.com/zouchangfu/go-zero-element-admin/internal/types"
	"strings"

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
	if err := l.svcCtx.UserDao.RemoveByIds(strings.Split(req.Ids, ",")).Error; err != nil {
		return errx.NewErrCode(errx.DbError)
	}
	return nil
}
