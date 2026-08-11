package storage

import (
	"gorm.io/gorm"
)

type ActorRepositoryImpl struct {
	DB *gorm.DB
}
