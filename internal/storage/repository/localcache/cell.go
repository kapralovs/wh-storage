package localcache

import (
	"errors"
	"sync"

	"github.com/kapralovs/wh-storage/internal/models"
)

type LocalRepo struct {
	mu    *sync.Mutex
	cells map[string]*models.Cell
}

func NewLocalRepo() *LocalRepo {
	return &LocalRepo{
		mu:    new(sync.Mutex),
		cells: make(map[string]*models.Cell),
	}
}

func (r *LocalRepo) CreateCell() error {
	return nil
}
func (r *LocalRepo) UpdateCell(strID string) error {
	return nil
}
func (r *LocalRepo) DeleteCell(strID string) error {
	return nil
}
func (r *LocalRepo) GetCell(strID string) (*models.Cell, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if cell, ok := r.cells[strID]; ok {
		return cell, nil
	}

	return nil, errors.New("нет ячейки с таким id")
}
func (r *LocalRepo) GetCellsList() ([]*models.Cell, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var cellsList = make([]*models.Cell, len(r.cells))

	for _, cell := range r.cells {
		cellsList = append(cellsList, cell)
	}

	return cellsList, nil
}
