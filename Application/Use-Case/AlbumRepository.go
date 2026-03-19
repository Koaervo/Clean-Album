package usecase

import entity "CLEAN/Application/Entity"

type AlbumRepository interface {
	Create(entity.Album) error
	Update(entity.Album) error
	Delete(int64) error
	GetAllAlbums() ([]entity.Album, error)
	GetAlbumById(int64) (entity.Album, error)
}
