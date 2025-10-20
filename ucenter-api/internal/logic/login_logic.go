package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"grpc-common/ucenter/types/login"
	"mscoin-common/tools"

	"time"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	regReq := &login.LoginReq{}
	if err := copier.Copy(regReq, req); err != nil {
		return nil, err
	}
	res, err := l.svcCtx.UCLoginRpc.Login(ctx, regReq)
	if err != nil {
		return nil, err
	}
	resp = &types.LoginRes{}
	if err := copier.Copy(resp, res); err != nil {
		return nil, err
	}
	logx.Info("Login logic")
	return
}

func (l *LoginLogic) CheckLogin(token string) (bool, error) {
	_, err := tools.ParseToken(token, l.svcCtx.Config.JWT.AccessSecret)
	if err != nil {
		return false, nil
	}
	return true, nil

}
