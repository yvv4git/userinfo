package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yvv4git/userinfo/internal/infrastructure/config"
	"github.com/yvv4git/userinfo/internal/infrastructure/logger"
	"github.com/yvv4git/userinfo/internal/utils"
)

func TestConfigLoad(t *testing.T) {
	// Change directory to project root
	require.NoError(t, utils.ChangeDirToProjectRoot("../../"), "Should change directory without errors")

	// Setup logger & config
	logger := logger.SetupDefaultLogger()
	cfg := config.NewConfig(logger)

	// Load config
	err := cfg.Load()
	assert.NoError(t, err, "Should load config without errors")
}
