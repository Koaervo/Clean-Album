package adapter

import (
	entity "CLEAN/Application/Entity"
	usecase "CLEAN/Application/Use-Case"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type HttpAlbumHandler struct {
	albumUseCase usecase.AlbumUseCase
}

func NewHttpAlbumHandler(useCase usecase.AlbumUseCase) *HttpAlbumHandler {
	return &HttpAlbumHandler{albumUseCase: useCase}
}

func (h *HttpAlbumHandler) CreateAlbum(c *fiber.Ctx) error {
	var album entity.Album
	if err := c.BodyParser(&album); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "bad request"})
	}

	if err := h.albumUseCase.Create(album); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(album)
}

func (h *HttpAlbumHandler) UpdateAlbum(c *fiber.Ctx) error {
	var album entity.Album
	if err := c.BodyParser(&album); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "bad request"})
	}

	if err := h.albumUseCase.Update(album); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(album)
}

func (h *HttpAlbumHandler) DeleteAlbum(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.ParseInt(idStr, 64, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "bad request"})
	}

	if err := h.albumUseCase.Delete(id); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON("")
}

func (h *HttpAlbumHandler) GetAllAlbum(c *fiber.Ctx) error {
	var albums []entity.Album

	albums, err := h.albumUseCase.GetAllAlbums()
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(albums)
}

func (h *HttpAlbumHandler) GetAlbumById(c *fiber.Ctx) error {
	var album entity.Album
	idStr := c.Params("id")

	id, err := strconv.ParseInt(idStr, 64, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "bad request"})
	}

	album, err = h.albumUseCase.GetAlbumById(id)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(album)
}
