// +build !production

package assets

import (
	"os"
	"path/filepath"
	"sync"
	"time"
)

var statstatus struct {
	mutex sync.Mutex
	fails bool
}

func ModTime(prefix, path string) (time.Time, error) {
	now := time.Now()
	if fails := func() bool {
		statstatus.mutex.Lock()
		defer statstatus.mutex.Unlock()
		return statstatus.fails
	}(); fails {
		return now, nil
	}
	fi, err := os.Stat(filepath.Join(prefix, path))
	if err != nil {
		func() {
			statstatus.mutex.Lock()
			defer statstatus.mutex.Unlock()
			statstatus.fails = true
		}()
		return time.Time{}, err
	}
	return fi.ModTime(), nil
}

var Uncompressedasset = Asset
