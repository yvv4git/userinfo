package usecases

import (
	"context"
	"log/slog"
)

type DstConsole struct {
	logger *slog.Logger
}

func NewDstConsole(logger *slog.Logger) *DstConsole {
	return &DstConsole{
		logger: logger,
	}
}

type ParamsProcessUserInfo struct {
	Data map[string]any
}

func (d *DstConsole) ProcessUserInfo(ctx context.Context, params *ParamsProcessUserInfo) error {
	d.logger.Info("Received JSON data", slog.Any("data", params.Data))

	return nil
}
