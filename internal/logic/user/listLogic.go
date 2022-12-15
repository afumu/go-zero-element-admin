package user

import (
	"context"
	"github.com/jinzhu/copier"
	errx2 "github.com/zouchangfu/go-zero-element-admin/internal/common/errx"
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

func (l *ListLogic) List() (resp *types.UserListResp, err error) {
	allUsers, sqlResult := l.svcCtx.UserDao.ListAll()
	if sqlResult.Error != nil {
		return nil, errx2.NewErrCode(errx2.DbError)
	}

	var users []*types.UserResp
	for _, u := range allUsers {
		userResp := types.UserResp{}
		copier.Copy(&userResp, &u)
		userResp.CreatedAt = u.CreatedAt.UnixMilli()
		userResp.UpdatedAt = u.UpdatedAt.UnixMilli()
		users = append(users, &userResp)
	}

	return &types.UserListResp{UserList: users}, nil
}
