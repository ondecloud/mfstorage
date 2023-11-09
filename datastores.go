package mfstorage

import (
	"sync"

	"github.com/ipfs/go-datastore"
)

var dss struct {
	sync.RWMutex
	cfgMaps map[string]ConfigFromMaps
}

func init() {
	dss.cfgMaps = map[string]ConfigFromMaps{
		"mount": MountDatastoreConfig,
	}
}

type StoreInfo map[string]interface{}

type DatastoreConfig interface {
	StoreInfo() StoreInfo

	// Create instantiate a new datastore from this config
	Create(path string) (datastore.Datastore, error)
}

// RegisterStorage registers a storage with the given name and configuration.
//
// Parameters:
// - name: the name of the storage to register.
// - cfg: the configuration for the storage.
func RegisterStorage(name string, cfg ConfigFromMaps) {
	dss.Lock()
	defer dss.Unlock()
	//panic if exists
	if _, ok := dss.cfgMaps[name]; ok {
		panic("storage already registered: " + name)
	}
	dss.cfgMaps[name] = cfg
}

// LoadStorage loads a datastore configuration based on the given name and configuration map.
//
// Parameters:
// - name: the name of the datastore to load.
// - cfg: the configuration map for the datastore.
//
// Returns:
// - DatastoreConfig: the loaded datastore configuration.
// - error: an error if the datastore is not found.
func LoadStorage(name string, cfg map[string]any) (DatastoreConfig, error) {
	dss.RLock()
	defer dss.RUnlock()
	configMaps, ok := dss.cfgMaps[name]
	if !ok {
		return nil, ErrStoreNotFound(name)
	}
	return configMaps(cfg)
}
