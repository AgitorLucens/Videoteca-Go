package storage

import (
	"gorm.io/gorm"
)

type GenreRepositoryImpl struct {
	DB *gorm.DB
}
