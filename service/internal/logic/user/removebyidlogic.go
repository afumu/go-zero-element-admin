package user

import (
	"context"

	"github.com/zouchangfu/go-zero-element-admin/service/internal/svc"
	"github.com/zouchangfu/go-zero-element-admin/service/internal/types"

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

func (l *RemoveByIdLogic) RemoveById(req *types.FormParamId) (resp *types.Result, err error) {
	// todo: add your logic here and delete this line

	return
}
