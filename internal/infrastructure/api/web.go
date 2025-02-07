package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/yvv4git/userinfo/internal/infrastructure/config"
	"github.com/yvv4git/userinfo/internal/service"
)

type Web struct {
	logger      *slog.Logger
	cfg         *config.API
	svcUserInfo *service.UserInfo
}

func NewWeb(logger *slog.Logger, cfg *config.API, svcUserInfo *service.UserInfo) *Web {
	return &Web{
		logger:      logger,
		cfg:         cfg,
		svcUserInfo: svcUserInfo,
	}
}

func (w *Web) Run(ctx context.Context) error {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.POST("/userinfo", w.UserInfoHandler)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", w.cfg.Host, w.cfg.Port),
		Handler: router,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			w.logger.Error("HTTP server error", slog.String("error", err.Error()))
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Duration(w.cfg.ShutdownTimeout)*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		w.logger.Error("HTTP server shutdown error", slog.String("error", err.Error()))
		return err
	}

	w.logger.Info("HTTP server stopped")
	return nil
}

func (w *Web) UserInfoHandler(c *gin.Context) {
	var jsonData map[string]interface{}

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		w.logger.Error("Failed to bind JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	err := w.svcUserInfo.ProcessUserInfo(
		c.Request.Context(),
		&service.ParamsProcessUserInfo{Data: jsonData},
	)
	if err != nil {
		w.logger.Error("Failed to process user info", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "JSON received and logged"})
}
