package domain

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"mscoin-common/msdb"
	"mscoin-common/tools"
	"ucenter/internal/dao"
	"ucenter/model"
	"ucenter/repo"
)

type MemberDomain struct {
	MemberRepo repo.MemberRepo
}

func (d MemberDomain) FindByPhone(ctx context.Context, phone string) (*model.Member, error) {
	mem, error := d.MemberRepo.FindByPhone(ctx, phone)
	if error != nil {
		logx.Error(error)
		return nil, errors.New("数据库异常")
	}
	return mem, nil
}

func (d *MemberDomain) Register(
	ctx context.Context,
	username string,
	phone string,
	password string,
	country string,
	promotion string,
	partner string) error {
	mem := model.NewMember()
	tools.Default(mem)
	mem.Id = 0
	//首先处理密码 密码要md5加密，但是md5不安全，所以我们给加上一个salt值
	pwd, salt := tools.Encode(password, nil)
	mem.Salt = salt
	mem.Password = pwd
	mem.MobilePhone = phone
	mem.Username = username
	mem.Country = country
	mem.PromotionCode = promotion
	mem.FillSuperPartner(partner)
	mem.MemberLevel = model.GENERAL
	mem.Avatar = "https://mszlu.oss-cn-beijing.aliyuncs.com/mscoin/defaultavatar.png"
	err := d.MemberRepo.Save(ctx, mem)
	if err != nil {
		return errors.New("注册失败")
	}
	return nil
}

func NewMemberDomain(db *msdb.MsDB) *MemberDomain {
	return &MemberDomain{
		MemberRepo: dao.NewMemberDao(db),
	}
}
