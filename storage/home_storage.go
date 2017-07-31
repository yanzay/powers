package storage

import (
	"encoding/json"

	"github.com/yanzay/powers/models"
)

// HomeStorage has methods for accessing player Homes
type HomeStorage interface {
	GetHome(int64) (*models.Home, error)
	SetHome(int64, *models.Home) error
}

var homesBucket = []byte("homes")

func (bs *boltStorage) GetHome(id int64) (*models.Home, error) {
	homeBytes, err := bs.get(homesBucket, id)
	if err != nil {
		return nil, err
	}
	home := &models.Home{}
	err = json.Unmarshal(homeBytes, home)
	return home, err
}

func (bs *boltStorage) SetHome(id int64, home *models.Home) error {
	homeBytes, err := json.Marshal(home)
	if err != nil {
		return err
	}
	return bs.set(homesBucket, id, homeBytes)
}
