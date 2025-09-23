package characters

import (
	"context"
	"github.com/jinzhu/copier"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCharacterByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCharacterByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCharacterByIdLogic {
	return &GetCharacterByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCharacterByIdLogic) GetCharacterById(req *types.GetCharacterByIdReq) (resp *types.GetCharacterByIdResp, err error) {
	character, err := l.svcCtx.CharactersModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return
	}

	resp = &types.GetCharacterByIdResp{}
	_ = copier.Copy(&resp.Character, character)
	return
}
