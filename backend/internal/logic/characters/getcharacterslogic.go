package characters

import (
	"context"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCharactersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCharactersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCharactersLogic {
	return &GetCharactersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCharactersLogic) GetCharacters(_ *types.GetCharactersReq) (resp *types.GetCharactersResp, err error) {
	characters, err := l.svcCtx.CharactersModel.List(l.ctx)
	if err != nil {
		return
	}

	resp = &types.GetCharactersResp{}
	for _, character := range characters {
		var c types.Character
		_ = copier.Copy(&c, character)
		resp.List = append(resp.List, &c)
	}

	return
}
