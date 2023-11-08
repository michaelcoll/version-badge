package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/patrickmn/go-cache"

	"github.com/michaelcoll/version-badge/internal/back/domain"
)

type ActuatorInfoCaller struct {
	domain.AppInfoCaller

	config YamlAppConfig
	c      *cache.Cache
}

func NewActuatorInfoCaller(config YamlAppConfig) *ActuatorInfoCaller {
	c := cache.New(1*time.Minute, 10*time.Second)
	return &ActuatorInfoCaller{c: c, config: config}
}

func (c *ActuatorInfoCaller) Get(ctx context.Context, env string, appName string) (*domain.AppInfo, error) {

	url, duration, compareWithEnv, err := c.config.Get(env, appName)
	if err != nil {
		return nil, err
	}

	actuatorInfo, err := c.getInfo(ctx, url, duration)
	if err != nil {
		return nil, err
	}

	upToDate := true
	if compareWithEnv != "" {
		url, duration, _, err := c.config.Get(compareWithEnv, appName)
		if err != nil {
			return nil, err
		}

		otherActuatorInfo, err := c.getInfo(ctx, url, duration)
		if err != nil {
			return nil, err
		}

		if actuatorInfo.Git.Commit.Id != otherActuatorInfo.Git.Commit.Id {
			upToDate = false
		}
	}

	return &domain.AppInfo{
		Version:   actuatorInfo.Build.Version,
		CommitSha: actuatorInfo.Git.Commit.Id,
		UpToDate:  upToDate,
	}, nil
}

func (c *ActuatorInfoCaller) getInfo(ctx context.Context, url string, duration time.Duration) (*ActuatorInfo, error) {

	if actuatorInfo, found := c.c.Get(url); found {
		return actuatorInfo.(*ActuatorInfo), nil
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Printf("%s Can't create request to %s\n", color.HiYellowString("i"), url)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("%s Can't access %s\n", color.HiYellowString("i"), url)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("%s Can't read response from %s\n", color.HiYellowString("i"), url)
		return nil, err
	}

	actuatorInfo := ActuatorInfo{}
	jsonErr := json.Unmarshal(body, &actuatorInfo)
	if jsonErr != nil {
		fmt.Printf("%s Can't unmarshal response from %s\n", color.HiYellowString("i"), url)
		return nil, err
	}

	c.c.Set(url, &actuatorInfo, duration)

	return &actuatorInfo, nil
}
