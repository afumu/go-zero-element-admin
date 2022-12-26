package dict

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zouchangfu/go-zero-element-admin/internal/common/errx"

	"github.com/zouchangfu/go-zero-element-admin/internal/svc"
	"github.com/zouchangfu/go-zero-element-admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List() (resp *types.DictListResp, err error) {
	allDicts, sqlResult := l.svcCtx.DictDao.ListAll()
	if sqlResult.Error != nil {
		return nil, errx.NewErrCode(errx.DbError)
	}

	var dicts []*types.DictResp
	for _, u := range allDicts {
		dictResp := types.DictResp{}
		copier.Copy(&dictResp, &u)
		dictResp.CreatedAt = u.CreatedAt.UnixMilli()
		dictResp.UpdatedAt = u.UpdatedAt.UnixMilli()
		dicts = append(dicts, &dictResp)
	}

	return &types.DictListResp{DictList: dicts}, nil
}
