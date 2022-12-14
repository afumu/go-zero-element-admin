package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zouchangfu/go-zero-element-admin/common/errx"

	"github.com/zouchangfu/go-zero-element-admin/service/internal/svc"
	"github.com/zouchangfu/go-zero-element-admin/service/internal/types"

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

	sysUser, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, errx.NewErrCode(errx.DbError)
	}

	u := types.UserResp{}
	if err := copier.Copy(&u, &sysUser); err != nil {
		return nil, errx.NewErrCode(errx.CopierError)
	}

	if sysUser.CreatedTime.Valid {
		u.CreatedTime = sysUser.CreatedTime.Time.UnixMilli()
	}

	if sysUser.CreatedTime.Valid {
		u.UpdatedTime = sysUser.UpdatedTime.Time.UnixMilli()
	}

	return &u, nil
}
