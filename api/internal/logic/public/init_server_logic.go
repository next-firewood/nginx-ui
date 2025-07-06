package public

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"rui/common/database"
	"rui/common/errorx"
	"rui/internal"
	"rui/internal/api/public"
	"rui/internal/repo"
	"rui/internal/svc"
	"strconv"
)

type InitServerLogic struct {
	Ctx    context.Context
	SvcCtx *svc.ServiceContext
}

func NewInitServerLogic(context context.Context, svcCtx *svc.ServiceContext) *InitServerLogic {
	return &InitServerLogic{
		Ctx:    context,
		SvcCtx: svcCtx,
	}
}

func (l *InitServerLogic) InitServerLogic(req *public.InitServerReq) (err error) {
	val, err := l.SvcCtx.Redis.Get(l.Ctx, internal.InitStatusRdsKey).Result()
	if val == "" || errors.Is(err, redis.Nil) {
		err = nil
		return errorx.BusinessErr("请勿重复初始化")
	}

	if err != nil {
		return err
	}

	status, err := strconv.Atoi(val)
	if err != nil {
		return err
	}

	ctx, clean := l.SvcCtx.Repo.BeginTx(l.Ctx)

	if status == internal.ServerUserStatus {
		if _, err = l.SvcCtx.Repo.UserNew(ctx, repo.User{
			GormModel: database.GormModel{Uuid: uuid.NewString()},
			Username:  req.Username, Password: req.Password, Level: internal.AdminLevel,
		}); err != nil {
			return clean(err)
		}
	}

	if err = l.SvcCtx.Repo.SysConfDetailNew(ctx, repo.SysConf{
		Code: internal.NginxConfKey, Value: req.NginxUrl, Node: req.Node, IsLocal: internal.True,
	}); err != nil {
		return clean(err)
	}

	l.SvcCtx.Redis.Del(l.Ctx, internal.InitStatusRdsKey)

	return clean(err)
}
