package user

import (
	"context"
	"github.com/acmestack/gorm-plus/gplus"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/zouchangfu/go-zero-element-admin/internal/common/errx"
	"github.com/zouchangfu/go-zero-element-admin/internal/model"
	"github.com/zouchangfu/go-zero-element-admin/internal/utils"
	"time"

	"github.com/zouchangfu/go-zero-element-admin/internal/svc"
	"github.com/zouchangfu/go-zero-element-admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserNoExistsError = errx.NewErrMsg("用户不存在")
var ErrGenerateTokenError = errx.NewErrMsg("生成token失败")
var ErrUsernamePwdError = errx.NewErrMsg("账号或密码不正确")

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {

	query := gplus.NewQuery[model.SysUser]()
	query.Eq(model.SysUserColumn.Username, req.Username)
	user, sqlResult := l.svcCtx.UserDao.GetOne(query)
	if sqlResult.Error != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DbError), "根据账号查询用户信息失败，username:%s,err:%v", req.Username, err)
	}

	if user == nil {
		return nil, errors.Wrapf(ErrUserNoExistsError, "username:%s", req.Username)
	}

	if !(utils.Md5ByString(req.Password) == user.Password) {
		return nil, errors.Wrap(ErrUsernamePwdError, "密码匹配出错")
	}

	tokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := tokenLogic.GenerateToken(user.Id)

	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", user.Id)
	}

	return &types.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type GenerateTokenResp struct {
	AccessToken  string
	AccessExpire int64
	RefreshAfter int64
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(userId int64) (*GenerateTokenResp, error) {
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, userId)
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "getJwtToken err userId:%d , err:%v", userId, err)
	}

	return &GenerateTokenResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *GenerateTokenLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
