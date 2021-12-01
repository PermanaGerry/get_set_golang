package inter

import (
	"get_set_golang/model"
)

type InterfaceRepoNota interface {
	GetAllNota() []string
	GetAllTotal() []string
	FindByNota(nota string) []model.DataNota
	FindByTotal(total int) []model.DataNota
}
