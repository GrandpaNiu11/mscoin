package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	common "mscoin-common"
	"mscoin-common/tools"

	"net/http"
	"ucenter-api/internal/logic"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"
)

type LoginHandler struct {
	svcCtx *svc.ServiceContext
}

func NewLoginHandler(svcCtx *svc.ServiceContext) *LoginHandler {
	return &LoginHandler{
		svcCtx: svcCtx,
	}
}

func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req types.LoginReq
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	//获取ip
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewLoginLogic(r.Context(), h.svcCtx)
	resp, error := l.Login(&req)
	result := common.NewResult().Deal(resp, error)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *LoginHandler) chackLogin(w http.ResponseWriter, r *http.Request) {
	result := common.NewResult()
	tolen := r.Header.Get("x-auth-token")
	l := logic.NewLoginLogic(r.Context(), h.svcCtx)
	flag, _ := l.CheckLogin(tolen)
	result = common.NewResult().Deal(flag, nil)
	httpx.OkJsonCtx(r.Context(), w, result)
}
