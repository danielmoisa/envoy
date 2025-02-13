package controller

import (
	"github.com/danielmoisa/envoy/src/cache"
	"github.com/danielmoisa/envoy/src/drive"
	"github.com/danielmoisa/envoy/src/storage"
)

type Controller struct {
	Storage *storage.Storage
	Cache   *cache.Cache
	Drive   *drive.Drive
}

func NewControllerForBackend(storage *storage.Storage, cache *cache.Cache, drive *drive.Drive) *Controller {
	return &Controller{
		Storage: storage,
		Cache:   cache,
		Drive:   drive,
	}
}
