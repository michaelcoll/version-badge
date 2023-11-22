package domain

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTemplateValues_WithActuatorError(t *testing.T) {
	service := BadgeService{}
	actuatorErr := errors.New("actuator error")

	result := service.buildTemplateValues(nil, "testEnv", actuatorErr)

	assert.Equal(t, Red, result.Color)
	assert.Equal(t, "DOWN", result.Tag)
	assert.Equal(t, "testEnv", result.Env)
}

func TestBuildTemplateValues_WithOutdatedAppInfo(t *testing.T) {
	service := BadgeService{}
	appInfo := &AppInfo{UpToDate: false, Version: "1.0.0", CommitSha: "abc123"}

	result := service.buildTemplateValues(appInfo, "testEnv", nil)

	assert.Equal(t, Yellow, result.Color)
	assert.Equal(t, "1.0.0", result.Tag)
	assert.Equal(t, "testEnv", result.Env)
}

func TestBuildTemplateValues_WithSnapshotVersion(t *testing.T) {
	service := BadgeService{}
	appInfo := &AppInfo{UpToDate: true, Version: "1.0.0-SNAPSHOT", CommitSha: "abc123"}

	result := service.buildTemplateValues(appInfo, "testEnv", nil)

	assert.Equal(t, Green, result.Color)
	assert.Equal(t, "abc123", result.Tag)
	assert.Equal(t, "testEnv", result.Env)
}

func TestBuildTemplateValues_WithRegularVersion(t *testing.T) {
	service := BadgeService{}
	appInfo := &AppInfo{UpToDate: true, Version: "1.0.0", CommitSha: "abc123"}

	result := service.buildTemplateValues(appInfo, "testEnv", nil)

	assert.Equal(t, Green, result.Color)
	assert.Equal(t, "1.0.0", result.Tag)
	assert.Equal(t, "testEnv", result.Env)
}
