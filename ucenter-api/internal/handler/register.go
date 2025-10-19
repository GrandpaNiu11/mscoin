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

type RegisterHandler struct {
	svcCtx *svc.ServiceContext
}

func NewRegisterHandler(svcCtx *svc.ServiceContext) *RegisterHandler {
	return &RegisterHandler{
		svcCtx: svcCtx,
	}
}

func (h *RegisterHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req types.Request
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	//获取ip
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewUcenterapiLogic(r.Context(), h.svcCtx)
	resp, error := l.Register(&req)
	result := common.NewResult().Deal(resp, error)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *RegisterHandler) SendCode(w http.ResponseWriter, r *http.Request) {
	var req types.CodeRequest
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	l := logic.NewUcenterapiLogic(r.Context(), h.svcCtx)
	resp, error := l.SendCode(&req)
	result := common.NewResult().Deal(resp, error)
	httpx.OkJsonCtx(r.Context(), w, result)
}
