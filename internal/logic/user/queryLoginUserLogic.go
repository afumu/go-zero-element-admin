package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zouchangfu/go-zero-element-admin/internal/common/errx"
	"github.com/zouchangfu/go-zero-element-admin/internal/common/utils"

	"github.com/zouchangfu/go-zero-element-admin/internal/svc"
	"github.com/zouchangfu/go-zero-element-admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryLoginUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryLoginUserLogic {
	return &QueryLoginUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryLoginUserLogic) QueryLoginUser() (resp *types.UserResp, err error) {
	userId := utils.GetLoginUserId(l.ctx)

	user, sqlResult := l.svcCtx.UserDao.GetById(userId)
	if sqlResult.Error != nil {
		return nil, errx.NewErrCode(errx.DbError)
	}

	userResp := types.UserResp{}
	if err := copier.Copy(&userResp, &user); err != nil {
		return nil, errx.NewErrCode(errx.CopierError)
	}

	return &userResp, nil
}
