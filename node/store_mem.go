package node

import (
	"sync"

	"github.com/ipfs/go-datastore"

	"github.com/celestiaorg/celestia-node/core"
	"github.com/celestiaorg/celestia-node/libs/keystore"
)

type memStore struct {
	keys keystore.Keystore
	data datastore.Batching
	core core.Store
	cfg  *Config
	cfgL sync.Mutex
}

// NewMemStore creates an in-memory Store for Node.
// Useful for testing.
func NewMemStore() Store {
	return &memStore{
		keys: keystore.NewMapKeystore(),
		data: datastore.NewMapDatastore(),
		core: core.NewMemStore(),
	}
}

func (m *memStore) Keystore() (keystore.Keystore, error) {
	return m.keys, nil
}

func (m *memStore) Datastore() (datastore.Batching, error) {
	return m.data, nil
}

func (m *memStore) Core() (core.Store, error) {
	return m.core, nil
}

func (m *memStore) Config() (*Config, error) {
	m.cfgL.Lock()
	defer m.cfgL.Unlock()
	return m.cfg, nil
}

func (m *memStore) PutConfig(cfg *Config) error {
	m.cfgL.Lock()
	defer m.cfgL.Unlock()
	m.cfg = cfg
	return nil
}

func (m *memStore) Path() string {
	return ""
}

func (m *memStore) Close() error {
	return nil
}
