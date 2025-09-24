package characters

import (
	"backend/pkg/code"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

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
		return nil, errors.Wrapf(code.NewInternalError("failed to get character by id"), "err: %v, req: %+v", err, req)
	}

	resp = &types.GetCharacterByIdResp{}
	_ = copier.Copy(&resp.Character, character)
	return
}
