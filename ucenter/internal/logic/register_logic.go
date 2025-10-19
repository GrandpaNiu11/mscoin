package logic

import (
	"context"
	"errors"
	"grpc-common/ucenter/types/register"
	"mscoin-common/tools"
	"time"
	"ucenter/internal/domain"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

const RegisterCacheKey = "REGISTER::"

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	CaptchaDomain *domain.CaptchaDomain
	MemberDomain  *domain.MemberDomain
}

func NewRegisterByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:           ctx,
		svcCtx:        svcCtx,
		Logger:        logx.WithContext(ctx),
		CaptchaDomain: domain.NewCaptchaDomain(),
		MemberDomain:  domain.NewMemberDomain(svcCtx.DB),
	}
}

func (l *RegisterLogic) RegisterByPhone(in *register.RegReq) (*register.RegRes, error) {
	// todo: add your logic here and delete this line
	logx.Info("我被调用了")
	//先交给人机是否通过
	isVerify := l.CaptchaDomain.Verifiy(
		in.Captcha.Server,
		l.svcCtx.Config.Captcha.Vid,
		l.svcCtx.Config.Captcha.Key,
		in.Captcha.Token, 2, in.Ip)
	if !isVerify {
		return nil, errors.New("人机校验不通过")
	}
	//校验验证码
	logx.Info("人机校验通过....")
	redisValue := ""
	err := l.svcCtx.Cache.GetCtx(context.Background(), RegisterCacheKey+in.Phone, &redisValue)
	if err != nil {
		return nil, errors.New("验证码获取错误")
	}
	if in.Code != redisValue {
		return nil, errors.New("验证码输入错误")
	}
	//查询手机号是否注册
	mem, error := l.MemberDomain.FindByPhone(context.Background(), in.Phone)
	if error != nil {
		return nil, errors.New("服务异常")
	}
	if mem != nil {
		return nil, errors.New("此手机号已经被注册")
	}
	logx.Info("第三步完成")

	//注册
	//生成member模型 存入数据库
	err = l.MemberDomain.Register(
		context.Background(),
		in.Username,
		in.Phone,
		in.Password,
		in.Country,
		in.Promotion,
		in.SuperPartner,
	)
	return &register.RegRes{}, nil
}

func (l *RegisterLogic) SendCode(req *register.CodeReq) (*register.NogRes, error) {
	code := tools.Rand4Num()
	//todo 验证码 目前就不发了 后续再说
	go func() {
		logx.Infof("调用成功验证码为:  " + code)
	}()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	err := l.svcCtx.Cache.SetWithExpireCtx(ctx, RegisterCacheKey+req.Phone, code, 5*time.Minute)
	if err != nil {
		return nil, errors.New("验证码存入失败")
	}
	return nil, nil
}
