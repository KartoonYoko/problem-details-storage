package buckets

import (
	"github.com/KartoonYoko/problem-details-storage/internal/storage"
)

type bucketsUsecase struct {
	storage *storage.Storage
}

func New(storage *storage.Storage) *bucketsUsecase {
	uc := new(bucketsUsecase)

	uc.storage = storage

	return uc
}
