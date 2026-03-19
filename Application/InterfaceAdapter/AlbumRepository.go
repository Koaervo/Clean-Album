package adapter

import (
	entity "CLEAN/Application/Entity"
	usecase "CLEAN/Application/Use-Case"

	"gorm.io/gorm"
)

type GormAlbumRepository struct {
	db *gorm.DB
}

func NewGormAlbumRepository(db *gorm.DB) usecase.AlbumRepository {
	return &GormAlbumRepository{db: db}
}

func (r *GormAlbumRepository) Create(album entity.Album) error {
	return r.db.Create(&album).Error
}

func (r *GormAlbumRepository) Update(album entity.Album) error {
	return r.db.Save(&album).Error
}

func (r *GormAlbumRepository) Delete(album entity.Album) error {
	return r.db.Delete(&album).Error
}

func (r *GormAlbumRepository) GetAllAlbums() ([]entity.Album, error) {
	var album []entity.Album
	err := r.db.Preload("Songs").Find(&album).Error
	return album, err
}

func (r *GormAlbumRepository) GetAlbumById(id int64) (entity.Album, error) {
	var album entity.Album
	err := r.db.Preload("Songs").Where("id = ?", id).Find(&album).Error
	return album, err
}
