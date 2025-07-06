package public

import (
	"context"
	"rui/internal"
	"rui/internal/api/public"
	"rui/internal/repo"
	"rui/internal/svc"
)

type InitStatusLogic struct {
	Ctx    context.Context
	SvcCtx *svc.ServiceContext
}

func NewInitStatusLogic(context context.Context, svcCtx *svc.ServiceContext) *InitStatusLogic {
	return &InitStatusLogic{
		Ctx:    context,
		SvcCtx: svcCtx,
	}
}

func (l *InitStatusLogic) InitStatusLogic() (resp *public.InitStatusRes, err error) {
	user, err := l.SvcCtx.Repo.UserDetail(l.Ctx, repo.User{Level: internal.AdminLevel})
	if err != nil {
		return resp, err
	}

	if user.Id <= 0 {
		return &public.InitStatusRes{Status: 2}, l.SvcCtx.Redis.Set(l.Ctx, internal.InitStatusRdsKey, 2, 0).Err()
	}

	sysConf, err := l.SvcCtx.Repo.SysConfDetailByNode(l.Ctx,
		repo.SysConf{Code: internal.NginxConfKey, IsLocal: internal.True})
	if err != nil {
		return resp, err
	}

	if sysConf.Id <= 0 {
		return &public.InitStatusRes{Status: 3}, l.SvcCtx.Redis.Set(l.Ctx, internal.InitStatusRdsKey, 2, 0).Err()
	}

	return &public.InitStatusRes{Status: 1}, l.SvcCtx.Redis.Del(l.Ctx, internal.InitStatusRdsKey).Err()
}
