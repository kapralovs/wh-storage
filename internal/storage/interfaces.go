package internal

import "github.com/kapralovs/wh-storage/internal/models"

type Repository interface {
	CreateCell() error
	UpdateCell(strID string) error
	DeleteCell(strID string) error
	GetCell(strID string) (*models.Cell, error)
	GetCellsList() ([]*models.Cell, error)
}

type Usecase interface {
	CreateCell() error
	UpdateCell(strID string) error
	DeleteCell(strID string) error
	GetCell(strID string) (*models.Cell, error)
	GetCellsList() ([]*models.Cell, error)
}
