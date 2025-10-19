package logic

import (
	"context"
	"grpc-common/ucenter/types/register"
	"time"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUcenterapiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.Request) (resp *types.Response, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	l.svcCtx.UCRegisterRpc.RegisterByPhone(ctx, &register.RegReq{})
	if err != nil {
		return nil, err
	}
	logx.Info("register logic")
	return
}

func (l *RegisterLogic) SendCode(res *types.CodeRequest) (resp *types.Response, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	l.svcCtx.UCRegisterRpc.SendCode(ctx, &register.CodeReq{
		Phone:   res.Phone,
		Country: res.Country,
	})
	if err != nil {
		return nil, err
	}
	logx.Info("register logic")
	return
}
