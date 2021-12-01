package inter

import (
	"test_pointer/model"
)

type InterfaceRepoNota interface {
	GetAllNota() []string
	GetAllTotal() []string
	FindByNota(nota string) []model.DataNota
	FindByTotal(total int) []model.DataNota
}
