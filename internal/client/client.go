package client

import (
	"github.com/goofansu/linkding-cli/internal/config"
	"github.com/piero-vic/go-linkding"
)

func NewClient() (*linkding.Client, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	return linkding.NewClient(cfg.Endpoint, cfg.APIKey), nil
}
