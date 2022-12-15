package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zouchangfu/go-zero-element-admin/common/errx"
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

func (l *QueryByIdLogic) QueryById(req *types.FormParamId) (resp *types.UserResp, err error) {
	user, sqlResult := l.svcCtx.UserDao.GetById(req.Id)
	if sqlResult.Error != nil {
		return nil, errx.NewErrCode(errx.DbError)
	}

	userResp := types.UserResp{}
	if err := copier.Copy(&userResp, &user); err != nil {
		return nil, errx.NewErrCode(errx.DbError)
	}
	return &userResp, nil
}
