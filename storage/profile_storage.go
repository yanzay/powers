package storage

import (
	"encoding/json"

	"github.com/yanzay/powers/models"
)

// ProfileStorage has methods for accessing player Profiles
type ProfileStorage interface {
	GetProfile(int64) (*models.Profile, error)
	SetProfile(int64, *models.Profile) error
}

var profilesBucket = []byte("profiles")

func (bs *boltStorage) GetProfile(id int64) (*models.Profile, error) {
	profBytes, err := bs.get(profilesBucket, id)
	if err != nil {
		return nil, err
	}
	prof := &models.Profile{}
	err = json.Unmarshal(profBytes, prof)
	return prof, err
}

func (bs *boltStorage) SetProfile(id int64, prof *models.Profile) error {
	profBytes, err := json.Marshal(prof)
	if err != nil {
		return err
	}
	return bs.set(profilesBucket, id, profBytes)
}
