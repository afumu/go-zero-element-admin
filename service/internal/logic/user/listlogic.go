package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zouchangfu/go-zero-element-admin/common/errx"

	"github.com/zouchangfu/go-zero-element-admin/service/internal/svc"
	"github.com/zouchangfu/go-zero-element-admin/service/internal/types"

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

func (l *ListLogic) List() (resp *types.UserListResp, err error) {
	allUsers, err := l.svcCtx.UserModel.FindAll(l.ctx)
	if err != nil {
		return nil, errx.NewErrCode(errx.DbError)
	}

	var userList []*types.UserResp
	for _, v := range allUsers {
		userResp := types.UserResp{}
		if err := copier.Copy(&userResp, &v); err != nil {
			return nil, errx.NewErrCode(errx.CopierError)
		}
		if v.CreatedTime.Valid {
			userResp.CreatedTime = v.CreatedTime.Time.UnixMilli()
		}
		if v.UpdatedTime.Valid {
			userResp.UpdatedTime = v.UpdatedTime.Time.UnixMilli()
		}
		userList = append(userList, &userResp)
	}

	return &types.UserListResp{UserList: userList}, nil
}
