package service

import (
	"context"

	"github.com/yvv4git/userinfo/internal/domain/usecases"
)

type UserInfoProcessor interface {
	ProcessUserInfo(ctx context.Context, params *usecases.ParamsProcessUserInfo) error
}

type UserInfo struct {
	ucUserInfo UserInfoProcessor
}

func NewUserInfo(ucUserInfo UserInfoProcessor) *UserInfo {
	return &UserInfo{
		ucUserInfo: ucUserInfo,
	}
}

type ParamsProcessUserInfo struct {
	Data map[string]any
}

func (u *UserInfo) ProcessUserInfo(ctx context.Context, params *ParamsProcessUserInfo) error {
	if params.Data == nil {
		return ErrGotEmptyData
	}

	u.ucUserInfo.ProcessUserInfo(ctx, &usecases.ParamsProcessUserInfo{
		Data: params.Data,
	})

	return nil
}
