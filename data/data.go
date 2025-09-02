package data

import (
	"api-client/dto"
	err "api-client/error"
	"encoding/json"
	"errors"
	"os"
	"sync"
)

var (
	Products []dto.Item
	once     sync.Once
)

func InitProducts() error {
	var err error
	once.Do(func() {
		Products, err = SyncProducts()
	})
	return err
}

func FindByName(name string) (dto.Item, error) {
	for _, v := range Products {
		if name == v.Name {
			return v, nil
		}
	}

	return dto.Item{}, err.ErrProductEmpty
}

func SyncProducts() ([]dto.Item, error) {
	file, err := os.Open("data.json")
	if err != nil {
		return nil, errors.New("cannot Open data.json")
	}
	defer file.Close()

	var items []dto.Item
	if err := json.NewDecoder(file).Decode(&items); err != nil {
		return nil, errors.New("failed to decode JSON")
	}

	return items, nil
}
