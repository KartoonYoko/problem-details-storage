package buckets

import (
	"github.com/KartoonYoko/problem-details-storage/internal/storage"
)

type Usecase struct {
	storage *storage.Storage
}

func New(storage *storage.Storage) *Usecase {
	uc := new(Usecase)

	uc.storage = storage

	return uc
}
