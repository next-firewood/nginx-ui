package public

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"rui/internal/api/public"
	"rui/internal/svc"
)

type LoginLogic struct {
	Ctx    context.Context
	SvcCtx *svc.ServiceContext
}

func NewLoginLogic(context context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Ctx:    context,
		SvcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *public.LoginReq) (resp *public.LoginResp, err error) {
	token, err := l.SvcCtx.Config.Auth.GenerateToken([]byte(l.SvcCtx.Config.Auth.SecretKey), jwt.MapClaims{
		"username": req.Username,
	})

	if err != nil {
		return resp, err
	}

	return &public.LoginResp{Token: token}, err
}
