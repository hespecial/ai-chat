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

type CharacterSkillsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCharacterSkillsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CharacterSkillsLogic {
	return &CharacterSkillsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CharacterSkillsLogic) CharacterSkills(req *types.CharacterSkillsReq) (resp *types.CharacterSkillsResp, err error) {
	internalError := code.NewInternalError("failed to get character skills")

	ids, err := l.svcCtx.AssocCharacterSkillModel.GetSkillIds(l.ctx, req.CharacterId)
	if err != nil {
		return nil, errors.Wrapf(internalError, "get skill ids error: %v", err)
	}

	resp = &types.CharacterSkillsResp{
		Skills: make([]*types.Skill, 0, len(ids)),
	}
	for _, id := range ids {
		var skill types.Skill
		target, err := l.svcCtx.SkillModel.FindOne(l.ctx, id)
		if err != nil {
			return nil, errors.Wrapf(internalError, "finding skill %v error: %v", id, err)
		}
		_ = copier.Copy(&skill, target)
		resp.Skills = append(resp.Skills, &skill)
	}

	return
}
