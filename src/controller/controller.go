package controller

import (
	"github.com/danielmoisa/envoy/src/cache"
	"github.com/danielmoisa/envoy/src/drive"
	"github.com/danielmoisa/envoy/src/repository"
)

type Controller struct {
	Repository *repository.Repository
	Cache      *cache.Cache
	Drive      *drive.Drive
}

func NewControllerForBackend(repository *repository.Repository, cache *cache.Cache, drive *drive.Drive) *Controller {
	return &Controller{
		Repository: repository,
		Cache:      cache,
		Drive:      drive,
	}
}
