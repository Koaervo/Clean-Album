package usecase

import entity "CLEAN/Application/Entity"

type AlbumUseCase interface {
	Create(entity.Album) error
	Update(entity.Album) error
	Delete(int64) error
	GetAllAlbums() ([]entity.Album, error)
	GetAlbumById(int64) (entity.Album, error)
}

type AlbumService struct {
	repo AlbumRepository
}

func NewAlbumService(repo AlbumRepository) AlbumUseCase {
	return &AlbumService{repo: repo}
}

func (as *AlbumService) Create(album entity.Album) error {
	return as.repo.Create(album)
}

func (as *AlbumService) Update(album entity.Album) error {
	return as.repo.Update(album)
}

func (as *AlbumService) Delete(id int64) error {
	var album entity.Album

	album.IsDelete = true

	return as.repo.Update(album)
}

func (as *AlbumService) GetAllAlbums() ([]entity.Album, error) {
	return as.repo.GetAllAlbums()
}

func (as *AlbumService) GetAlbumById(id int64) (entity.Album, error) {
	return as.repo.GetAlbumById(id)
}
