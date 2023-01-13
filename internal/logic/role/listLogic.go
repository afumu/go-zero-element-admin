package role

import (
	"context"
	"github.com/acmestack/gorm-plus/gplus"
	"github.com/jinzhu/copier"
	"github.com/zouchangfu/go-zero-element-admin/internal/common/errx"
	"github.com/zouchangfu/go-zero-element-admin/internal/common/utils"
	"github.com/zouchangfu/go-zero-element-admin/internal/model"

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

func (l *ListLogic) List(req *types.RolePageReq) (resp *types.RolePageResp, err error) {
	page := gplus.NewPage[model.SysRole](req.Current, req.Size)
	pageResult, sqlResult := l.svcCtx.RoleDao.PageAll(page)
	if sqlResult.Error != nil {
		return nil, errx.NewErrCode(errx.DbError)
	}

	var rolePageResp types.RolePageResp
	if err := copier.CopyWithOption(&rolePageResp, &pageResult, utils.CovertTimeToString()); err != nil {
		return nil, errx.NewErrCode(errx.ServerCommonError)
	}

	return &rolePageResp, nil
}
