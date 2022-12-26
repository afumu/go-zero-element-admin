package dict

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zouchangfu/go-zero-element-admin/internal/common/errx"

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

func (l *QueryByIdLogic) QueryById(req *types.FormParamId) (resp *types.DictResp, err error) {
	dict, sqlResult := l.svcCtx.DictDao.GetById(req.Id)
	if sqlResult.Error != nil {
		return nil, errx.NewErrCode(errx.DbError)
	}

	dictResp := types.DictResp{}
	if err := copier.Copy(&dictResp, &dict); err != nil {
		return nil, errx.NewErrCode(errx.DbError)
	}
	return &dictResp, nil
}
